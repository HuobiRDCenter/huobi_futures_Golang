package common

type GetSwapEliteAccountRatioResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   struct {
		Symbol       string `json:"symbol"`
		ContractCode string `json:"contract_code"`
		List         []struct {
			BuyRatio    float64 `json:"buy_ratio"`
			SellRatio   float64 `json:"sell_ratio"`
			LockedRatio float64 `json:"locked_ratio"`
			Ts          int64   `json:"ts"`
		} `json:"list,omitempty"`
	} `json:"data,omitempty"`
}
