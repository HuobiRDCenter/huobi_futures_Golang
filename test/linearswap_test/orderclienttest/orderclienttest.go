package orderclienttest

import (
	"github.com/HuobiRDCenter/huobi_futures_Golang/config"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful"
	requestorder "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/request/order"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/response/order"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
)

func RunAllExamples() {
	IsolatedPlaceOrderAsync()
	CrossPlaceOrderAsync()
	IsolatedPlaceBatchOrderAsync()
	CrossPlaceBatchOrderAsync()
	IsolatedCancelOrderAsync()
	CrossCancelOrderAsync()
	IsolatedSwitchLeverRateAsync()
	CrossSwitchLeverRateAsync()
	IsolatedGetOrderInfoAsync()
	CrossGetOrderInfoAsync()
	IsolatedGetOrderDetailAsync()
	CrossGetOrderDetailAsync()
	IsolatedGetOpenOrderAsync()
	CrossGetOpenOrderAsync()
	IsolatedGetHisOrderAsync()
	CrossGetHisOrderAsync()
	IsolatedGetHisOrderExactAsync()
	CrossGetHisOrderExactAsync()
	IsolatedGetHisMatchAsync()
	CrossGetHisMatchAsync()
	IsolatedGetHisMatchExactAsync()
	CrossGetHisMatchExactAsync()
	IsolatedLightningCloseAsync()
	CrossLightningCloseAsync()
	LinearCancelAfterAsync()
	SwapSwitchPositionModeAsync()
	SwapCrossSwitchPositionModeAsync()
	SwapHisordersAsync()
}

func IsolatedPlaceOrderAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.PlaceOrderResponse)
	go client.IsolatedPlaceOrderAsync(resp, requestorder.PlaceOrderRequest{})
	x := <-resp
	log.Info("%v", x)
}

func CrossPlaceOrderAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.PlaceOrderResponse)
	go client.CrossPlaceOrderAsync(resp, requestorder.PlaceOrderRequest{})
	x := <-resp
	log.Info("%v", x)
}

func IsolatedPlaceBatchOrderAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.PlaceBatchOrderResponse)
	go client.IsolatedPlaceBatchOrderAsync(resp, requestorder.BatchPlaceOrderRequest{})
	x := <-resp
	log.Info("%v", x)
}

func CrossPlaceBatchOrderAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.PlaceBatchOrderResponse)
	go client.CrossPlaceBatchOrderAsync(resp, requestorder.BatchPlaceOrderRequest{})
	x := <-resp
	log.Info("%v", x)
}

func IsolatedCancelOrderAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.CancelOrderResponse)
	go client.IsolatedCancelOrderAsync(resp, "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func CrossCancelOrderAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.CancelOrderResponse)
	go client.CrossCancelOrderAsync(resp, "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func IsolatedSwitchLeverRateAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.SwitchLeverRateResponse)
	go client.IsolatedSwitchLeverRateAsync(resp, "", 0)
	x := <-resp
	log.Info("%v", x)
}

func CrossSwitchLeverRateAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.SwitchLeverRateResponse)
	go client.CrossSwitchLeverRateAsync(resp, "", 0, "", "")
	x := <-resp
	log.Info("%v", x)
}

func IsolatedGetOrderInfoAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.GetOrderInfoResponse)
	go client.IsolatedGetOrderInfoAsync(resp, "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func CrossGetOrderInfoAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.GetOrderInfoResponse)
	go client.CrossGetOrderInfoAsync(resp, "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func IsolatedGetOrderDetailAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.GetOrderDetailResponse)
	go client.IsolatedGetOrderDetailAsync(resp, "", 1234, 0, 0, 0, 100)
	x := <-resp
	log.Info("%v", x)
}

func CrossGetOrderDetailAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.GetOrderDetailResponse)
	go client.CrossGetOrderDetailAsync(resp, "", 1234, 0, 0, 0, 100, "")
	x := <-resp
	log.Info("%v", x)
}

func IsolatedGetOpenOrderAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.GetOpenOrderResponse)
	go client.IsolatedGetOpenOrderAsync(resp, "", 0, 100, "", 0)
	x := <-resp
	log.Info("%v", x)
}

func CrossGetOpenOrderAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.GetOpenOrderResponse)
	go client.CrossGetOpenOrderAsync(resp, "", 0, 100, "", 0)
	x := <-resp
	log.Info("%v", x)
}

func IsolatedGetHisOrderAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.GetHisOrderResponse)
	go client.IsolatedGetHisOrderAsync(resp, "", 0, 0, "", 0, 0, 100, "")
	x := <-resp
	log.Info("%v", x)
}

func CrossGetHisOrderAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.GetHisOrderResponse)
	go client.CrossGetHisOrderAsync(resp, "", 0, 0, "", 0, 0, 100, "")
	x := <-resp
	log.Info("%v", x)
}

func IsolatedGetHisOrderExactAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.GetHisOrderExactResponse)
	go client.IsolatedGetHisOrderExactAsync(resp, "", 0, 0, "", "", 0, 100, 0, 100, "")
	x := <-resp
	log.Info("%v", x)
}

func CrossGetHisOrderExactAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.GetHisOrderExactResponse)
	go client.CrossGetHisOrderExactAsync(resp, "", 0, 0, "", "", 0, 100, 0, 100, "")
	x := <-resp
	log.Info("%v", x)
}
func IsolatedGetHisMatchAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.GetHisMatchResponse)
	go client.IsolatedGetHisMatchAsync(resp, "", 0, 0, 0, 100)
	x := <-resp
	log.Info("%v", x)
}

func CrossGetHisMatchAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.GetHisMatchResponse)
	go client.CrossGetHisMatchAsync(resp, "", 0, 0, 0, 100)
	x := <-resp
	log.Info("%v", x)
}

func IsolatedGetHisMatchExactAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.GetHisMatchExactResponse)
	go client.IsolatedGetHisMatchExactAsync(resp, "", 0, 0, 100, 0, 0, "")
	x := <-resp
	log.Info("%v", x)
}

func CrossGetHisMatchExactAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.GetHisMatchExactResponse)
	go client.CrossGetHisMatchExactAsync(resp, "", 0, 0, 100, 0, 100, "")
	x := <-resp
	log.Info("%v", x)
}

func IsolatedLightningCloseAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.LightningCloseResponse)
	go client.IsolatedLightningCloseAsync(resp, "", 0, "", 1234, "")
	x := <-resp
	log.Info("%v", x)
}

func CrossLightningCloseAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.LightningCloseResponse)
	go client.CrossLightningCloseAsync(resp, "", 0, "", 1234, "", "")
	x := <-resp
	log.Info("%v", x)
}

func LinearCancelAfterAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.LinearCancelAfterResponse)
	go client.LinearCancelAfterAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapSwitchPositionModeAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.SwapSwitchPositionModeResponse)
	go client.SwapSwitchPositionModeAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapCrossSwitchPositionModeAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.SwapSwitchPositionModeResponse)
	go client.SwapCrossSwitchPositionModeAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapHisordersAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.SwapHisordersResponse)
	go client.SwapHisordersAsync(resp, "", "", "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}
