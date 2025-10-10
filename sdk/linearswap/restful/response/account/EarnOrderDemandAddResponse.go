package account

type EarnOrderDemandAddResponse struct {
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Data    *OrderSubscribeResDTO `json:"data"`
}
type OrderSubscribeResDTO struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
	OrderId  int64  `json:"orderId"`
	Status   int    `json:"status"`
}
