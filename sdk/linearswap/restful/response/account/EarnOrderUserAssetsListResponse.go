package account

type EarnOrderUserAssetsListResponse struct {
	Code    int                                `json:"code"`
	Message string                             `json:"message"`
	Data    *PageInfoResVo1                    `json:"data"`
	Total   int                                `json:"total"`
	Items   []UserAssetsInfoCouponExpandResDto `json:"items"`
}
type PageInfoResVo1 struct {
	Total int                                `json:"total"`
	Items []UserAssetsInfoCouponExpandResDto `json:"items"`
}

type UserAssetsInfoCouponExpandResDto struct {
	ProjectId         int64  `json:"projectId"`
	OrderId           int64  `json:"orderId"`
	ProjectType       int    `json:"projectType"`
	Currency          string `json:"currency"`
	YesterdayIncome   string `json:"yesterdayIncome"`
	TotalIncomeAmount string `json:"totalIncomeAmount"`
	TotalAmount       string `json:"totalAmount"`
	MiningYearRate    string `json:"miningYearRate"`
	ApyType           int    `json:"apyType"`
}
