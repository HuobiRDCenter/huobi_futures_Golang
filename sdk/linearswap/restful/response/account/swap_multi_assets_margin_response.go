package account

type SwapMultiAssetsMarginResponse struct {
	Code string `json:"code"`

	Data []struct {
		AssetsMode int `json:"assets_mode"`
	} `json:"data,omitempty"`

	Message int `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
