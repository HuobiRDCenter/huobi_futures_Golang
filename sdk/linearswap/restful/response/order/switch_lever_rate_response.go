package order

type SwitchLeverRateResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		ContractCode string `json:"contract_code"`

		MarginMode string `json:"margin_mode"`

		LeverRate int `json:"lever_rate"`

		ContractType string `json:"contract_type"`

		Pair string `json:"pair"`

		BusinessType string `json:"business_type"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
