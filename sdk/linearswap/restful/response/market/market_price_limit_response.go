package market

type MarketPriceLimitResponse struct {
    Code    int                     `json:"code"`
    Data    []PriceLimitData        `json:"data"`
    Message string                  `json:"message"`
    Ts      int64                   `json:"ts"`
}

type PriceLimitData struct {
    ContractCode string `json:"contract_code"`
    HighLimit    string `json:"high_limit"`
    LowLimit     string `json:"low_limit"`
}