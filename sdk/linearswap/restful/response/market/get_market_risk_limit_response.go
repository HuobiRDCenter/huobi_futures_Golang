package market

type GetMarketRiskLimitResponse struct {
	Code int `json:"code"`

	Data []struct {
		Currency []string `json:"currency" bson:"currency"`
	} `json:"data,omitempty"`

	Message string `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
