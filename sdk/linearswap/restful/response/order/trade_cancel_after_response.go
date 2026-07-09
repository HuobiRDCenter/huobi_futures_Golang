package order

type TradeCancelAfterResponse struct {
    Code    int64                     `json:"code"`
    Data    TradeCancelAfterData      `json:"data"`
    Message string                    `json:"message"`
    Ts      int64                     `json:"ts"`
}

type TradeCancelAfterData struct {
    CurrentTime string `json:"current_time"`
    TriggerTime string `json:"trigger_time"`
}