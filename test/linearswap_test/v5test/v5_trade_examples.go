package v5test

import (
	"fmt"
	"strconv"
	"time"

	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful"
	requestalgo "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/request/algo"
	responsealgo "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/response/algo"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/response/order"
)

// 策略单用 BTC-USDT：trigger_price 远低于市价（10000，市价≈11万），永不触发，0 风险。
// 口径与 C++ test/v5_restful_algo_test.hpp 一致。
const (
	v5AlgoContract     = "BTC-USDT"
	v5AlgoTriggerPrice = "10000"
)

func tradeClient() *restful.OrderClient {
	return new(restful.OrderClient).Init(accessKey(), secretKey(), host())
}

func algoClient() *restful.AlgoClient {
	return new(restful.AlgoClient).Init(accessKey(), secretKey(), host())
}

// placeFarOrder 挂一个远价限价买单（1000，市价≈7600，永不成交），返回 order_id。
// dual_side 模式下单必带 position_side。
func placeFarOrder() string {
	ch := make(chan order.SwapTradeOrderResponse, 1)
	go tradeClient().SwapTradeOrderAsync(ch, v5Contract, "cross", "long", "buy", "limit",
		"", "", v5FarPrice, "1", 0, "", "", "", "", "", "", "", "", "", false, false)
	r := <-ch
	if r.Code == 200 && r.Data.OrderId != "" {
		return r.Data.OrderId
	}
	return ""
}

// cancelOrder 清理挂单
func cancelOrder(orderId string) {
	ch := make(chan order.SwapTradeCancelOrderResponse, 1)
	go tradeClient().SwapTradeCancelOrderAsync(ch, v5Contract, orderId, "")
	<-ch
}

// -------------------------------------------------------------
// V5 trade（13 个）

func restTradeOrder() {
	name := "POST /v5/trade/order"
	ch := make(chan order.SwapTradeOrderResponse, 1)
	go tradeClient().SwapTradeOrderAsync(ch, v5Contract, "cross", "long", "buy", "limit",
		"", "", v5FarPrice, "1", 0, "", "", "", "", "", "", "", "", "", false, false)
	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
	if r.Code == 200 {
		if r.Data.OrderId == "" {
			bad(name, "code=200 但 data.order_id 为空")
			return
		}
		cancelOrder(r.Data.OrderId) // 立即撤单，不留挂单
	}
}

func restTradeBatchOrder() {
	name := "POST /v5/trade/batch_orders"
	ch := make(chan order.SwapTradeBatchOrderResponse, 1)
	go tradeClient().SwapTradeBatchOrderAsync(ch, v5Contract, "cross", "long", "buy", "limit",
		"", "", v5FarPrice, "1", 0, "", "", "", "", "", "", "", "", "", false)
	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
	if r.Code == 200 {
		if len(r.Data) == 0 || r.Data[0].OrderId == "" {
			bad(name, "code=200 但 data[0].order_id 为空（批量下单 body 应为顶层数组）")
			return
		}
		cancelOrder(r.Data[0].OrderId)
	}
}

func restTradeCancelOrder() {
	name := "POST /v5/trade/cancel_order"
	oid := placeFarOrder()
	if oid == "" {
		skip(name, "前置挂单失败，无单可撤")
		return
	}
	ch := make(chan order.SwapTradeCancelOrderResponse, 1)
	go tradeClient().SwapTradeCancelOrderAsync(ch, v5Contract, oid, "")

	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
	if r.Code == 200 && r.Data.OrderId != oid {
		bad(name, "回显 order_id 与请求不一致")
	}
}

func restTradeCancelBatchOrders() {
	name := "POST /v5/trade/cancel_batch_orders"
	oid := placeFarOrder()
	if oid == "" {
		skip(name, "前置挂单失败，无单可撤")
		return
	}
	ch := make(chan order.SwapTradeCancelBatchOrdersResponse, 1)
	go tradeClient().SwapTradeCancelBatchOrdersAsync(ch, v5Contract, oid, "")

	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
	if r.Code == 200 && (len(r.Data) == 0 || r.Data[0].OrderId != oid) {
		bad(name, "回显 order_id 与请求不一致")
	}
}

