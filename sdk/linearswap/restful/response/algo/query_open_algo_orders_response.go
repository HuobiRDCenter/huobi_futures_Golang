package algo

type QueryOpenAlgoOrdersResponse struct {
    Code    int                       `json:"code"`
    Message string                    `json:"message"`
    Ts      int64                     `json:"ts"`
    Data    []OpenAlgoOrderData       `json:"data"`
}

type OpenAlgoOrderData struct {
    Id                    string `json:"id"`
    Volume                string `json:"volume"`
    Type                  string `json:"type"`
    State                 string `json:"state"`
    Side                  string `json:"side"`
    AlgoId                string `json:"algo_id"`
    AlgoClientOrderId     string `json:"algo_client_order_id"`
    ContractCode          string `json:"contract_code"`
    PositionSide          string `json:"position_side"`
    MarginMode            string `json:"margin_mode"`
    CreatedTime           string `json:"created_time"`
    UpdatedTime           string `json:"updated_time"`
    OrderSource           string `json:"order_source"`
    ActivePrice           string `json:"active_price,omitempty"`
    CallbackRate          string `json:"callback_rate,omitempty"`
    ReduceOnly            *bool  `json:"reduce_only,omitempty"`
    OrderPriceType        string `json:"order_price_type,omitempty"`
    
    // 可选字段，根据不同的订单类型可能有不同的值
    TpTriggerPrice        string `json:"tp_trigger_price,omitempty"`
    TpOrderPrice          string `json:"tp_order_price,omitempty"`
    TpType                string `json:"tp_type,omitempty"`
    TpTriggerPriceType    string `json:"tp_trigger_price_type,omitempty"`
    
    SlTriggerPrice        string `json:"sl_trigger_price,omitempty"`
    SlOrderPrice          string `json:"sl_order_price,omitempty"`
    SlType                string `json:"sl_type,omitempty"`
    SlTriggerPriceType    string `json:"sl_trigger_price_type,omitempty"`
    
    Price                 string `json:"price,omitempty"`
    PriceMatch            string `json:"price_match,omitempty"`
    TriggerPrice          string `json:"trigger_price,omitempty"`
    TriggerPriceType      string `json:"trigger_price_type,omitempty"`
    
    ActualVolume          string `json:"actual_volume,omitempty"`
    ActualPrice           string `json:"actual_price,omitempty"`
    ActualTime            string `json:"actual_time,omitempty"`
    RelationOrderId       string `json:"relation_order_id,omitempty"`
}