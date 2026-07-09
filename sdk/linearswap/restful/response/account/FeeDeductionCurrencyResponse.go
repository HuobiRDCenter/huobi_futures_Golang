package account

type FeeDeductionCurrencyResponse struct {
	Code int `json:"code"`

	Data struct {
		FeeOption         int    `json:"fee_option"`
		DeductionCurrency string `json:"deduction_currency"`
	} `json:"data,omitempty"`

	Message string `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
