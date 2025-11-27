package account

type FeeDeductionCurrencyResponse struct {
	FeeOption         int    `json:"fee_option"`
	DeductionCurrency string `json:"deduction_currency"`
}