func restTradeCancelAllOrders() {
	name := "POST /v5/trade/cancel_all_orders"
	ch := make(chan order.SwapTradeAllOrdersResponse, 1)
	go tradeClient().SwapTradeAllOrdersAsync(ch, v5Contract, "buy", "long")

	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
}

// 全平：无持仓时 1048（无仓可平），接口可达即通过，口径同 C++
func restTradePosition() {
	name := "POST /v5/trade/position"
	ch := make(chan order.SwapTradePositionResponse, 1)
	go tradeClient().SwapTradePositionAsync(ch, v5Contract, "cross", "long", "")

	r := <-ch
	if r.Code == 200 || r.Code == 1048 {
		ok(name, int64(r.Code), r.Message)
		return
	}
	check(name, int64(r.Code), r.Ts, r.Message)
}

// 有仓平仓：覆盖 /v5/trade/position 的 data 为**单体对象**、且 order_id 是**数字**的线上形态
// （无仓时走上面 1048 分支，验不到这两点；SPX500 单张面值低，一次开平仅手续费成本）
func restTradePositionWithPosition() {
	name := "POST /v5/trade/position（有仓平仓）"
	oc := make(chan order.SwapTradeOrderResponse, 1)
	go tradeClient().SwapTradeOrderAsync(oc, v5Contract, "cross", "long", "buy", "market",
		"", "", "", "1", 0, "", "", "", "", "", "", "", "", "", false, false)
	o := <-oc
	if o.Code != 200 {
		skip(name, fmt.Sprintf("前置市价建仓失败 code=%d，无仓可平", o.Code))
		return
	}
	time.Sleep(time.Second)

	ch := make(chan order.SwapTradePositionResponse, 1)
	go tradeClient().SwapTradePositionAsync(ch, v5Contract, "cross", "long", "")
	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
	if r.Code == 200 && r.Data.OrderId == "" {
		bad(name, "code=200 但 data.order_id 为空（data 对象形态没解析出来？）")
	}

	// 收尾：撤干净、平干净，避免污染后续用例
	wsCancelAll()
	wsCloseAll()
}

func restTradePositionAll() {
	name := "POST /v5/trade/position_all"
	ch := make(chan order.SwapTradePositionAllResponse, 1)
	go tradeClient().SwapTradePositionAllAsync(ch)
	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
	// 有仓可平时 data 是元素自带 code/order_id 的数组，无仓时是 []（实测 2026-09-20）
	for i, d := range r.Data {
		if d.Code != 200 || d.OrderId == "" {
			bad(name, "元素 %d：code=%d order_id=%q（有仓时元素应带 code=200 与 order_id）", i, d.Code, d.OrderId)
			return
		}
	}
}

func restTradeOrderOpens() {
	name := "GET  /v5/trade/order/opens"
	oid := placeFarOrder()
	ch := make(chan order.GetTradeOpensResponse, 1)
	go acctClient().GetTradeOpensAsync(ch, v5Contract, "cross", "", "", 0, 0, "")

	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
	if oid != "" {
		cancelOrder(oid)
	}
}

func restTradeOrderDetails() {
	name := "GET  /v5/trade/order/details"
	oid := placeFarOrder()
	if oid == "" {
		skip(name, "前置挂单失败，无单可查成交")
		return
	}
	ch := make(chan order.GetTradeOrderTradesResponse, 1)
	go acctClient().GetTradeOrderTradesAsync(ch, v5Contract, "", "cross", oid, "", 0, 0, "")

	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
	cancelOrder(oid)

	// 挂单不会成交 → 再用最近一笔已成交订单查一次，覆盖 data 元素与字符串时间戳字段
	hc := make(chan order.GetTradeOrderHistoryResponse, 1)
	go acctClient().GetTradeOrderHistoryAsync(hc, v5Contract, "", "", "", "", "", 0, 0, "", "cross", "")
	h := <-hc
	if h.Code != 200 {
		return
	}
	filled := ""
	for _, d := range h.Data {
		if d.State == "filled" {
			filled = d.OrderID
			break
		}
	}
	if filled == "" {
		return
	}
	name = "GET  /v5/trade/order/details（已成交订单）"
	fc := make(chan order.GetTradeOrderTradesResponse, 1)
	go acctClient().GetTradeOrderTradesAsync(fc, v5Contract, "", "cross", filled, "", 0, 0, "")
	f := <-fc
	check(name, int64(f.Code), f.Ts, f.Message)
	if f.Code == 200 && len(f.Data) > 0 {
		d := f.Data[0]
		if d.TradeId == "" || d.CreatedTime == "" || d.UpdatedTime == "" {
			bad(name, "成交明细字段缺失：trade_id=%q created_time=%q updated_time=%q（时间戳线上是字符串）",
				d.TradeId, d.CreatedTime, d.UpdatedTime)
		}
	}
}

