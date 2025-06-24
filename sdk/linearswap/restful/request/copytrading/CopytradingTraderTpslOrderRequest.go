package copytrading

type CopytradingTraderTpslOrderRequest struct {
	SubPositionID  string `json:"sub_position_id"`
	TpTriggerPrice string `json:"tp_trigger_price"`
	SlTriggerPrice string `json:"sl_trigger_price"`
}
