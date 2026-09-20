package v5

import (
	"bytes"
	"encoding/json"
)

// NotifyData V5 推送的 data 字段。线上两种形态都有（2026-09-20 抓原始帧确认）：
//   - 数组：trade / trade_detail / positions / match_orders / algo_orders
//   - 单体对象：orders（订单状态变更）、account（资产变更）
//
// 统一归一成数组，调用方一律按 Data[0] 取单体；用 []json.RawMessage 直接接会把
// 对象形态的帧整帧解析失败（Op/Topic 能解出、Data 为空），推送就静默丢了。
type NotifyData []json.RawMessage

func (d *NotifyData) UnmarshalJSON(b []byte) error {
	trimmed := bytes.TrimSpace(b)
	if len(trimmed) == 0 || string(trimmed) == "null" {
		*d = nil
		return nil
	}
	if trimmed[0] == '{' { // 单体对象：包成单元素数组
		var one json.RawMessage
		if err := json.Unmarshal(trimmed, &one); err != nil {
			return err
		}
		*d = NotifyData{one}
		return nil
	}
	var arr []json.RawMessage
	if err := json.Unmarshal(trimmed, &arr); err != nil {
		return err
	}
	*d = NotifyData(arr)
	return nil
}

// NotificationResponse V5 推送（wss://api.hbdm.com/ws/v5/notification）的通用应答壳。
// 7 个 topic 共用：orders / trade / trade_detail / positions / account / match_orders / algo_orders。
// data 的具体结构随 topic 而异，由调用方按 Topic 自行解析（与 Java V5NotificationResponse、
// C++ WSSwapNotificationClient 的原始串回调口径一致）。
//
// 实测帧样例（2026-09-20）：
//
//	{"op":"notify","topic":"positions","contract_code":"BTC-USDT","ts":1789878955727,
//	 "uid":"502061937","event":"snapshot","data":[]}                        // data 为数组
//	{"op":"notify","topic":"account","contract_code":"","ts":1789880764508,
//	 "uid":"502061937","event":"create_order","data":{"equity":"0",...}}    // data 为单体对象
type NotificationResponse struct {
	Op string `json:"op"`

	// Topic 推送主题，路由回调即按此字段
	Topic string `json:"topic"`

	// ContractCode 订阅时传入的合约代码
	ContractCode string `json:"contract_code"`

	Ts  int64  `json:"ts"`
	Uid string `json:"uid"`

	// Event snapshot（订阅后的首帧全量）/ update（增量）
	Event string `json:"event"`

	// ErrCode 仅订阅/鉴权应答帧携带，推送帧为 0
	ErrCode int64 `json:"err-code"`

	// Data 各 topic 的推送体，数组或单体对象都归一成数组（见 NotifyData）
	Data NotifyData `json:"data"`
}
