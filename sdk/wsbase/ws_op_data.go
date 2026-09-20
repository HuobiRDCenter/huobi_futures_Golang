package wsbase

type WSOpData struct {
	Op string `json:"op"`

	Cid string `json:"cid,omitempty"`

	Topic string `json:"topic"`

	// ContractCode V5 订阅帧的合约代码（{"op":"sub","topic":"orders","contract_code":"BTC-USDT"}）。
	// v3 的 topic 里已内嵌合约代码，此字段为空时 omitempty 不会出现在帧里。
	ContractCode string `json:"contract_code,omitempty"`
}
