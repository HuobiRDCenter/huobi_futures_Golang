package account

type GetValidLeverRateResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		ContractCode string `json:"contract_code"`

		MarginMode string `json:"margin_mode"`

		AvailableLeverRate string `json:"available_level_rate"`

		Pair string `json:"pair"`

		BusinessType string `json:"business_type"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
