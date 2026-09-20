package ws

import (
	"encoding/json"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/ws/response/v5"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/wsbase"
	"reflect"
)

// WSV5NotifyClient V5 推送客户端，对应 wss://api.hbdm.com/ws/v5/notification。
//
// 与 v3 的差别（均已实测确认）：
//   - 路径 /ws/v5/notification（v3 是 /linear-swap-notification）
//   - 订阅帧把合约代码放在独立字段：{"op":"sub","topic":"orders","contract_code":"BTC-USDT"}，
//     而不是像 v3 那样内嵌进 topic（orders.btc-usdt）
//   - 推送帧 op 为 "notify"、按 topic 字段路由，data 结构随 topic 而异
//
// 需鉴权：Init 时传入 accessKey/secretKey 即会自动发送 auth 帧（签名口径与 v3 相同）。
type WSV5NotifyClient struct {
	WebSocketOp
}

// Init 连接并鉴权。host 传空则用默认 host。
func (wsV5 *WSV5NotifyClient) Init(accessKey string, secretKey string, host string) *WSV5NotifyClient {
	if host == "" {
		host = linearswap.LINEAR_SWAP_DEFAULT_HOST
	}
	// 订阅应答（op=sub/unsub）也按 topic 路由交回调，调用方据此判断订阅就绪（见 WebSocketOp.v5Mode）。
	// 必须在 open（起读循环）之前置位，避免与读循环竞争
	wsV5.v5Mode = true
	wsV5.open("/ws/v5/notification", host, accessKey, secretKey, true)
	return wsV5
}

// OnV5NotificationResponse 7 个 topic 共用的回调类型，data 由调用方按 Topic 解析
type OnV5NotificationResponse func(*v5.NotificationResponse)

// subTopic 构造 V5 订阅帧。cid 为空时不带该字段（V5 订阅应答按 topic 回，cid 仅作透传）。
func (wsV5 *WSV5NotifyClient) subTopic(op string, topic string, contractCode string,
	callbackFun OnV5NotificationResponse, cid string) {

	opData := wsbase.WSOpData{Op: op, Topic: topic, ContractCode: contractCode}
	if cid != "" {
		opData.Cid = cid
	}
	jdata, _ := json.Marshal(opData)

	if op == "sub" {
		wsV5.sub(jdata, topic, callbackFun, reflect.TypeOf(v5.NotificationResponse{}))
	} else {
		wsV5.unsub(jdata, topic)
	}
}

// -------------------------------------------------------------
// orders

func (wsV5 *WSV5NotifyClient) SubOrders(contractCode string, callbackFun OnV5NotificationResponse, cid string) {
	wsV5.subTopic("sub", "orders", contractCode, callbackFun, cid)
}

func (wsV5 *WSV5NotifyClient) UnsubOrders(contractCode string, cid string) {
	wsV5.subTopic("unsub", "orders", contractCode, nil, cid)
}

// -------------------------------------------------------------
// trade

func (wsV5 *WSV5NotifyClient) SubTrade(contractCode string, callbackFun OnV5NotificationResponse, cid string) {
	wsV5.subTopic("sub", "trade", contractCode, callbackFun, cid)
}

func (wsV5 *WSV5NotifyClient) UnsubTrade(contractCode string, cid string) {
	wsV5.subTopic("unsub", "trade", contractCode, nil, cid)
}

// -------------------------------------------------------------
// trade_detail

func (wsV5 *WSV5NotifyClient) SubTradeDetail(contractCode string, callbackFun OnV5NotificationResponse, cid string) {
	wsV5.subTopic("sub", "trade_detail", contractCode, callbackFun, cid)
}

func (wsV5 *WSV5NotifyClient) UnsubTradeDetail(contractCode string, cid string) {
	wsV5.subTopic("unsub", "trade_detail", contractCode, nil, cid)
}

// -------------------------------------------------------------
// positions

func (wsV5 *WSV5NotifyClient) SubPositions(contractCode string, callbackFun OnV5NotificationResponse, cid string) {
	wsV5.subTopic("sub", "positions", contractCode, callbackFun, cid)
}

func (wsV5 *WSV5NotifyClient) UnsubPositions(contractCode string, cid string) {
	wsV5.subTopic("unsub", "positions", contractCode, nil, cid)
}

// -------------------------------------------------------------
// account

func (wsV5 *WSV5NotifyClient) SubAccount(contractCode string, callbackFun OnV5NotificationResponse, cid string) {
	wsV5.subTopic("sub", "account", contractCode, callbackFun, cid)
}

func (wsV5 *WSV5NotifyClient) UnsubAccount(contractCode string, cid string) {
	wsV5.subTopic("unsub", "account", contractCode, nil, cid)
}

// -------------------------------------------------------------
// match_orders

func (wsV5 *WSV5NotifyClient) SubMatchOrders(contractCode string, callbackFun OnV5NotificationResponse, cid string) {
	wsV5.subTopic("sub", "match_orders", contractCode, callbackFun, cid)
}

func (wsV5 *WSV5NotifyClient) UnsubMatchOrders(contractCode string, cid string) {
	wsV5.subTopic("unsub", "match_orders", contractCode, nil, cid)
}

// -------------------------------------------------------------
// algo_orders

func (wsV5 *WSV5NotifyClient) SubAlgoOrders(contractCode string, callbackFun OnV5NotificationResponse, cid string) {
	wsV5.subTopic("sub", "algo_orders", contractCode, callbackFun, cid)
}

func (wsV5 *WSV5NotifyClient) UnsubAlgoOrders(contractCode string, cid string) {
	wsV5.subTopic("unsub", "algo_orders", contractCode, nil, cid)
}
