package ws

import (
	"fmt"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/ws/response/v5"
	"reflect"
	"strconv"
	"sync/atomic"
)

// WSV5TradeClient V5 交易指令客户端，对应 wss://api.hbdm.com/linear-swap-trade。
//
// 与推送客户端（/ws/v5/notification）的差别：
//   - 5 个指令 op：place_order / place_batch_orders / cancel_order / cancel_batch_orders / cancel_all_orders
//   - 请求帧 {"op":"<op>","cid":"<n>","data":{...}}，cid 由本客户端自动递增生成并作为各方法的返回值，
//     调用方用它与应答里的 cid 关联
//   - 应答帧**没有 op 字段**，靠 cid 路由；帧为 gzip 二进制，已由 WebSocketOp.readLoop 解压
//
// 需鉴权：Init 时传入 accessKey/secretKey。data 传原始 JSON（对象或数组），与 C++ 版一致。
type WSV5TradeClient struct {
	WebSocketOp

	cidSeq      int64
	callbackFun OnV5TradeResponse
}

// OnV5TradeResponse 5 个 op 共用的回调类型，用应答里的 Cid 关联到具体请求
type OnV5TradeResponse func(*v5.TradeResponse)

// Init 连接并鉴权。host 传空则用默认 host。
func (wsTd *WSV5TradeClient) Init(accessKey string, secretKey string, host string,
	callbackFun OnV5TradeResponse) *WSV5TradeClient {

	if host == "" {
		host = linearswap.LINEAR_SWAP_DEFAULT_HOST
	}
	wsTd.callbackFun = callbackFun
	wsTd.open("/linear-swap-trade", host, accessKey, secretKey, true)
	return wsTd
}

// sendOp 发送指令帧并返回本次 cid
func (wsTd *WSV5TradeClient) sendOp(op string, data string) string {
	if data == "" {
		data = "{}"
	}

	cid := strconv.FormatInt(atomic.AddInt64(&wsTd.cidSeq, 1), 10)
	frame := fmt.Sprintf("{\"op\":\"%s\",\"cid\":\"%s\",\"data\":%s}", op, cid, data)

	wsTd.req([]byte(frame), cid, wsTd.callbackFun, reflect.TypeOf(v5.TradeResponse{}))
	return cid
}

// PlaceOrder 下单。data 为下单参数对象，例：
// {"contract_code":"BTC-USDT","margin_mode":"cross","position_side":"long",
//
//	"side":"buy","type":"limit","price":"1","volume":"1"}
//
// 返回本次请求的 cid
func (wsTd *WSV5TradeClient) PlaceOrder(data string) string {
	return wsTd.sendOp("place_order", data)
}

// PlaceBatchOrders 批量下单。data 形如 {"orders_data":[{...},{...}]}
func (wsTd *WSV5TradeClient) PlaceBatchOrders(data string) string {
	return wsTd.sendOp("place_batch_orders", data)
}

// CancelOrder 撤单。data 形如 {"contract_code":"BTC-USDT","order_id":"123"}
func (wsTd *WSV5TradeClient) CancelOrder(data string) string {
	return wsTd.sendOp("cancel_order", data)
}

// CancelBatchOrders 批量撤单。data 形如 {"contract_code":"BTC-USDT","order_id":["1","2"]}
func (wsTd *WSV5TradeClient) CancelBatchOrders(data string) string {
	return wsTd.sendOp("cancel_batch_orders", data)
}

// CancelAllOrders 全撤。data 形如 {"contract_code":"BTC-USDT"}
func (wsTd *WSV5TradeClient) CancelAllOrders(data string) string {
	return wsTd.sendOp("cancel_all_orders", data)
}
