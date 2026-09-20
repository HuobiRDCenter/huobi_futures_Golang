package v5test

import (
	"encoding/json"
	"fmt"
	"time"

	requestalgo "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/request/algo"
	responsealgo "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/response/algo"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/response/order"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/ws"
	v5 "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/ws/response/v5"
)

// WS 交易指令的挂单价：7300 在 SPX500 的价格带内（low_limit≈7271），
// 会被正常受理并返回 order_id，随即撤掉。口径同 C++ test/v5_ws_trade_test.hpp。
const v5WsFarPrice = "7300"

// -------------------------------------------------------------
// 触发器（REST 侧制造推送）

// triggerLimitOrder 挂远价限价单，返回 order_id
func triggerLimitOrder() string {
	ch := make(chan order.SwapTradeOrderResponse, 1)
	go tradeClient().SwapTradeOrderAsync(ch, v5Contract, "cross", "long", "buy", "limit",
		"", "", v5WsFarPrice, "1", 0, "", "", "", "", "", "", "", "", "", false, false)
	r := <-ch
	if r.Code == 200 && r.Data.OrderId != "" {
		return r.Data.OrderId
	}
	return ""
}

// wsCancelAll 兜底撤掉该合约所有买单
func wsCancelAll() {
	ch := make(chan order.SwapTradeAllOrdersResponse, 1)
	go tradeClient().SwapTradeAllOrdersAsync(ch, v5Contract, "buy", "long")
	<-ch
}

// wsCloseAll 兜底全平
func wsCloseAll() {
	ch := make(chan order.SwapTradePositionAllResponse, 1)
	go tradeClient().SwapTradePositionAllAsync(ch)
	<-ch
}

// triggerMarketFillAndClose 市价买 1 张成交（触发 trade / trade_detail / positions /
// account / match_orders 推送），随后全平清理。单张 SPX500 名义价值低，损耗仅手续费。
func triggerMarketFillAndClose() {
	wsCancelAll()
	wsCloseAll()

	ch := make(chan order.SwapTradeOrderResponse, 1)
	go tradeClient().SwapTradeOrderAsync(ch, v5Contract, "cross", "long", "buy", "market",
		"", "", "", "1", 0, "", "", "", "", "", "", "", "", "", false, false)
	r := <-ch
	if r.Code != 200 {
		return
	}
	time.Sleep(800 * time.Millisecond)

	closeCh := make(chan order.SwapTradePositionResponse, 1)
	go tradeClient().SwapTradePositionAsync(closeCh, v5Contract, "cross", "long", "")
	<-closeCh
	time.Sleep(300 * time.Millisecond)

	wsCloseAll()
	wsCancelAll()
}

// triggerAlgoOrder 建一个永不触发的策略单（trigger_price=100），返回 algo_id。
// price 必填：type=trigger 时即使 order_price_type=optimal_20，缺 price 也会 1066
// （The price field cannot be empty），建单失败就没有推送。
func triggerAlgoOrder() string {
	req := requestalgo.AlgoOrderRequest{
		ContractCode:     v5Contract,
		Type:             "trigger",
		PositionSide:     "long",
		Side:             "buy",
		MarginMode:       "cross",
		Volume:           "1",
		TriggerPrice:     "100",
		Price:            "100",
		TriggerPriceType: "last",
		OrderPriceType:   "optimal_20",
	}
	ch := make(chan responsealgo.AlgoOrderResponse, 1)
	go algoClient().AlgoOrderAsync(ch, req)
	r := <-ch
	if r.Code == 200 && len(r.Data) > 0 {
		return r.Data[0].AlgoId
	}
	return ""
}

func cancelAlgoById(algoId string) {
	ch := make(chan responsealgo.CancelAlgoOrdersResponse, 1)
	go algoClient().CancelAlgoOrderAsync(ch, requestalgo.CancelAlgoOrdersRequest{
		AlgoId:       algoId,
		ContractCode: v5Contract,
	})
	<-ch
}

