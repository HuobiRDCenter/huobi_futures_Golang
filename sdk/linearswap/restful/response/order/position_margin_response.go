package order

type PositionMarginResponse struct {
	Code    int64              `json:"code"`
	Data    PositionMarginData `json:"data"`
	Message string             `json:"message"`
	Ts      int64              `json:"ts"`
}

type PositionMarginData struct {
}
