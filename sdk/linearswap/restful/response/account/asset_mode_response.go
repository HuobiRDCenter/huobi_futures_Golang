package account

type AssetModeResponse struct {
	Code int `json:"code"`

	Data struct {
		AssetMode int `json:"asset_mode"`
	} `json:"data,omitempty"`

	Message string `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
