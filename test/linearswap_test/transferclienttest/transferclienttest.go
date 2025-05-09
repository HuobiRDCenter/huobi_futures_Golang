package transferclienttest

import (
	"github.com/HuobiRDCenter/huobi_futures_Golang/config"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/response/transfer"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
)

func RunAllExamples() {
	TransferAsync()
}

func TransferAsync() {
	client := new(restful.TransferClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan transfer.TransferResponse)
	go client.TransferAsync(resp, "", "", 0, "", "")
	x := <-resp
	log.Info("%v", x)
}
