package account

// 执行划转请求
type UniversalTransferRequest struct {
	Currency       string `json:"currency"`
	Amount         string `json:"amount"`
	FromAccountType string `json:"from_account_type"`
	ToAccountType   string `json:"to_account_type"`
	FromAssetType   string `json:"from_asset_type,omitempty"`
	ToAssetType     string `json:"to_asset_type,omitempty"`
}