func restTradeOrderHistory() {
	name := "GET  /v5/trade/order/history"
	ch := make(chan order.GetTradeOrderHistoryResponse, 1)
	go acctClient().GetTradeOrderHistoryAsync(ch, v5Contract, "", "", "", "", "", 0, 0, "", "cross", "")

	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
}

// 带业务参数的私有 GET：order_id 必须进签名，否则 401
func restGetTradeOrder() {
	name := "GET  /v5/trade/order"
	oid := placeFarOrder()
	if oid == "" {
		skip(name, "前置挂单失败，无单可查")
		return
	}
	ch := make(chan order.GetTradeOrderResponse, 1)
	go acctClient().GetTradeOrderAsync(ch, v5Contract, "cross", oid, "")

	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
	if r.Code == 200 && r.Data.OrderID != oid {
		bad(name, "回显 order_id 与请求不一致")
	}
	cancelOrder(oid)
}

// 关自动撤单：on_off 是必填（空参 404），此处显式传 0 做幂等关闭
func restTradeCancelAfter() {
	name := "POST /v5/trade/cancel-after"
	ch := make(chan order.TradeCancelAfterResponse, 1)
	go acctClient().TradeCancelAfterAsync(ch, "0", "")
	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
}

// -------------------------------------------------------------
// V5 position（7 个）

func restPositionOpens() {
	name := "GET  /v5/trade/position/opens"
	ch := make(chan order.GetTradePositionOpensResponse, 1)
	go acctClient().GetTradePositionOpensAsync(ch, v5Contract)

	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
}

// position/history 不在 V5 官方清单（position 恰 7 个）内：C++ const_val.hpp 标注为
// "非V5清单范围，但C++代码使用"，线上实测返回 404 Resources do not exist。
// 方法保留以对齐 Java/C++ 代码，实测按 SKIP 处理，不计入失败。
func restPositionHistory() {
	skip("GET  /v5/trade/position/history", "不在 V5 清单内（C++ 标注非清单范围），线上 404，接口未开放")
}

func restPositionLever() {
	name := "GET  /v5/position/lever"
	ch := make(chan order.GetTradePositionLeverResponse, 1)
	go acctClient().GetTradePositionLeverAsync(ch, v5Contract, "cross")

	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
}

// 幂等设当前杠杆值，避免改动账户真实配置。
// 用 BTC-USDT + isolated：SPX500 上有刚挂的测试单会报 1493（有挂单不许调杠杆）。
func restSetPositionLever() {
	name := "POST /v5/position/lever"
	curCh := make(chan order.GetTradePositionLeverResponse, 1)
	go acctClient().GetTradePositionLeverAsync(curCh, "BTC-USDT", "isolated")

	cur := <-curCh
	if cur.Code != 200 || len(cur.Data) == 0 || cur.Data[0].LeverRate == 0 {
		skip(name, fmt.Sprintf("前置查询杠杆失败或无数据 code=%d，不做幂等写", cur.Code))
		return
	}

	ch := make(chan order.SwapTradePositionLeverResponse, 1)
	go tradeClient().SwapTradePositionLeverAsync(ch, "BTC-USDT", "isolated", strconv.FormatInt(cur.Data[0].LeverRate, 10))

	r := <-ch
	if r.Code == 200 || r.Code == 1493 {
		ok(name, int64(r.Code), r.Message)
		return
	}
	check(name, int64(r.Code), r.Ts, r.Message)
}