// -------------------------------------------------------------
// 推送 topic 的必填字段（口径同 C++ checkNotifyData）

var wsTopicFields = map[string]struct {
	base   []string
	filled []string
}{
	"orders": {
		base: []string{"order_id", "contract_code", "side", "type", "volume", "state",
			"margin_mode", "position_side", "lever_rate", "created_time", "updated_time"},
		filled: []string{"trade_avg_price", "trade_volume", "trade_turnover"},
	},
	"trade": {
		base: []string{"trade_id", "order_id", "contract_code", "direction", "trade_price",
			"trade_volume", "trade_turnover", "role", "position_side", "created_time"},
	},
	"trade_detail": {
		base: []string{"trade_id", "order_id", "contract_code", "direction", "trade_price",
			"trade_volume", "trade_turnover", "trade_fee", "fee_currency", "position_side", "created_time"},
	},
	"positions": {
		base: []string{"contract_code", "position_side", "direction", "margin_mode", "volume",
			"lever_rate", "open_avg_price", "margin", "state", "created_time"},
	},
	"account": {
		base: []string{"equity", "state", "details", "initial_margin", "maintenance_margin", "available_margin"},
	},
	"match_orders": {
		base: []string{"order_id", "contract_code", "side", "type", "volume", "state",
			"match_time", "position_side"},
		filled: []string{"trade_id", "trade_price", "trade_volume", "role"},
	},
	"algo_orders": {
		base: []string{"algo_id", "contract_code", "type", "side", "volume", "state",
			"margin_mode", "position_side", "trigger_price", "created_time"},
	},
}

// runNotifyCase 订阅 → 等 sub ack → 触发 → 等带数据的 notify 帧 → 校验必填字段
func runNotifyCase(topic string, sub func(*ws.WSV5NotifyClient, ws.OnV5NotificationResponse),
	trigger func()) {

	name := "WS   sub " + topic
	spec := wsTopicFields[topic]

	subAck := make(chan struct{}, 1)
	dataCh := make(chan *v5.NotificationResponse, 32)

	cb := func(r *v5.NotificationResponse) {
		if r.Op == "sub" && r.ErrCode == 0 {
			select {
			case subAck <- struct{}{}:
			default:
			}
			return
		}
		if r.Op != "notify" || r.Topic != topic || len(r.Data) == 0 {
			return
		}
		select {
		case dataCh <- r:
		default:
		}
	}

	client := new(ws.WSV5NotifyClient).Init(accessKey(), secretKey(), host())
	sub(client, cb)

	select {
	case <-subAck:
	case <-time.After(10 * time.Second):
		bad(name, "10s 内未收到订阅应答（op=sub, err-code=0）")
		return
	}

	trigger()

	deadline := time.After(20 * time.Second)
	var lastErr string
	for {
		select {
		case r := <-dataCh:
			var m map[string]interface{}
			if err := json.Unmarshal(r.Data[0], &m); err != nil {
				lastErr = fmt.Sprintf("data 解析失败：%v", err)
				continue
			}
			if err := checkNotifyFields(topic, m, spec.base, spec.filled); err != "" {
				lastErr = err
				continue
			}
			ok(name, 200, fmt.Sprintf("event=%s fields=%d", r.Event, len(m)))
			return
		case <-deadline:
			if lastErr != "" {
				bad(name, "20s 内无合格数据帧，最后一次校验失败：%s", lastErr)
			} else {
				bad(name, "20s 内未收到 %s 的 notify 数据帧", topic)
			}
			return
		}
	}
}

// checkNotifyFields 校验必填字段，返回空串表示通过
func checkNotifyFields(topic string, m map[string]interface{}, base, filled []string) string {
	state, _ := m["state"].(string)
	required := base
	if state == "filled" {
		required = append(append([]string{}, base...), filled...)
	}
	for _, f := range required {
		v, exists := m[f]
		if !exists || v == nil {
			return fmt.Sprintf("缺少字段 %s（state=%s）", f, state)
		}
	}
	return ""
}

