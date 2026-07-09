package market

type MarketLiquidationOrdersResponse struct {
    Code    int                     `json:"code"`
    Data    []LiquidationOrderData  `json:"data"`
    Message string                  `json:"message"`
    Ts      int64                   `json:"ts"`
}

type LiquidationOrderData struct {
    Id              string `json:"id"`
    ContractCode    string `json:"contract_code"`
    LiquidationTime string `json:"liquidation_time"`
    Side            string `json:"side"`
    PositionSide    string `json:"position_side"`
    Volume          string `json:"volume"`
    Amount          string `json:"amount"`
    BankruptPrice   string `json:"bankrupt_price"`
    TradeTurnover   string `json:"trade_turnover"`
}