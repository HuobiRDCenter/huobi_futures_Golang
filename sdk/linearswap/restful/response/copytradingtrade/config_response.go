package copytradingtrade

type ConfigResponse struct {
    Enable             bool   `json:"enable"`
    ProfitSharingRatio string `json:"profitSharingRatio"`
    MaxFollowers       string `json:"maxFollowers"`
}