func restPositionMode() {
	name := "GET  /v5/position/mode"
	ch := make(chan order.GetTradePositionModeResponse, 1)
	go acctClient().GetTradePositionModeAsync(ch)
	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
}

// 幂等设当前持仓模式；有挂单/持仓时 1493 属预期
func restSetPositionMode() {
	name := "POST /v5/position/mode"
	curCh := make(chan order.GetTradePositionModeResponse, 1)
	go acctClient().GetTradePositionModeAsync(curCh)
	cur := <-curCh
	if cur.Code != 200 || cur.Data.PositionMode == "" {
		skip(name, fmt.Sprintf("前置查询持仓模式失败或无数据 code=%d，不做幂等写", cur.Code))
		return
	}

	ch := make(chan order.GetTradePositionModeResponse, 1)
	go tradeClient().SwapTradePositionModeAsync(ch, cur.Data.PositionMode)

	r := <-ch
	if r.Code == 200 || r.Code == 1493 {
		ok(name, int64(r.Code), r.Message)
		return
	}
	check(name, int64(r.Code), r.Ts, r.Message)
}

func restPositionRiskLimit() {
	name := "GET  /v5/position/risk/limit"
	ch := make(chan order.GetTradePositionRiskLimitResponse, 1)
	go acctClient().GetTradePositionRiskLimitAsync(ch, v5Contract, "cross", "")

	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
	if r.Code == 200 && len(r.Data) == 0 {
		bad(name, "code=200 但 data 为空数组")
	}
}

func restPositionRiskLimitTier() {
	name := "GET  /v5/position/risk/limit_tier"
	ch := make(chan order.PositionRiskLimitTierResponse, 1)
	go acctClient().PositionRiskLimitTierAsync(ch, v5Contract, "cross")

	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
}

// 调整逐仓保证金：无 isolated 持仓时返回业务错误属预期，
// 真实调整需先开逐仓仓，成本高且改变账户状态，故弱验证（口径同 C++）
func restPositionMargin() {
	name := "POST /v5/position/margin"
	ch := make(chan order.PositionMarginResponse, 1)
	go acctClient().PositionMarginAsync(ch, "BTC-USDT", "long", "1", "add")

	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
}

// -------------------------------------------------------------
// V5 algo（5 个）

