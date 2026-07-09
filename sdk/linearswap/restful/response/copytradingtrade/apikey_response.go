package copytradingtrade

type ApikeyResponse struct {
    Label     string   `json:"label"`
    AccessKey string   `json:"accessKey"`
    SecretKey string   `json:"secretKey"`
    Perm      []string `json:"perm"`
    Ts        string   `json:"ts"`
}