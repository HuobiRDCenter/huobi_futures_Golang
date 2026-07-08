package notify

// AccountSubResponse 账户变动订阅响应
type AccountSubResponse struct {
    Op    string        `json:"op"`
    Topic string        `json:"topic"`
    Ts    int64         `json:"ts"`
    Event string        `json:"event"`
    UID   string        `json:"uid"`
    Data  []AccountData `json:"data"`
}

// AccountData 账户数据
type AccountData struct {
    State                string            `json:"state"`
    Equity               string            `json:"equity"`
    InitialMargin        string            `json:"initial_margin"`
    MaintenanceMargin    string            `json:"maintenance_margin"`
    MaintenanceMarginRate string           `json:"maintenance_margin_rate"`
    ProfitUnreal         string            `json:"profit_unreal"`
    AvailableMargin      string            `json:"available_margin"`
    VoucherValue         string            `json:"voucher_value"`
    CreatedTime          string            `json:"created_time"`
    UpdatedTime          string            `json:"updated_time"`
    Version              string            `json:"version"`
    Details              []CurrencyDetail  `json:"details"`
}

// CurrencyDetail 币种详情
type CurrencyDetail struct {
    Currency               string `json:"currency"`
    Equity                 string `json:"equity"`
    IsolatedEquity         string `json:"isolated_equity"`
    Available              string `json:"available"`
    WithdrawAvailable      string `json:"withdraw_available"`
    ProfitUnreal           string `json:"profit_unreal"`
    IsolatedProfitUnreal   string `json:"isolated_profit_unreal"`
    InitialMargin          string `json:"initial_margin"`
    MaintenanceMargin      string `json:"maintenance_margin"`
    MaintenanceMarginRate  string `json:"maintenance_margin_rate"`
    InitialMarginRate      string `json:"initial_margin_rate"`
    Voucher                string `json:"voucher"`
    VoucherValue           string `json:"voucher_value"`
    CreatedTime            string `json:"created_time"`
    UpdatedTime            string `json:"updated_time"`
}

// OnAccountSubResponse 账户变动回调函数类型
type OnAccountSubResponse func(*AccountSubResponse)