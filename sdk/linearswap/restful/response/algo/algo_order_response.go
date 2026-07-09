package algo

type AlgoOrderResponse struct {
    Code    int               `json:"code"`
    Message string            `json:"message"`
    Ts      int64             `json:"ts"`
    Data    []AlgoOrderData   `json:"data"`
}

type AlgoOrderData struct {
    AlgoId             string `json:"algo_id"`
    AlgoClientOrderId  string `json:"algo_client_order_id"`
}