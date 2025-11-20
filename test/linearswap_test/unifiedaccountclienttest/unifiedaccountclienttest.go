package unifiedaccountclienttest

import (
	"github.com/HuobiRDCenter/huobi_futures_Golang/config"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/response/unifiedaccount"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
)

func RunAllExamples() {
	GetUnifiedAccountInfoAsync()
	GetLinearSwapOverviewAccountInfoAsync()
	LinearSwapFeeSwitchAsync()
	FixPositionMarginChangeAsync()
	FixPositionMarginChangeRecordAsync()
	SwapUnifiedAccountTypeAsync()
	SwapSwitchAccountTypeAsync()
}

func GetUnifiedAccountInfoAsync() {
	client := new(restful.UnifiedAccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan unifiedaccount.GetUnifiedAccountInfoResponse)
	go client.GetUnifiedAccountInfoAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func GetLinearSwapOverviewAccountInfoAsync() {
	client := new(restful.UnifiedAccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan unifiedaccount.GetLinearSwapOverviewAccountInfoResponse)
	go client.GetLinearSwapOverviewAccountInfoAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func LinearSwapFeeSwitchAsync() {
	client := new(restful.UnifiedAccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan unifiedaccount.LinearSwapFeeSwitchResponse)
	go client.LinearSwapFeeSwitchAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func FixPositionMarginChangeAsync() {
	client := new(restful.UnifiedAccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan unifiedaccount.FixPositionMarginChangeResponse)
	go client.FixPositionMarginChangeAsync(resp, "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func FixPositionMarginChangeRecordAsync() {
	client := new(restful.UnifiedAccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan unifiedaccount.FixPositionMarginChangeRecordResponse)
	go client.FixPositionMarginChangeRecordAsync(resp, "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapUnifiedAccountTypeAsync() {
	client := new(restful.UnifiedAccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan unifiedaccount.SwapUnifiedAccountTypeResponse)
	go client.SwapUnifiedAccountTypeAsync(resp)
	x := <-resp
	log.Info("%v", x)
}

func SwapSwitchAccountTypeAsync() {
	client := new(restful.UnifiedAccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan unifiedaccount.SwapSwitchAccountTypeResponse)
	go client.SwapSwitchAccountTypeAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}
