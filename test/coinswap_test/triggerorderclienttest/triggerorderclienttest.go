package triggerorderclienttest

import (
	"github.com/HuobiRDCenter/huobi_futures_Golang/config"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinswap/restful"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinswap/restful/response/triggerorder"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
)

func RunAllExamples() {
	SwapTriggerOrderAsync()
	SwapTriggerCancelAsync()
	SwapTriggerCancelAllAsync()
	SwapTriggerOpenOrdersAsync()
	SwapTriggerHisOrdersAsync()
	SwapTpslOrderAsync()
	SwapTpslCancelAsync()
	SwapTpslCancelAllAsync()
	SwapTpslOpenOrdersAsync()
	SwapTpslHisOrdersAsync()
	SwapRelationTpslOrderAsync()
	SwapTrackOrderAsync()
	SwapTrackCancelAsync()
	SwapTrackCancelAllAsync()
	SwapTrackOpenOrdersAsync()
	SwapTrackHisOrdersAsync()
}

func SwapTriggerOrderAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.SwapTriggerOrderResponse)
	go client.SwapTriggerOrderAsync(resp, "", "", "", "", "", "", "", "", "")
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

func SwapTriggerCancelAllAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.SwapTriggerCancelAllResponse)
	go client.SwapTriggerCancelAllAsync(resp, "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapTriggerOpenOrdersAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.SwapTriggerOpenOrdersResponse)
	go client.SwapTriggerOpenOrdersAsync(resp, "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapTriggerHisOrdersAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.SwapTriggerHisOrdersResponse)
	go client.SwapTriggerHisOrdersAsync(resp, "", "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapTpslOrderAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.SwapTpslOrderResponse)
	go client.SwapTpslOrderAsync(resp, "", "", "", "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapTpslCancelAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.SwapTpslCancelResponse)
	go client.SwapTpslCancelAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapTpslCancelAllAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.SwapTpslCancelAllResponse)
	go client.SwapTpslCancelAllAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapTpslOpenOrdersAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.SwapTpslOpenOrdersResponse)
	go client.SwapTpslOpenOrdersAsync(resp, "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapTpslHisOrdersAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.SwapTpslHisOrdersResponse)
	go client.SwapTpslHisOrdersAsync(resp, "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapRelationTpslOrderAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.SwapRelationTpslOrderResponse)
	go client.SwapRelationTpslOrderAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapTrackOrderAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.SwapTrackOrderResponse)
	go client.SwapTrackOrderAsync(resp, "", "", "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapTrackCancelAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.SwapTrackCancelResponse)
	go client.SwapTrackCancelAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapTrackCancelAllAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.SwapTrackCancelAllResponse)
	go client.SwapTrackCancelAllAsync(resp, "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapTrackOpenOrdersAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.SwapTrackOpenOrdersResponse)
	go client.SwapTrackOpenOrdersAsync(resp, "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapTrackHisOrdersAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.SwapTrackHisOrdersResponse)
	go client.SwapTrackHisOrdersAsync(resp, "", "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}
