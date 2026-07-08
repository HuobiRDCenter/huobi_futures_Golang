package market

type MarketOpenInterestResponse struct {
    Code    int                     `json:"code"`
    Data    OpenInterestData        `json:"data"`
    Message string                  `json:"message"`
    Success bool                    `json:"success"`
    Ts      int64                   `json:"ts"`
}

type OpenInterestData struct {
    ContractCode string `json:"contract_code"`
    Amount       string `json:"amount"`
    Volume       string `json:"volume"`
    Value        string `json:"value"`
    TradeAmount  string `json:"trade_amount"`
    TradeVolume  string `json:"trade_volume"`
    TradeTurnover string `json:"trade_turnover"`
}