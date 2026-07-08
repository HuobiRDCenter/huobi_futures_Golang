package algo

type AlgoOrderRequest struct {
    ContractCode string `json:"contract_code"`
    Type         string `json:"type"`
    PositionSide string `json:"position_side"`
    Side         string `json:"side"`
    MarginMode   string `json:"margin_mode"`
    
    // 可选字段
    AlgoClientOrderId   string  `json:"algo_client_order_id,omitempty"`
    Volume              string  `json:"volume,omitempty"`
    
    // 止盈相关参数
    TpTriggerPrice      string `json:"tp_trigger_price,omitempty"`
    TpOrderPrice        string `json:"tp_order_price,omitempty"`
    TpType              string `json:"tp_type,omitempty"`
    TpTriggerPriceType  string `json:"tp_trigger_price_type,omitempty"`
    
    // 止损相关参数
    SlTriggerPrice      string `json:"sl_trigger_price,omitempty"`
    SlOrderPrice        string `json:"sl_order_price,omitempty"`
    SlType              string `json:"sl_type,omitempty"`
    SlTriggerPriceType  string `json:"sl_trigger_price_type,omitempty"`
    
    // 计划委托相关参数
    Price               string `json:"price,omitempty"`
    PriceMatch          string `json:"price_match,omitempty"`
    TriggerPrice        string `json:"trigger_price,omitempty"`
    TriggerPriceType    string `json:"trigger_price_type,omitempty"`
    
    // 跟踪委托相关参数
    ActivePrice         string `json:"active_price,omitempty"`
    OrderPriceType      string `json:"order_price_type,omitempty"`
    CallbackRate        string `json:"callback_rate,omitempty"`
    ReduceOnly          *bool  `json:"reduce_only,omitempty"`
}