package account

type AccountBillsResponse struct {
    Code    int64               `json:"code"`
    Data    []AccountBillsData  `json:"data"`
    Message string              `json:"message"`
    Ts      int64               `json:"ts"`
}

type AccountBillsData struct {
    Id           string `json:"id"`
    ContractCode string `json:"contract_code"`
    MarginMode   string `json:"margin_mode"`
    Type         string `json:"type"`
    Currency     string `json:"currency"`
    Amount       string `json:"amount"`
    CreatedTime  string `json:"created_time"`
}