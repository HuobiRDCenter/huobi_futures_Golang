package algo

type QueryOpenAlgoOrdersRequest struct {
    ContractCode       string `json:"contract_code,omitempty"`
    AlgoId             string `json:"algo_id,omitempty"`
    AlgoClientOrderId  string `json:"algo_client_order_id,omitempty"`
    Type               string `json:"type"`
    From               *int64 `json:"from,omitempty"`
    Limit              *int   `json:"limit,omitempty"`
    Direct             string `json:"direct,omitempty"`
}