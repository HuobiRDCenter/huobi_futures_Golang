package copytradingtrade

type StatisticsResponse struct {
    InstType          string                 `json:"instType"`
    TotalFollowerNum  string                 `json:"totalFollowerNum"`
    CurFollowerNum    string                 `json:"curFollowerNum"`
    TotalPnl          string                 `json:"totalPnl"`
    FollowerTotalPnl  string                 `json:"followerTotalPnl"`
    WinRate           string                 `json:"winRate"`
    WinNum            string                 `json:"winNum"`
    LossNum           string                 `json:"lossNum"`
    ProfitDetails90d  []ProfitDetailResponse `json:"profitDetails90d"`
}

type ProfitDetailResponse struct {
    PnlRate string `json:"pnlRate"`
    Pnl     string `json:"pnl"`
    Ts      string `json:"ts"`
}