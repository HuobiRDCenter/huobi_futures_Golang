package transferclienttest

import (
	"github.com/HuobiRDCenter/huobi_futures_Golang/config"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinfutures/restful"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinfutures/restful/response/transfer"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
)

func RunAllExamples() {
	FuturesTransferAsync()
	AccountTransferAsync()
}

func FuturesTransferAsync() {
	client := new(restful.TransferClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan transfer.FuturesTransferResponse)
	go client.FuturesTransferAsync(resp, "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func AccountTransferAsync() {
	client := new(restful.TransferClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan transfer.AccountTransferResponse)
	go client.AccountTransferAsync(resp, "", "", 0, "", "")
	x := <-resp
	log.Info("%v", x)
}
