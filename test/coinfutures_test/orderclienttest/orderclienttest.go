package orderclienttest

import (
	"github.com/HuobiRDCenter/huobi_futures_Golang/config"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinfutures/restful"
	requestorder "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinfutures/restful/request/order"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinfutures/restful/response/order"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
)

func RunAllExamples() {
	ContractCancelAfterAsync()
	ContractOrderAsync()
	ContractBatchOrderAsync()
	ContractCancelAsync()
	ContractCancelAllAsync()
	ContractSwitchLeverRateAsync()
	ContractOrderInfoAsync()
	ContractOrderDetailAsync()
	ContractOpenOrdersAsync()
	ContractHisordersAsync()
	ContractHisOrdersExactAsync()
	ContractMatchResultsAsync()
	ContractMatchResultsExactAsync()
	LightningClosePositionAsync()
}

func ContractCancelAfterAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.ContractCancelAfterResponse)
	go client.ContractCancelAfterAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractOrderAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.ContractOrderResponse)
	go client.ContractOrderAsync(resp, requestorder.ContractOrderRequest{})
	x := <-resp
	log.Info("%v", x)
}

func ContractBatchOrderAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.ContractBatchOrderResponse)
	go client.ContractBatchOrderAsync(resp, requestorder.BatchContractOrderRequest{})
	x := <-resp
	log.Info("%v", x)
}

func ContractCancelAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.ContractCancelResponse)
	go client.ContractCancelAsync(resp, "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractCancelAllAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.ContractCancelAllResponse)
	go client.ContractCancelAllAsync(resp, "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractSwitchLeverRateAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.ContractSwitchLeverRateResponse)
	go client.ContractSwitchLeverRateAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractOrderInfoAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.ContractOrderInfoResponse)
	go client.ContractOrderInfoAsync(resp, "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractOrderDetailAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.ContractOrderDetailResponse)
	go client.ContractOrderDetailAsync(resp, "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractOpenOrdersAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.ContractOpenOrdersResponse)
	go client.ContractOpenOrdersAsync(resp, "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractHisordersAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.ContractHisordersResponse)
	go client.ContractHisordersAsync(resp, "", "", "", "", "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractHisOrdersExactAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.ContractHisOrdersExactResponse)
	go client.ContractHisOrdersExactAsync(resp, "", "", "", "", "", "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractMatchResultsAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.ContractMatchResultsResponse)
	go client.ContractMatchResultsAsync(resp, "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractMatchResultsExactAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.ContractMatchResultsExactResponse)
	go client.ContractMatchResultsExactAsync(resp, "", "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func LightningClosePositionAsync() {
	client := new(restful.OrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan order.LightningClosePositionResponse)
	go client.LightningClosePositionAsync(resp, "", "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}
