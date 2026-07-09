package copytradingtrade

type ProfitSharingHistoryResponse struct {
    Id              string `json:"id"`
    InstType        string `json:"instType"`
    FollowerName    string `json:"followerName"`
    Ccy             string `json:"ccy"`
    ProfitSharingAmt string `json:"profitSharingAmt"`
    Ts              string `json:"ts"`
}