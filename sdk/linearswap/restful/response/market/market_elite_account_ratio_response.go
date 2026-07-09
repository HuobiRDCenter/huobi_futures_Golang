package market

type MarketEliteAccountRatioResponse struct {
    Code    int                     `json:"code"`
    Data    []EliteAccountRatioData `json:"data"`
    Message string                  `json:"message"`
    Ts      int64                   `json:"ts"`
}

type EliteAccountRatioData struct {
    ContractCode string `json:"contract_code"`
    BuyRatio     string `json:"buy_ratio"`
    SellRatio    string `json:"sell_ratio"`
    Ts           string `json:"ts"`
}