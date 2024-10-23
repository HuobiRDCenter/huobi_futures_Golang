package wscenternotifyclienttest

import (
	"github.com/HuobiRDCenter/huobi_futures_Golang/config"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinswap/ws"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinswap/ws/response/centernotify"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
	"time"
)

func RunAllExamples() {
	Heartbeat()
}
func Heartbeat() {
	var wscnfClient = new(ws.WSCenterNotifyClient).Init(config.AccessKey, config.SecretKey, "", config.Sign)
	wscnfClient.SubHeartbeat("*", func(m *centernotify.SubHeartbeatResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(500) * time.Second)
	wscnfClient.UnsubHeartbeat("*", "")
	time.Sleep(time.Duration(500) * time.Second)
}
