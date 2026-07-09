package copytradingtrade

type FollowerSettingsResponse struct {
    InstType           string `json:"instType"`
    Enable             bool   `json:"enable"`
    ProfitSharingRatio string `json:"profitSharingRatio"`
    MaxFollowers       string `json:"maxFollowers"`
}