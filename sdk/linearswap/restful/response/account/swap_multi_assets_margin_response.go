package account

type SwapMultiAssetsMarginResponse struct {
	Code int `json:"code"`

	Data []struct {
		AssetsMode int `json:"assets_mode"`
	} `json:"data,omitempty"`

	Message string `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
