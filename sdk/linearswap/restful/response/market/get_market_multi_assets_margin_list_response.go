package market

type GetMarketMultiAssetsMarginListResponse struct {
	Code string `json:"code"`

	Data []struct {
		MultiAssets []string `json:"multi_assets" bson:"multi_assets"`
	} `json:"data,omitempty"`

	Message int `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
