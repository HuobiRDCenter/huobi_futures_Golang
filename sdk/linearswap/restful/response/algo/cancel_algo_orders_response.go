package algo

type CancelAlgoOrdersResponse struct {
    Code    int                      `json:"code"`
    Message string                   `json:"message"`
    Ts      int64                    `json:"ts"`
    Data    []CancelAlgoOrderData    `json:"data"`
}

type CancelAlgoOrderData struct {
    AlgoId             string `json:"algo_id"`
    AlgoClientOrderId  string `json:"algo_client_order_id"`
    Success            bool   `json:"success"`
    ErrorMsg           string `json:"error_msg,omitempty"`
    ErrorCode          int    `json:"error_code,omitempty"`
}