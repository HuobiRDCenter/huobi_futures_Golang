package orderclienttest

import (
	"github.com/HuobiRDCenter/huobi_futures_Golang/config"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinswap/restful"
	requestorder "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinswap/restful/request/order"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinswap/restful/response/order"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
)

func RunAllExamples() {
	SwapCancelAfterAsync()
	SwapOrderAsync()
	SwapBatchOrderAsync()
	SwapCancelAsync()
	SwapCancelAllAsync()
	SwapSwitchLeverRateAsync()
	SwapOrderInfoAsync()
	SwapOrderDetailAsync()
	SwapOpenOrdersAsync()
	SwapHisOrdersAsync()
	SwapHisOrdersExactAsync()
	SwapMatchResultsAsync()
	SwapMatchResultsExactAsync()
	SwapLightningClosePositionAsync()
}

func SwapCancelAfterAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.SwapCancelAfterResponse)
	go client.SwapCancelAfterAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapOrderAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.SwapOrderResponse)
	go client.SwapOrderAsync(resp, requestorder.SwapOrderRequest{})
	x := <-resp
	log.Info("%v", x)
}

func SwapBatchOrderAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.SwapBatchOrderResponse)
	go client.SwapBatchOrderAsync(resp, requestorder.BatchSwapOrderRequest{})
	x := <-resp
	log.Info("%v", x)
}

func SwapCancelAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.SwapCancelResponse)
	go client.SwapCancelAsync(resp, "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapCancelAllAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.SwapCancelAllResponse)
	go client.SwapCancelAllAsync(resp, "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapSwitchLeverRateAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.SwapSwitchLeverRateResponse)
	go client.SwapSwitchLeverRateAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapOrderInfoAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.SwapOrderInfoResponse)
	go client.SwapOrderInfoAsync(resp, "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapOrderDetailAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.SwapOrderDetailResponse)
	go client.SwapOrderDetailAsync(resp, "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapOpenOrdersAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.SwapOpenOrdersResponse)
	go client.SwapOpenOrdersAsync(resp, "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapHisOrdersAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.SwapHisOrdersResponse)
	go client.SwapHisOrdersAsync(resp, "", "", "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapHisOrdersExactAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.SwapHisOrdersExactResponse)
	go client.SwapHisOrdersExactAsync(resp, "", "", "", "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapMatchResultsAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.SwapMatchResultsResponse)
	go client.SwapMatchResultsAsync(resp, "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapMatchResultsExactAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.SwapMatchResultsExactResponse)
	go client.SwapMatchResultsExactAsync(resp, "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapLightningClosePositionAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.SwapLightningClosePositionResponse)
	go client.SwapLightningClosePositionAsync(resp, "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}
