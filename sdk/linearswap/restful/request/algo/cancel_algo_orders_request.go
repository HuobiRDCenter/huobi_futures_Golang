package algo

type CancelAlgoOrdersRequest struct {
    AlgoId             string `json:"algo_id,omitempty"`
    AlgoClientOrderId  string `json:"algo_client_order_id,omitempty"`
    ContractCode       string `json:"contract_code"`
}