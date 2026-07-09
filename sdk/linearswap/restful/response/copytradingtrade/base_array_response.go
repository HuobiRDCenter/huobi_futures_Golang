package copytradingtrade

type BaseArrayResponse[T any] struct {
    Code string `json:"code"`
    Msg  string `json:"msg"`
    Data []T    `json:"data"`
}