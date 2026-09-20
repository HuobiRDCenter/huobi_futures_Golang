package algo

type CancelAlgoOrdersResponse struct {
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Ts      int64                 `json:"ts"`
	Data    []CancelAlgoOrderData `json:"data"`
}

// CancelAlgoOrderData 批量撤单的逐项结果。
// 线上返回形如 {"code":200,"message":"Success","algo_id":"...","algo_client_order_id":null}，
// 外层 code=200 不代表每项都成功，元素自带 code/message（如 1061 单不存在）。
type CancelAlgoOrderData struct {
	Code              int    `json:"code"`
	Message           string `json:"message"`
	AlgoId            string `json:"algo_id"`
	AlgoClientOrderId string `json:"algo_client_order_id"`
}
