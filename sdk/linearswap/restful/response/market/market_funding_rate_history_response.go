package market

type MarketFundingRateHistoryResponse struct {
    Code    int                               `json:"code"`
    Data    []FundingRateHistoryData          `json:"data"`
    Message string                            `json:"message"`
    Success bool                              `json:"success"`
    Ts      int64                             `json:"ts"`
}

type FundingRateHistoryData struct {
    Id           string `json:"id"`
    ContractCode string `json:"contract_code"`
    FundingRate  string `json:"funding_rate"`
    FundingTime  string `json:"funding_time"`
}