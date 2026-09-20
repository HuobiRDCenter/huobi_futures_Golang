package v5

import "encoding/json"

// TradeResponse V5 交易指令（wss://api.hbdm.com/linear-swap-trade）的通用应答壳。
// 5 个 op 共用：place_order / place_batch_orders / cancel_order / cancel_batch_orders / cancel_all_orders。
//
// 应答帧没有 op 字段，靠 Cid 关联请求；帧为 gzip 二进制，由 WebSocketOp.readLoop 解压。
//
// 实测帧样例（2026-09-20，参数非法时的应答）：
//
//	{"code":1067,"message":"Illegal parameter contract_code.","data":null,"args":null,
//	 "ts":1789878995020,"cid":"probe-1","rate_limit":{"limit":"72","interval":"3000",
//	 "remaining":"71","reset":"1789878998018"}}
type TradeResponse struct {
	Code int64 `json:"code"`

	Message string `json:"message"`

	// Cid 与请求帧 {"op":..,"cid":..,"data":..} 里的 cid 一致，用于关联请求
	Cid string `json:"cid"`

	Ts int64 `json:"ts"`

	// Data 随 op 而异（下单为对象数组、撤单为对象数组、失败为 null），由调用方按 op 解析
	Data json.RawMessage `json:"data"`

	Args json.RawMessage `json:"args"`

	RateLimit json.RawMessage `json:"rate_limit"`
}
