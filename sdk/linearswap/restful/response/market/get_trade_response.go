package market

type GetTradeResponse struct {
	Ch string `json:"ch,omitempty"`

	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Tick struct {
		Id int64 `json:"id"`

		Data []struct {
			Amount string `json:"amount"`

			Direction string `json:"direction"`

			Id int64 `json:"id"`

			Price string `json:"price"`

			Ts int64 `json:"ts"`

			Quantity float64 `json:"quantity"`

			TradeTurnover float64 `json:"trade_turnover"`

			ContractCode string `json:"contract_code"`

			BusinessType string `json:"business_type"`
		} `json:"data"`

		Ts int64 `json:"ts"`
	} `json:"tick,omitempty"`

	Ts int64 `json:"ts"`
}
