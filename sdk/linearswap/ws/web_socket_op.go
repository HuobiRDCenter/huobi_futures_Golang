package ws

import (
	"container/list"
	"encoding/json"
	"fmt"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/reqbuilder"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/wsbase"
	"reflect"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type MethonInfo struct {
	fun   interface{}
	param reflect.Type
}

func (wsOp *MethonInfo) Init(fun interface{}, param reflect.Type) *MethonInfo {
	wsOp.fun = fun
	wsOp.param = param
	return wsOp
}

type WebSocketOp struct {
	host string
	path string

	conn      *websocket.Conn
	accessKey string
	secretKey string

	autoConnect bool
	allSubStrs  list.List

	onSubCallbackFuns map[string]*MethonInfo
	onReqCallbackFuns map[string]*MethonInfo

	authOk bool

	// v5Mode 由 V5 客户端（WSV5NotifyClient）Init 时置位。
	// V5 推送把订阅/退订应答（op=sub/unsub）也按 topic 路由交回调，调用方据此判断订阅就绪；
	// 口径对齐 C++ WSSwapNotificationClient（"sub/unsub ack 按 topic 路由交 callback"）。
	// v3 旧端点（/linear-swap-notification 等）沿用只记日志的原行为，避免给既有用户多出一次回调。
	v5Mode bool
}

func (wsOp *WebSocketOp) open(path string, host string, accessKey string, secretKey string, autoConnect bool) bool {
	if host == "" {
		wsOp.host = linearswap.LINEAR_SWAP_DEFAULT_HOST
	}

	wsOp.host = host
	wsOp.path = path

	wsOp.accessKey = accessKey
	wsOp.secretKey = secretKey

	wsOp.autoConnect = autoConnect

	ret := wsOp.connServer()
	if ret {
		wsOp.allSubStrs = list.List{}
		wsOp.onSubCallbackFuns = make(map[string]*MethonInfo)
		wsOp.onReqCallbackFuns = make(map[string]*MethonInfo)
	}
	return ret
}

func (wsOp *WebSocketOp) close() {
	wsOp.conn.Close()
	wsOp.conn = nil
}

func (wsOp *WebSocketOp) connServer() bool {
	url := fmt.Sprintf("wss://%s%s", wsOp.host, wsOp.path)
	var err error
	wsOp.conn, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Error("WebSocket connected error: %s", err)
		return false
	}
	log.Info("WebSocket connected")
	wsOp.conn.SetCloseHandler(wsOp.onClose)

	go wsOp.readLoop(wsOp.conn)

	if wsOp.accessKey == "" && wsOp.secretKey == "" {
		wsOp.authOk = true
		return true
	}
	wsOp.authOk = false
	return wsOp.sendAuth(wsOp.conn, wsOp.host, wsOp.path, wsOp.accessKey, wsOp.secretKey)
}

func (wsOp *WebSocketOp) onClose(code int, text string) error {
	log.Info("WebSocket close.")

	wsOp.conn = nil
	if wsOp.autoConnect {
		if !wsOp.connServer() {
			return fmt.Errorf("reconnect server error")
		}
	}

	return fmt.Errorf("")
}

func (wsOp *WebSocketOp) sendAuth(conn *websocket.Conn, host string, path string, accessKey string, secretKey string) bool {
	if conn == nil {
		log.Error("websocket conn is null")
		return false
	}

	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05")

	req := new(reqbuilder.GetRequest).Init()
	req.AddParam("AccessKeyId", accessKey)
	req.AddParam("SignatureMethod", "HmacSHA256")
	req.AddParam("SignatureVersion", "2")
	req.AddParam("Timestamp", timestamp)

	sign := new(reqbuilder.Signer).Init(secretKey)
	signature := sign.Sign("GET", host, path, req.BuildParams())

	auth := wsbase.WSAuthData{
		Op:               "auth",
		AtType:           "api",
		AccessKeyId:      accessKey,
		SignatureMethod:  "HmacSHA256",
		SignatureVersion: "2",
		Timestamp:        timestamp,
		Signature:        signature}

	jdata, error := json.Marshal(&auth)
	if error != nil {
		log.Error("Auth to json error.")
		return false
	}
	conn.WriteMessage(websocket.TextMessage, jdata)
	return true
}

