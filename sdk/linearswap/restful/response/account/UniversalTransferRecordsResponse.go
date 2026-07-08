package account

// 查询划转记录响应
type UniversalTransferRecordsResponse struct {
	Code    int                          `json:"code"`
	Msg     string                       `json:"msg"`
	Ts      int64                        `json:"ts"`
	Data    []UniversalTransferRecord    `json:"data"`
}

// 划转记录详情
type UniversalTransferRecord struct {
	Id              int64   `json:"id"`
	TransferId      int64   `json:"transfer_id"`
	Amount          string  `json:"amount"`
	Currency        string  `json:"currency"`
	Status          string  `json:"status"`           // success, pending, failed
	FromAccountType string  `json:"from_account_type"` // spot, linear-swap, futures, etc.
	ToAccountType   string  `json:"to_account_type"`
	FromAssetType   string  `json:"from_asset_type,omitempty"`
	ToAssetType     string  `json:"to_asset_type,omitempty"`
	Ts              int64   `json:"ts"`
}