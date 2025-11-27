package account

type QueryEarnProjectListResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    *PageInfoResVo `json:"data,omitempty"`
}
type PageInfoResVo struct {
	Total int                `json:"total"`
	Items []SavingProjectDTO `json:"items,omitempty"`
}

type SavingProjectDTO struct {
	ProjectId         int64          `json:"projectId"`
	ProductId         int64          `json:"productId"`
	CalculationType   int            `json:"calculationType"`
	Type              int            `json:"type"`
	ViewYearRate      float64        `json:"viewYearRate"`
	FinishAmount      float64        `json:"finishAmount"`
	ProjectStatus     int            `json:"projectStatus"`
	TotalAmount       float64        `json:"totalAmount"`
	Currency          string         `json:"currency"`
	StartAmount       float64        `json:"startAmount"`
	ApyType           int            `json:"apyType"`
	TieredRates       []TieredRateVo `json:"tieredRates,omitempty"`
	MarketPerkUpLimit *string        `json:"marketPerkUpLimit,omitempty"`
	MarketTimeApy     *float64       `json:"marketTimeApy,omitempty"`
	MarketPerkApy     *float64       `json:"marketPerkApy,omitempty"`
}

type TieredRateVo struct {
	AmountStart float64 `json:"amountStart"`
	AmountEnd   float64 `json:"amountEnd"`
	Rate        float64 `json:"rate"`
}