func (wsOp *WebSocketOp) readLoop(conn *websocket.Conn) {
	for conn != nil {
		msgType, buf, err := conn.ReadMessage()
		if err != nil {
			log.Error("Read error: %s", err)
			continue
		}
		var message string
		if msgType == websocket.BinaryMessage {
			message, err = wsbase.GZipDecompress(buf)
			if err != nil {
				log.Error("UnGZip data error: %s", err)
				continue
			}
		} else if msgType == websocket.TextMessage {
			message = string(buf)
		}

		var jdata map[string]interface{}
		err = json.Unmarshal([]byte(message), &jdata)
		if err != nil {
			log.Error("msg to map[string]json.RawMessage error: %s", err)
			continue
		}

		if ts, found := jdata["ping"]; found { // market heartbeat
			ts = int64(ts.(float64))
			//log.Info("WebSocket received data: \"ping\":%d", ts)

			pongData := fmt.Sprintf("{\"pong\":%d }", ts)
			wsOp.conn.WriteMessage(websocket.TextMessage, []byte(pongData))

			//log.Info("WebSocket replied data: %s", pongData)
		} else if op, found := jdata["op"]; found {
			switch op {
			case "ping": // order heartbeat
				ts := jdata["ts"]
				//log.Info("WebSocket received data, { \"op\":\"%s\", \"ts\": \"%s\" }", op, ts)

				pongData := fmt.Sprintf("{ \"op\":\"pong\", \"ts\": \"%s\" }", ts)
				wsOp.conn.WriteMessage(websocket.TextMessage, []byte(pongData))

				//log.Info("WebSocket replied data, %s", pongData)

			case "close":
				log.Error("Some error occurres when authentication in server side.")

			case "error":
				log.Error("Illegal op or internal error, but websoket is still connected.")

			case "auth":
				code := int64(jdata["err-code"].(float64))
				if code == 0 {
					log.Info("Authentication success.")
					wsOp.authOk = true
					for e := wsOp.allSubStrs.Front(); e != nil; e = e.Next() {
						wsOp.conn.WriteMessage(websocket.TextMessage, []byte(e.Value.(string)))
					}
				} else {
					msg := jdata["err-msg"].(string)
					log.Error("Authentication failure: %d/%s", code, msg)
					wsOp.close()
				}
			case "notify":
				topic := jdata["topic"].(string)
				wsOp.handleSubCallbackFun(topic, message, jdata)
			case "sub":
				topic := jdata["topic"]
				log.Info("sub: \"%s\"", topic)

				// V5 推送把订阅应答也交给 topic 回调，调用方据此判断订阅就绪
				// （v3 只记日志，理由见 WebSocketOp.v5Mode 注释）
				if wsOp.v5Mode {
					if t, ok := topic.(string); ok {
						wsOp.handleSubAckFun(t, message)
					}
				}
			case "unsub":
				topic := jdata["topic"].(string)
				log.Info("unsub: \"%s\"", topic)

				if wsOp.v5Mode {
					wsOp.handleSubAckFun(topic, message)
				}
				if _, found := wsOp.onSubCallbackFuns[topic]; found {
					delete(wsOp.onSubCallbackFuns, topic)
				}
			default:
				log.Info("WebSocket received unknow data: %s", jdata)
			}
		} else if topic, found := jdata["subbed"]; found { // sub success reply
			log.Info("\"subbed\": \"%s\"", topic)
		} else if topic, found := jdata["unsubbed"]; found { // unsub success reply
			log.Info("\"unsubbed\": \"%s\"", topic)

			if _, found := wsOp.onSubCallbackFuns[topic.(string)]; found {
				delete(wsOp.onSubCallbackFuns, topic.(string))
			}
		} else if topic, found := jdata["ch"]; found { // market sub reply data

			wsOp.handleSubCallbackFun(topic.(string), message, jdata)
		} else if topic, found := jdata["rep"]; found { // market request reply data
			wsOp.handleReqCallbackFun(topic.(string), message, jdata)
		} else if cid, found := jdata["cid"]; found { // V5 交易指令应答：无 op 字段，靠 cid 关联请求
			wsOp.handleReqCallbackFun(fmt.Sprintf("%v", cid), message, jdata)
		} else if code, found := jdata["err-code"]; found { // market request reply data
			msg := jdata["err-msg"]
			log.Error("%d:%s", code, msg)
		} else {
			log.Info("WebSocket received unknow data: %s", jdata)
		}
	}
}

