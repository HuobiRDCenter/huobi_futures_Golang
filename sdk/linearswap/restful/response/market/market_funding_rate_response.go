package market

type MarketFundingRateResponse struct {
    Code    int                         `json:"code"`
    Data    []MarketFundingRateData     `json:"data"`
    Message string                      `json:"message"`
    Success bool                        `json:"success"`
    Ts      int64                       `json:"ts"`
}

type MarketFundingRateData struct {
    ContractCode      string `json:"contract_code"`
    FundingRate       string `json:"funding_rate"`
    FundingTime       string `json:"funding_time"`
    NextFundingTime   string `json:"next_funding_time"`
    MinFundingRate    string `json:"min_funding_rate"`
    MaxFundingRate    string `json:"max_funding_rate"`
}