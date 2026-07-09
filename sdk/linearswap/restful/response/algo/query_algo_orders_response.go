package algo

type QueryAlgoOrdersResponse struct {
    Code    int                    `json:"code"`
    Message string                 `json:"message"`
    Ts      int64                  `json:"ts"`
    Data    []QueryAlgoOrderData   `json:"data"`
}

type QueryAlgoOrderData struct {
    Id                    string `json:"id"`
    AlgoId                string `json:"algo_id"`
    AlgoClientOrderId     string `json:"algo_client_order_id"`
    ContractCode          string `json:"contract_code"`
    Volume                string `json:"volume"`
    Type                  string `json:"type"`
    State                 string `json:"state"`
    PositionSide          string `json:"position_side"`
    MarginMode            string `json:"margin_mode"`
    Side                  string `json:"side"`
    
    // 止盈相关字段
    TpTriggerPrice        string `json:"tp_trigger_price,omitempty"`
    TpOrderPrice          string `json:"tp_order_price,omitempty"`
    TpType                string `json:"tp_type,omitempty"`
    TpTriggerPriceType    string `json:"tp_trigger_price_type,omitempty"`
    
    // 止损相关字段
    SlTriggerPrice        string `json:"sl_trigger_price,omitempty"`
    SlOrderPrice          string `json:"sl_order_price,omitempty"`
    SlType                string `json:"sl_type,omitempty"`
    SlTriggerPriceType    string `json:"sl_trigger_price_type,omitempty"`
    
    // 计划委托相关字段
    Price                 string `json:"price,omitempty"`
    PriceMatch            string `json:"price_match,omitempty"`
    TriggerPrice          string `json:"trigger_price,omitempty"`
    TriggerPriceType      string `json:"trigger_price_type,omitempty"`
    
    // 跟踪委托相关字段
    ActivePrice           string `json:"active_price,omitempty"`
    OrderPriceType        string `json:"order_price_type,omitempty"`
    CallbackRate          string `json:"callback_rate,omitempty"`
    ReduceOnly            *bool  `json:"reduce_only,omitempty"`
    
    // 执行相关字段
    ActualVolume          string `json:"actual_volume,omitempty"`
    ActualPrice           string `json:"actual_price,omitempty"`
    ActualTime            string `json:"actual_time,omitempty"`
    RelationOrderId       string `json:"relation_order_id,omitempty"`
    CreatedTime           string `json:"created_time"`
    UpdatedTime           string `json:"updated_time"`
    OrderSource           string `json:"order_source"`
}