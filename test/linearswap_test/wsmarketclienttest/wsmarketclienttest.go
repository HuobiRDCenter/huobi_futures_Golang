package wsmarketclienttest

import (
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/ws"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/ws/response/market"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
	"time"
	"github.com/HuobiRDCenter/huobi_futures_Golang/config"
)

func RunAllExamples() {
	SubKLine()
	ReqKLine()
	SubDepth()
	SubIncrementalDepth()
	SubDetail()
	SubBBO()
	SubTradeDetail()
	ReqTradeDetail()
}
func SubKLine() {
	var wsmkClient = new(ws.WSMarketClient).Init("api.hbdm.com", config.Sign)
	wsmkClient.SubKLine("BTC-USDT", "15min", func(m *market.SubKLineResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(10) * time.Second)
}

func ReqKLine() {
	var wsmkClient = new(ws.WSMarketClient).Init("api.hbdm.com", config.Sign)
	wsmkClient.ReqKLine("BTC-USDT", "1min", func(m *market.ReqKLineResponse) {
		log.Info("%v", *m)
	}, 1604395758, 1604396758, "")
	time.Sleep(time.Duration(10) * time.Second)
}

func SubDepth() {
	var wsmkClient = new(ws.WSMarketClient).Init("api.hbdm.com", config.Sign)
	wsmkClient.SubDepth("BTC-USDT", "step0", func(m *market.SubDepthResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(5) * time.Second)
}

func SubIncrementalDepth() {
	var wsmkClient = new(ws.WSMarketClient).Init("api.hbdm.com", config.Sign)
	wsmkClient.SubIncrementalDepth("BTC-USDT", "20", func(m *market.SubDepthResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(5) * time.Second)
}

func SubDetail() {
	var wsmkClient = new(ws.WSMarketClient).Init("api.hbdm.com", config.Sign)
	wsmkClient.SubDetail("BTC-USDT", func(m *market.SubKLineResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(5) * time.Second)
}

func SubBBO() {
	var wsmkClient = new(ws.WSMarketClient).Init("api.hbdm.com", config.Sign)
	wsmkClient.SubBBO("BTC-USDT", func(m *market.SubBBOResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(5) * time.Second)
}

func SubTradeDetail() {
	var wsmkClient = new(ws.WSMarketClient).Init("api.hbdm.com", config.Sign)
	wsmkClient.SubTradeDetail("BTC-USDT", func(m *market.SubTradeDetailResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(5) * time.Second)
}

func ReqTradeDetail() {
	var wsmkClient = new(ws.WSMarketClient).Init("api.hbdm.com", config.Sign)
	wsmkClient.ReqTradeDetail("BTC-USDT", func(m *market.ReqTradeDetailResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(5) * time.Second)
}