func (wsOp *WebSocketOp) handleSubCallbackFun(ch string, data string, jdata map[string]interface{}) {
	var mi *MethonInfo = nil
	ch = strings.ToLower(ch)

	if _, found := wsOp.onSubCallbackFuns[ch]; found {
		mi = wsOp.onSubCallbackFuns[ch]
	} else if ch == "accounts" || ch == "positions" { // isolated
		data_array := jdata["data"].([]interface{})
		contract_code := data_array[0].(map[string]interface{})["contract_code"].(string)
		contract_code = strings.ToLower(contract_code)

		full_ch := fmt.Sprintf("%s.%s", ch, contract_code)
		if _, found := wsOp.onSubCallbackFuns[full_ch]; found {
			mi = wsOp.onSubCallbackFuns[full_ch]
		} else if _, found := wsOp.onSubCallbackFuns[fmt.Sprintf("%s.*", ch)]; found {
			mi = wsOp.onSubCallbackFuns[fmt.Sprintf("%s.*", ch)]
		}
	} else if strings.HasPrefix(ch, "orders.") {
		if _, found := wsOp.onSubCallbackFuns["orders.*"]; found {
			mi = wsOp.onSubCallbackFuns["orders.*"]
		}
	} else if strings.HasPrefix(ch, "matchorders.") {
		if _, found := wsOp.onSubCallbackFuns["matchorders.*"]; found {
			mi = wsOp.onSubCallbackFuns["matchorders.*"]
		}
	} else if strings.HasPrefix(ch, "trigger_order.") {
		if _, found := wsOp.onSubCallbackFuns["trigger_order.*"]; found {
			mi = wsOp.onSubCallbackFuns["trigger_order.*"]
		}
	} else if ch == "accounts_cross" { // isolated
		data_array := jdata["data"].([]interface{})
		margin_account := data_array[0].(map[string]interface{})["margin_account"].(string)
		margin_account = strings.ToLower(margin_account)

		full_ch := fmt.Sprintf("%s.%s", ch, margin_account)
		if _, found := wsOp.onSubCallbackFuns[full_ch]; found {
			mi = wsOp.onSubCallbackFuns[full_ch]
		} else if _, found := wsOp.onSubCallbackFuns[fmt.Sprintf("%s.*", ch)]; found {
			mi = wsOp.onSubCallbackFuns[fmt.Sprintf("%s.*", ch)]
		}
	} else if ch == "positions_cross" {
		data_array := jdata["data"].([]interface{})
		contract_code := data_array[0].(map[string]interface{})["contract_code"].(string)
		contract_code = strings.ToLower(contract_code)

		full_ch := fmt.Sprintf("%s.%s", ch, contract_code)
		if _, found := wsOp.onSubCallbackFuns[full_ch]; found {
			mi = wsOp.onSubCallbackFuns[full_ch]
		} else if _, found := wsOp.onSubCallbackFuns[fmt.Sprintf("%s.*", ch)]; found {
			mi = wsOp.onSubCallbackFuns[fmt.Sprintf("%s.*", ch)]
		}
	} else if strings.HasPrefix(ch, "orders_cross.") {
		if _, found := wsOp.onSubCallbackFuns["orders_cross.*"]; found {
			mi = wsOp.onSubCallbackFuns["orders_cross.*"]
		}
	} else if strings.HasPrefix(ch, "matchorders_cross.") {
		if _, found := wsOp.onSubCallbackFuns["matchorders_cross.*"]; found {
			mi = wsOp.onSubCallbackFuns["matchorders_cross.*"]
		}
	} else if strings.HasPrefix(ch, "trigger_order_cross.") {
		if _, found := wsOp.onSubCallbackFuns["trigger_order_cross.*"]; found {
			mi = wsOp.onSubCallbackFuns["trigger_order_cross.*"]
		}
	} else if strings.HasSuffix(ch, ".liquidation_orders") { // General
		if _, found := wsOp.onSubCallbackFuns["public.*.liquidation_orders"]; found {
			mi = wsOp.onSubCallbackFuns["public.*.liquidation_orders"]
		}
	} else if strings.HasSuffix(ch, ".funding_rate") {
		if _, found := wsOp.onSubCallbackFuns["public.*.funding_rate"]; found {
			mi = wsOp.onSubCallbackFuns["public.*.funding_rate"]
		}
	} else if strings.HasSuffix(ch, ".contract_info") {
		if _, found := wsOp.onSubCallbackFuns["public.*.contract_info"]; found {
			mi = wsOp.onSubCallbackFuns["public.*.contract_info"]
		}
	}

	if mi == nil {
		log.Error("no callback function to handle: %s", jdata)
		return
	}

	wsOp.runFunction(mi, data)
}

// handleSubAckFun 把 V5 的订阅/退订应答帧（op=sub/unsub）按 topic 交回调。
// 应答帧只带 topic、不带 data，不能复用 handleSubCallbackFun：那里的隔离/通配分支会去取 data 字段。
func (wsOp *WebSocketOp) handleSubAckFun(topic string, message string) {
	mi, found := wsOp.onSubCallbackFuns[strings.ToLower(topic)]
	if !found || mi == nil {
		return
	}
	wsOp.runFunction(mi, message)
}

