package market

type MarketElitePositionRatioResponse struct {
    Code    int                     `json:"code"`
    Data    []ElitePositionRatioData `json:"data"`
    Message string                  `json:"message"`
    Ts      int64                   `json:"ts"`
}

type ElitePositionRatioData struct {
    ContractCode string `json:"contract_code"`
    BuyRatio     string `json:"buy_ratio"`
    SellRatio    string `json:"sell_ratio"`
    Ts           string `json:"ts"`
}