// placeFarAlgoOrder 建一个远价限价的计划委托，返回 algo_id
func placeFarAlgoOrder() string {
	req := requestalgo.AlgoOrderRequest{
		ContractCode: v5AlgoContract,
		Type:         "trigger",
		PositionSide: "long",
		Side:         "buy",
		MarginMode:   "cross",
		Volume:       "1",
		TriggerPrice: v5AlgoTriggerPrice,
		// trigger 类型服务端强制要求 price（缺则 1066），口径同 C++ 测试传 price=trigger_price
		Price:            v5AlgoTriggerPrice,
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

func cancelAlgoOrder(algoId string) {
	// 策略单创建后有注册延迟，立即撤会返回 1061 The order does not exist（实测确认）
	time.Sleep(time.Second)
	ch := make(chan responsealgo.CancelAlgoOrdersResponse, 1)
	go algoClient().CancelAlgoOrderAsync(ch, requestalgo.CancelAlgoOrdersRequest{
		AlgoId:       algoId,
		ContractCode: v5AlgoContract,
	})
	<-ch
}

func restAlgoOrder() {
	name := "POST /v5/algo/order"
	req := requestalgo.AlgoOrderRequest{
		ContractCode: v5AlgoContract,
		Type:         "trigger",
		PositionSide: "long",
		Side:         "buy",
		MarginMode:   "cross",
		Volume:       "1",
		TriggerPrice: v5AlgoTriggerPrice,
		// trigger 类型服务端强制要求 price（缺则 1066），口径同 C++ 测试传 price=trigger_price
		Price:            v5AlgoTriggerPrice,
		TriggerPriceType: "last",
		OrderPriceType:   "optimal_20",
	}
	ch := make(chan responsealgo.AlgoOrderResponse, 1)
	go algoClient().AlgoOrderAsync(ch, req)
	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
	if r.Code == 200 {
		if len(r.Data) == 0 || r.Data[0].AlgoId == "" {
			bad(name, "code=200 但 data[0].algo_id 为空")
			return
		}
		cancelAlgoOrder(r.Data[0].AlgoId)
	}
}

func restAlgoCancelOrders() {
	name := "POST /v5/algo/cancel_orders"
	algoId := placeFarAlgoOrder()
	if algoId == "" {
		skip(name, "前置建策略单失败，无单可撤")
		return
	}
	time.Sleep(time.Second)
	ch := make(chan responsealgo.CancelAlgoOrdersResponse, 1)
	go algoClient().CancelAlgoOrderAsync(ch, requestalgo.CancelAlgoOrdersRequest{
		AlgoId:       algoId,
		ContractCode: v5AlgoContract,
	})

	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
	// 批量撤单的应答元素自带 code：外层 200 但元素非 200 时（如 1061 单不存在）仍是失败
	if r.Code == 200 && len(r.Data) > 0 && r.Data[0].Code != 200 {
		bad(name, "外层 200 但元素 code=%d message=%s", r.Data[0].Code, r.Data[0].Message)
	}
}

func restAlgoQueryOrder() {
	name := "GET  /v5/algo/order"
	algoId := placeFarAlgoOrder()
	if algoId == "" {
		skip(name, "前置建策略单失败，无单可查")
		return
	}
	ch := make(chan responsealgo.QueryAlgoOrdersResponse, 1)
	go algoClient().QueryAlgoOrderAsync(ch, requestalgo.QueryAlgoOrdersRequest{
		Type:         "trigger",
		AlgoId:       algoId,
		ContractCode: v5AlgoContract,
	})
	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
	cancelAlgoOrder(algoId)
}

func restAlgoQueryOpenOrders() {
	name := "GET  /v5/algo/order/opens"
	ch := make(chan responsealgo.QueryOpenAlgoOrdersResponse, 1)
	go algoClient().QueryOpenAlgoOrdersAsync(ch, requestalgo.QueryOpenAlgoOrdersRequest{
		Type:         "trigger",
		ContractCode: v5AlgoContract,
	})
	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
}

func restAlgoQueryOrderHistory() {
	name := "GET  /v5/algo/order/history"
	ch := make(chan responsealgo.QueryAlgoOrderHistoryResponse, 1)
	go algoClient().QueryAlgoOrderHistoryAsync(ch, requestalgo.QueryAlgoOrderHistoryRequest{
		Type:         "trigger",
		ContractCode: v5AlgoContract,
		MarginMode:   "cross",
	})
	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
}

// RunRestTradePositionAlgo 跑 V5 trade（13）、position（7）、algo（5）
func RunRestTradePositionAlgo() {
	Section("V5 REST · trade（私有，13）")
	restTradeOrder()
	restTradeBatchOrder()
	restTradeCancelOrder()
	restTradeCancelBatchOrders()
	restTradeCancelAllOrders()
	restTradePosition()
	restTradePositionWithPosition()
	restTradePositionAll()
	restTradeOrderOpens()
	restTradeOrderDetails()
	restTradeOrderHistory()
	restGetTradeOrder()
	restTradeCancelAfter()
	restPositionOpens() // #13 GET /v5/trade/position/opens 也属 trade 段

	Section("V5 REST · position（私有，7）")
	restPositionHistory()
	restPositionLever()
	restSetPositionLever()
	restPositionMode()
	restSetPositionMode()
	restPositionRiskLimit()
	restPositionRiskLimitTier()
	restPositionMargin()

	Section("V5 REST · algo（私有，5）")
	restAlgoOrder()
	restAlgoCancelOrders()
	restAlgoQueryOrder()
	restAlgoQueryOpenOrders()
	restAlgoQueryOrderHistory()
}
