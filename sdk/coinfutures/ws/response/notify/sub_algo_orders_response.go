package notify

// AlgoOrdersSubResponse 算法订单订阅响应
type AlgoOrdersSubResponse struct {
    Op            string `json:"op"`
    Topic         string `json:"topic"`
    Ts            int64  `json:"ts"`
    UID           string `json:"uid"`
    ContractCode  string `json:"contract_code"`
    Data          []AlgoOrderData `json:"data"`
}

// AlgoOrderData 算法订单数据
type AlgoOrderData struct {
    ID                   string `json:"id"`
    AlgoID               string `json:"algo_id"`
    AlgoClientOrderID    string `json:"algo_client_order_id"`
    ContractCode         string `json:"contract_code"`
    Volume               string `json:"volume"`
    Type                 string `json:"type"`
    State                string `json:"state"`
    PositionSide         string `json:"position_side"`
    MarginMode           string `json:"margin_mode"`
    Side                 string `json:"side"`
    TPTriggerPrice       string `json:"tp_trigger_price"`
    TPOrderPrice         string `json:"tp_order_price"`
    TPType               string `json:"tp_type"`
    TPTriggerPriceType   string `json:"tp_trigger_price_type"`
    SLTriggerPrice       string `json:"sl_trigger_price"`
    SLOrderPrice         string `json:"sl_order_price"`
    SLType               string `json:"sl_type"`
    SLTriggerPriceType   string `json:"sl_trigger_price_type"`
    Price                string `json:"price"`
    PriceMatch           string `json:"price_match"`
    TriggerPrice         string `json:"trigger_price"`
    TriggerPriceType     string `json:"trigger_price_type"`
    ActivePrice          string `json:"active_price"`
    OrderPriceType       string `json:"order_price_type"`
    CallbackRate         string `json:"callback_rate"`
    ReduceOnly           bool   `json:"reduce_only"`
    ActualVolume         string `json:"actual_volume"`
    ActualPrice          string `json:"actual_price"`
    ActualTime           string `json:"actual_time"`
    RelationOrderID      string `json:"relation_order_id"`
    CreatedTime          string `json:"created_time"`
    UpdatedTime          string `json:"updated_time"`
    OrderSource          string `json:"order_source"`
}

// OnAlgoOrdersSubResponse 算法订单回调函数类型
type OnAlgoOrdersSubResponse func(*AlgoOrdersSubResponse)