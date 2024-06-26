package common

type GetSwapHistoricalFundingRateResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   struct {
		TotalPage   int `json:"total_page"`
		CurrentPage int `json:"current_page"`
		TotalSize   int `json:"total_size"`
		Data        []struct {
			Symbol          string `json:"symbol"`
			ContractCode    string `json:"contract_code"`
			FeeAsset        string `json:"fee_asset"`
			FundingTime     string `json:"funding_time"`
			FundingRate     string `json:"funding_rate"`
			RealizedRate    string `json:"realized_rate,omitempty"`
			AvgPremiumIndex string `json:"avg_premium_index"`
		} `json:"data,omitempty"`
	} `json:"data,omitempty"`
}
