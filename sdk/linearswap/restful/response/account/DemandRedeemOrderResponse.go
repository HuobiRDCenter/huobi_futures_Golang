package account

type DemandRedeemOrderResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    *RedeemDto `json:"data"`
}
type RedeemDto struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
	OrderId  int64  `json:"orderId"`
	Status   int    `json:"status"`
}
