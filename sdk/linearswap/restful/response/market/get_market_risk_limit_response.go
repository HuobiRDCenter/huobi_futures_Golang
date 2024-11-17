package market

type GetMarketRiskLimitResponse struct {
	Code string `json:"code"`

	Data []struct {
		Currency []string `json:"currency" bson:"currency"`
	} `json:"data,omitempty"`

	Message int `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