func (wsOp *WebSocketOp) handleReqCallbackFun(ch string, data string, jdata map[string]interface{}) {
	var mi *MethonInfo = nil
	var found bool = false
	ch = strings.ToLower(ch)

	if mi, found = wsOp.onReqCallbackFuns[ch]; !found {
		log.Error("no callback function to handle: %s", jdata)
		return
	}

	if mi == nil {
		log.Error("no callback function to handle: %s", jdata)
		return
	}

	wsOp.runFunction(mi, data)
}

func (wsOp *WebSocketOp) runFunction(mi *MethonInfo, data string) {
	param := reflect.New(mi.param).Interface()
	if err := json.Unmarshal([]byte(data), &param); err != nil {
		// 解析失败时仍回调（与 Java/C++ 的"原样交回调"口径一致），但必须留下日志，
		// 否则 DTO 与线上字段不一致会完全静默
		log.Error("unmarshal ws response error: %s, raw: %s", err, data)
	}

	rargs := make([]reflect.Value, 1)
	rargs[0] = reflect.ValueOf(param)

	fun := reflect.ValueOf(mi.fun)
	fun.Call(rargs)
}

// wsAuthTimeout 是等待 WS 鉴权完成的上限。
const wsAuthTimeout = 15 * time.Second

// waitAuth 等待鉴权完成，返回 false 表示超时或连接已断开。
// 原实现是 `for !authOk { time.Sleep(10) }`：10 是 10 纳秒（未乘毫秒）的忙等，
// 且没有超时——一旦连接失败或鉴权被拒，调用方会永久卡死。
func (wsOp *WebSocketOp) waitAuth(timeout time.Duration) bool {
	deadline := time.Now().Add(timeout)
	for !wsOp.authOk {
		if wsOp.conn == nil {
			log.Error("websocket 未连接或连接已断开，订阅/请求无法发送")
			return false
		}
		if time.Now().After(deadline) {
			log.Error("websocket 鉴权等待超时（%s），请检查网络与 AccessKey/SecretKey", timeout)
			return false
		}
		time.Sleep(10 * time.Millisecond)
	}
	return true
}

func (wsOp *WebSocketOp) sub(subStr []byte, ch string, fun interface{}, param reflect.Type) bool {

	if !wsOp.waitAuth(wsAuthTimeout) {
		return false
	}
	ch = strings.ToLower(ch)

	var mi *MethonInfo = nil
	var found bool
	if mi, found = wsOp.onSubCallbackFuns[ch]; found {
		mi = new(MethonInfo).Init(fun, param)
		wsOp.onSubCallbackFuns[ch] = mi
		return true
	}
	wsOp.conn.WriteMessage(websocket.TextMessage, subStr)
	log.Info("websocket has send data: %s", subStr)

	wsOp.allSubStrs.PushBack(string(subStr))

	mi = new(MethonInfo).Init(fun, param)
	wsOp.onSubCallbackFuns[ch] = mi

	return true
}

func (wsOp *WebSocketOp) unsub(unsubStr []byte, ch string) bool {

	if !wsOp.waitAuth(wsAuthTimeout) {
		return false
	}
	ch = strings.ToLower(ch)

	if _, found := wsOp.onSubCallbackFuns[ch]; !found {
		return true
	}
	wsOp.conn.WriteMessage(websocket.TextMessage, unsubStr)
	log.Info("websocket has send data: %s", unsubStr)

	var next *list.Element
	for e := wsOp.allSubStrs.Front(); e != nil; e = next {
		next = e.Next()
		val := e.Value.(string)
		if val == string(unsubStr) {
			wsOp.allSubStrs.Remove(e)
		}
	}

	return true
}

func (wsOp *WebSocketOp) req(subStr []byte, ch string, fun interface{}, param reflect.Type) bool {

	if !wsOp.waitAuth(wsAuthTimeout) {
		return false
	}
	ch = strings.ToLower(ch)

	var mi *MethonInfo = nil
	var found bool
	if mi, found = wsOp.onReqCallbackFuns[ch]; found {
		mi = new(MethonInfo).Init(fun, param)
		wsOp.onReqCallbackFuns[ch] = mi
		return true
	}
	wsOp.conn.WriteMessage(websocket.TextMessage, subStr)
	log.Info("websocket has send data: %s", subStr)

	mi = new(MethonInfo).Init(fun, param)
	wsOp.onReqCallbackFuns[ch] = mi

	return true
}

func (wsOp *WebSocketOp) unreq(unsubStr []byte, ch string) bool {

	if !wsOp.waitAuth(wsAuthTimeout) {
		return false
	}
	ch = strings.ToLower(ch)

	if _, found := wsOp.onReqCallbackFuns[ch]; !found {
		return true
	}
	wsOp.conn.WriteMessage(websocket.TextMessage, unsubStr)
	log.Info("websocket has send data: %s", unsubStr)

	return true
}
