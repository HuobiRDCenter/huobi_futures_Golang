package algo

type QueryAlgoOrdersRequest struct {
    AlgoId             string `json:"algo_id,omitempty"`
    AlgoClientOrderId  string `json:"algo_client_order_id,omitempty"`
    Type               string `json:"type"`
    ContractCode       string `json:"contract_code,omitempty"`
}