package triggerorderclienttest

import (
	"github.com/HuobiRDCenter/huobi_futures_Golang/config"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful"
	requesttiggerorder "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/request/triggerorder"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/response/triggerorder"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
)

func RunAllExamples() {
	IsolatedPlaceOrderAsync()
	CrossPlaceOrderAsync()
	IsolatedCancelOrderAsync()
	CrossCancelOrderAsync()
	IsolatedGetOpenOrderAsync()
	CrossGetOpenOrderAsync()
	IsolatedGetHisOrderAsync()
	CrossGetHisOrderAsync()
	IsolatedTpslOrderAsync()
	CrossTpslOrderAsync()
	IsolatedTpslCancelAsync()
	CrossTpslCancelAsync()
	IsolatedGetTpslOpenOrderAsync()
	CrossGetTpslOpenOrderAsync()
	IsolatedGetTpslHisOrderAsync()
	CrossGetTpslHisOrderAsync()
	IsolatedGetRelationTpslOrderAsync()
	CrossGetRelationTpslOrderAsync()
}

func IsolatedPlaceOrderAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.PlaceOrderResponse)
	go client.IsolatedPlaceOrderAsync(resp, requesttiggerorder.PlaceOrderRequest{})
	x := <-resp
	log.Info("%v", x)
}

func CrossPlaceOrderAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.PlaceOrderResponse)
	go client.CrossPlaceOrderAsync(resp, requesttiggerorder.PlaceOrderRequest{})
	x := <-resp
	log.Info("%v", x)
}

func IsolatedCancelOrderAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.CancelOrderResponse)
	go client.IsolatedCancelOrderAsync(resp, "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func CrossCancelOrderAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.CancelOrderResponse)
	go client.CrossCancelOrderAsync(resp, "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func IsolatedGetOpenOrderAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.GetOpenOrderResponse)
	go client.IsolatedGetOpenOrderAsync(resp, "", 0, 100, 0)
	x := <-resp
	log.Info("%v", x)
}

func CrossGetOpenOrderAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.GetOpenOrderResponse)
	go client.CrossGetOpenOrderAsync(resp, "", 0, 100, 0, "")
	x := <-resp
	log.Info("%v", x)
}

func IsolatedGetHisOrderAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.GetHisOrderResponse)
	go client.IsolatedGetHisOrderAsync(resp, "", 0, "", 0, 0, 100, "")
	x := <-resp
	log.Info("%v", x)
}

func CrossGetHisOrderAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.GetHisOrderResponse)
	go client.CrossGetHisOrderAsync(resp, "", 0, "", 0, 0, 100, "", "")
	x := <-resp
	log.Info("%v", x)
}

func IsolatedTpslOrderAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.TpslOrderResponse)
	go client.IsolatedTpslOrderAsync(resp, requesttiggerorder.TpslOrderRequest{})
	x := <-resp
	log.Info("%v", x)
}

func CrossTpslOrderAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.TpslOrderResponse)
	go client.CrossTpslOrderAsync(resp, requesttiggerorder.TpslOrderRequest{})
	x := <-resp
	log.Info("%v", x)
}

func IsolatedTpslCancelAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.CancelOrderResponse)
	go client.IsolatedTpslCancelAsync(resp, "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func CrossTpslCancelAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.CancelOrderResponse)
	go client.CrossTpslCancelAsync(resp, "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func IsolatedGetTpslOpenOrderAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.GetOpenOrderResponse)
	go client.IsolatedGetTpslOpenOrderAsync(resp, "", 0, 100, 0)
	x := <-resp
	log.Info("%v", x)
}

func CrossGetTpslOpenOrderAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.GetOpenOrderResponse)
	go client.CrossGetTpslOpenOrderAsync(resp, "", 0, 100, 0, "")
	x := <-resp
	log.Info("%v", x)
}

func IsolatedGetTpslHisOrderAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.GetHisOrderResponse)
	go client.IsolatedGetTpslHisOrderAsync(resp, "", "", 0, 0, 100, "")
	x := <-resp
	log.Info("%v", x)
}

func CrossGetTpslHisOrderAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.GetHisOrderResponse)
	go client.CrossGetTpslHisOrderAsync(resp, "", "", 0, 0, 100, "", "")
	x := <-resp
	log.Info("%v", x)
}

func IsolatedGetRelationTpslOrderAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.GetRelationTpslOrderResponse)
	go client.IsolatedGetRelationTpslOrderAsync(resp, "", 1234)
	x := <-resp
	log.Info("%v", x)
}

func CrossGetRelationTpslOrderAsync() {
	client := new(restful.TriggerOrderClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan triggerorder.GetRelationTpslOrderResponse)
	go client.CrossGetRelationTpslOrderAsync(resp, "", 1234, "")
	x := <-resp
	log.Info("%v", x)
}