// -------------------------------------------------------------
// WS 交易指令（5 个）：发指令 → 等 cid 应答 → 校验 code/data

// runTradeOp 发指令并等对应 cid 的应答，返回应答帧
func runTradeOp(name string, send func(*ws.WSV5TradeClient) string) (*v5.TradeResponse, string) {
	respCh := make(chan *v5.TradeResponse, 8)
	var wantCid string

	client := new(ws.WSV5TradeClient).Init(accessKey(), secretKey(), host(), func(r *v5.TradeResponse) {
		if r.Cid != wantCid {
			return
		}
		select {
		case respCh <- r:
		default:
		}
	})

	wantCid = send(client)

	select {
	case r := <-respCh:
		return r, wantCid
	case <-time.After(15 * time.Second):
		bad(name, "15s 内未收到 cid=%s 的应答帧", wantCid)
		return nil, wantCid
	}
}

func wsTradePlaceOrder() {
	name := "WS   place_order"
	data := fmt.Sprintf(`{"contract_code":"%s","margin_mode":"cross","position_side":"long",`+
		`"side":"buy","type":"limit","price":"%s","volume":"1"}`, v5Contract, v5WsFarPrice)
	r, _ := runTradeOp(name, func(c *ws.WSV5TradeClient) string { return c.PlaceOrder(data) })
	if r == nil {
		return
	}
	check(name, r.Code, r.Ts, r.Message)
	if r.Code == 200 {
		if oid := extractOrderId(r.Data); oid == "" {
			bad(name, "code=200 但应答 data 里没有 order_id：%s", truncate(string(r.Data), 120))
			return
		} else {
			cancelOrder(oid) // 立即撤单，不留挂单
		}
	}
}

func wsTradePlaceBatchOrders() {
	name := "WS   place_batch_orders"
	data := fmt.Sprintf(`[{"contract_code":"%s","margin_mode":"cross","position_side":"long",`+
		`"side":"buy","type":"limit","price":"%s","volume":"1"}]`, v5Contract, v5WsFarPrice)
	r, _ := runTradeOp(name, func(c *ws.WSV5TradeClient) string { return c.PlaceBatchOrders(data) })
	if r == nil {
		return
	}
	check(name, r.Code, r.Ts, r.Message)
	if r.Code == 200 {
		if oid := extractOrderId(r.Data); oid == "" {
			bad(name, "code=200 但应答 data 里没有 order_id：%s", truncate(string(r.Data), 120))
			return
		} else {
			cancelOrder(oid)
		}
	}
}

func wsTradeCancelOrder() {
	name := "WS   cancel_order"
	oid := triggerLimitOrder()
	if oid == "" {
		skip(name, "前置挂单失败，无单可撤")
		return
	}
	data := fmt.Sprintf(`{"contract_code":"%s","order_id":"%s"}`, v5Contract, oid)
	r, _ := runTradeOp(name, func(c *ws.WSV5TradeClient) string { return c.CancelOrder(data) })
	if r == nil {
		return
	}
	check(name, r.Code, r.Ts, r.Message)
	if r.Code == 200 && extractOrderId(r.Data) != oid {
		bad(name, "回显 order_id 与请求不一致")
	}
}

func wsTradeCancelBatchOrders() {
	name := "WS   cancel_batch_orders"
	oid := triggerLimitOrder()
	if oid == "" {
		skip(name, "前置挂单失败，无单可撤")
		return
	}
	data := fmt.Sprintf(`{"contract_code":"%s","order_id":["%s"]}`, v5Contract, oid)
	r, _ := runTradeOp(name, func(c *ws.WSV5TradeClient) string { return c.CancelBatchOrders(data) })
	if r == nil {
		return
	}
	check(name, r.Code, r.Ts, r.Message)
	if r.Code == 200 && extractOrderId(r.Data) != oid {
		bad(name, "回显 order_id 与请求不一致")
	}
}

