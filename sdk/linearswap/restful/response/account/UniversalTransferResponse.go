package account

// 执行划转响应
type UniversalTransferResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Ts      int64  `json:"ts"`
	Data    int64  `json:"data"` // 返回划转ID
}