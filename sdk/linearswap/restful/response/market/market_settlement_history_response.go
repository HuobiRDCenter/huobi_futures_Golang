package market

type MarketSettlementHistoryResponse struct {
    Code    int                     `json:"code"`
    Data    []SettlementHistoryData `json:"data"`
    Message string                  `json:"message"`
    Ts      int64                   `json:"ts"`
}

type SettlementHistoryData struct {
    Id              string `json:"id"`
    ContractCode    string `json:"contract_code"`
    SettlementTime  string `json:"settlement_time"`
    ClawbackRatio   string `json:"clawback_ratio"`
    SettlementPrice string `json:"settlement_price"`
}