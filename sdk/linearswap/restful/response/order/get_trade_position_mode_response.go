package order

type GetTradePositionModeResponse struct {
	Code string `json:"code"`

	Data []struct {
		PositionMode string `json:"position_mode"`
	} `json:"data,omitempty"`

	Message int `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
