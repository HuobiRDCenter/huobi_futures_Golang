package market

type GetMarketMultiAssetsMarginListResponse struct {
	Code int `json:"code"`

	// V5 多资产保证金列表：data 是**对象** {"multi_assets":["BTC","ETH"]}，不是数组
	Data struct {
		MultiAssets []string `json:"multi_assets" bson:"multi_assets"`
	} `json:"data,omitempty"`

	Message string `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
