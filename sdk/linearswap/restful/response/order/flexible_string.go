package order

import (
	"bytes"
	"encoding/json"
)

// FlexibleString 兼容"同一字段线上既可能是字符串也可能是数字"的情况。
//
// 实测（2026-09-20，SPX500-USDT）：
//
//	POST /v5/trade/order     →  "order_id":"1551218163888664576"   （字符串）
//	POST /v5/trade/position  →  "order_id": 1551218221863182336   （数字）
//
// 若字段直接声明成 string，数字形态会让**整个 data 对象**解析失败，调用方连
// client_order_id 都拿不到；声明成数字则相反。故统一用本类型承接，取值时 string(x)。
type FlexibleString string

func (f *FlexibleString) UnmarshalJSON(b []byte) error {
	b = bytes.TrimSpace(b)
	if len(b) == 0 || string(b) == "null" {
		*f = ""
		return nil
	}
	if b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		*f = FlexibleString(s)
		return nil
	}
	// 数字（或 true/false）原样落成文本
	*f = FlexibleString(string(b))
	return nil
}

// String 便于直接当字符串用（fmt、拼接等）
func (f FlexibleString) String() string { return string(f) }
