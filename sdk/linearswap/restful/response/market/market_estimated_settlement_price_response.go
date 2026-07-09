package market

type MarketEstimatedSettlementPriceResponse struct {
    Code    int                             `json:"code"`
    Data    []EstimatedSettlementPriceData  `json:"data"`
    Message string                          `json:"message"`
    Ts      int64                           `json:"ts"`
}

type EstimatedSettlementPriceData struct {
    ContractCode             string `json:"contract_code"`
    SettlementType           string `json:"settlement_type"`
    EstimatedSettlementPrice string `json:"estimated_settlement_price"`
}