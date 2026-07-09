package algo

type QueryAlgoOrderHistoryRequest struct {
    ContractCode string   `json:"contract_code,omitempty"`
    MarginMode   string   `json:"margin_mode,omitempty"`
    States       []string `json:"states,omitempty"`
    Type         string   `json:"type"`
    StartTime    *int64   `json:"start_time,omitempty"`
    EndTime      *int64   `json:"end_time,omitempty"`
    From         *int64   `json:"from,omitempty"`
    Limit        *int     `json:"limit,omitempty"`
    Direct       string   `json:"direct,omitempty"`
}