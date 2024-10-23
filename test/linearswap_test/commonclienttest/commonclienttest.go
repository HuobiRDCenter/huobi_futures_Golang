package commonclienttest

import (
	"github.com/HuobiRDCenter/huobi_futures_Golang/config"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/response/common"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
)

func RunAllExamples() {
	GetSwapUnifiedAccountTypeAsync()
	SwapSwitchAccountTypeAsync()
}

func GetSwapUnifiedAccountTypeAsync() {
	client := new(restful.CommonClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan common.GetSwapUnifiedAccountTypeResponse)
	go client.GetSwapUnifiedAccountTypeAsync(resp)
	x := <-resp
	log.Info("%v", x)
}

func SwapSwitchAccountTypeAsync() {
	client := new(restful.CommonClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan common.SwapSwitchAccountTypeResponse)
	go client.SwapSwitchAccountTypeAsync(resp, 0)
	x := <-resp
	log.Info("%v", x)
}
