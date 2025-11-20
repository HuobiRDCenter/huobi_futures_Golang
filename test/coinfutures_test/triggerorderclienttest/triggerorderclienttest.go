package triggerorderclienttest

import (
	"github.com/HuobiRDCenter/huobi_futures_Golang/config"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinfutures/restful"
	requesttiggerorder "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinfutures/restful/request/triggerorder"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinfutures/restful/response/triggerorder"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
)

func RunAllExamples() {
	ContractTriggerOrderAsync()
	SwapTriggerCancelAsync()
	ContractTriggerCancelAllAsync()
	ContractTriggerOpenOrdersAsync()
	ContractTriggerHisOrdersAsync()
	ContractTpslOrderAsync()
	ContractTpslCancelAsync()
	ContractTpslCancelAllAsync()
	ContractTpslOpenOrdersAsync()
	ContractTpslHisOrdersAsync()
	ContractRelationTpslOrderAsync()
	ContractTrackOrderAsync()
	ContractTrackCancelAsync()
	ContractTrackCancelAllAsync()
	ContractTrackOpenOrdersAsync()
	ContractTrackHisOrdersAsync()
}

func ContractTriggerOrderAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.ContractTriggerOrderResponse)
	go client.ContractTriggerOrderAsync(resp, requesttiggerorder.ContractTriggerOrderRequest{})
	x := <-resp
	log.Info("%v", x)
}

func SwapTriggerCancelAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.SwapTriggerCancelResponse)
	go client.SwapTriggerCancelAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractTriggerCancelAllAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.ContractTriggerCancelAllResponse)
	go client.ContractTriggerCancelAllAsync(resp, "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractTriggerOpenOrdersAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.ContractTriggerOpenOrdersResponse)
	go client.ContractTriggerOpenOrdersAsync(resp, "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractTriggerHisOrdersAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.ContractTriggerHisOrdersResponse)
	go client.ContractTriggerHisOrdersAsync(resp, "", "", "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractTpslOrderAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.ContractTpslOrderResponse)
	go client.ContractTpslOrderAsync(resp, requesttiggerorder.ContractTriggerOrderRequest{})
	x := <-resp
	log.Info("%v", x)
}

func ContractTpslCancelAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.ContractTpslCancelResponse)
	go client.ContractTpslCancelAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractTpslCancelAllAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.ContractTpslCancelAllResponse)
	go client.ContractTpslCancelAllAsync(resp, "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractTpslOpenOrdersAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.ContractTpslOpenOrdersResponse)
	go client.ContractTpslOpenOrdersAsync(resp, "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractTpslHisOrdersAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.ContractTpslHisOrdersResponse)
	go client.ContractTpslHisOrdersAsync(resp, "", "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractRelationTpslOrderAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.ContractRelationTpslOrderResponse)
	go client.ContractRelationTpslOrderAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractTrackOrderAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.ContractTrackOrderResponse)
	go client.ContractTrackOrderAsync(resp, "", "", "", "", "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}
func ContractTrackCancelAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.ContractTrackCancelResponse)
	go client.ContractTrackCancelAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractTrackCancelAllAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.ContractTrackCancelAllResponse)
	go client.ContractTrackCancelAllAsync(resp, "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractTrackOpenOrdersAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.ContractTrackOpenOrdersResponse)
	go client.ContractTrackOpenOrdersAsync(resp, "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractTrackHisOrdersAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.ContractTrackHisOrdersResponse)
	go client.ContractTrackHisOrdersAsync(resp, "", "", "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}