func wsTradeCancelAllOrders() {
	name := "WS   cancel_all_orders"
	triggerLimitOrder() // 先挂一单，保证有东西可撤
	data := fmt.Sprintf(`{"contract_code":"%s","side":"buy","position_side":"long"}`, v5Contract)
	r, _ := runTradeOp(name, func(c *ws.WSV5TradeClient) string { return c.CancelAllOrders(data) })
	if r == nil {
		return
	}
	check(name, r.Code, r.Ts, r.Message)
}

// extractOrderId 从应答 data 里取 order_id。V5 不同 op 的 data 可能是对象或数组，
// 批量接口还可能是 {"orders_data":[...]}，三种都兼容。
func extractOrderId(raw json.RawMessage) string {
	if len(raw) == 0 {
		return ""
	}
	var v interface{}
	if err := json.Unmarshal(raw, &v); err != nil {
		return ""
	}
	return findOrderId(v)
}

func findOrderId(v interface{}) string {
	switch t := v.(type) {
	case map[string]interface{}:
		if s, ok := t["order_id"].(string); ok && s != "" {
			return s
		}
		for _, key := range []string{"orders_data", "data"} {
			if inner, ok := t[key]; ok {
				if s := findOrderId(inner); s != "" {
					return s
				}
			}
		}
	case []interface{}:
		for _, e := range t {
			if s := findOrderId(e); s != "" {
				return s
			}
		}
	}
	return ""
}

// RunWs 跑 V5 WS 推送 7 个 topic + 交易指令 5 个 op
func RunWs() {
	Section("V5 WS · 推送 /ws/v5/notification（私有，7 topic）")

	runNotifyCase("orders", func(c *ws.WSV5NotifyClient, cb ws.OnV5NotificationResponse) {
		c.SubOrders(v5Contract, cb, "")
	}, func() {
		oid := triggerLimitOrder()
		time.Sleep(500 * time.Millisecond)
		if oid != "" {
			cancelOrder(oid)
		}
	})

	runNotifyCase("trade", func(c *ws.WSV5NotifyClient, cb ws.OnV5NotificationResponse) {
		c.SubTrade(v5Contract, cb, "")
	}, triggerMarketFillAndClose)

	runNotifyCase("trade_detail", func(c *ws.WSV5NotifyClient, cb ws.OnV5NotificationResponse) {
		c.SubTradeDetail(v5Contract, cb, "")
	}, triggerMarketFillAndClose)

	runNotifyCase("positions", func(c *ws.WSV5NotifyClient, cb ws.OnV5NotificationResponse) {
		c.SubPositions(v5Contract, cb, "")
	}, triggerMarketFillAndClose)

	runNotifyCase("account", func(c *ws.WSV5NotifyClient, cb ws.OnV5NotificationResponse) {
		c.SubAccount(v5Contract, cb, "")
	}, triggerMarketFillAndClose)

	runNotifyCase("match_orders", func(c *ws.WSV5NotifyClient, cb ws.OnV5NotificationResponse) {
		c.SubMatchOrders(v5Contract, cb, "")
	}, triggerMarketFillAndClose)

	runNotifyCase("algo_orders", func(c *ws.WSV5NotifyClient, cb ws.OnV5NotificationResponse) {
		c.SubAlgoOrders(v5Contract, cb, "")
	}, func() {
		algoId := triggerAlgoOrder()
		if algoId != "" {
			time.Sleep(500 * time.Millisecond)
			cancelAlgoById(algoId)
		}
	})

	Section("V5 WS · 交易指令 /linear-swap-trade（私有，5 op）")
	wsTradePlaceOrder()
	wsTradePlaceBatchOrders()
	wsTradeCancelOrder()
	wsTradeCancelBatchOrders()
	wsTradeCancelAllOrders()
}
