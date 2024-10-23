package wsindexclienttest

import (
	"github.com/HuobiRDCenter/huobi_futures_Golang/config"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinswap/ws"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinswap/ws/response/index"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
	"time"
)

func RunAllExamples() {
	SubPremiumIndexKLine()
	ReqPremiumIndexKLine()
	SubMarkPriceKLine()
	ReqMarkPriceKLine()
	SubEstimatedRateKLine()
	ReqEstimatedRateKLine()
	SubBasis()
	ReqBasis()
	SubIndexKLine()
	ReqIndexKLine()
}
func SubPremiumIndexKLine() {
	var wsixClient = new(ws.WSIndexClient).Init("", config.Sign)
	wsixClient.SubPremiumIndexKLine("BTC-USDT", "15min", func(m *index.SubIndexKLineResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(20) * time.Second)
}

func ReqPremiumIndexKLine() {
	var wsixClient = new(ws.WSIndexClient).Init("", config.Sign)
	wsixClient.ReqPremiumIndexKLine("BTC-USDT", "15min", func(m *index.ReqIndexKLineResponse) {
		log.Info("%v", *m)
	}, 1604395758, 1604396758, "")
	time.Sleep(time.Duration(20) * time.Second)
}

func SubMarkPriceKLine() {
	var wsixClient = new(ws.WSIndexClient).Init("", config.Sign)
	wsixClient.SubMarkPriceKLine("BTC-USDT", "15min", func(m *index.SubIndexKLineResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(20) * time.Second)
}

func ReqMarkPriceKLine() {
	var wsixClient = new(ws.WSIndexClient).Init("", config.Sign)
	wsixClient.ReqMarkPriceKLine("BTC-USDT", "1min", func(m *index.ReqIndexKLineResponse) {
		log.Info("%v", *m)
	}, 1612434060, 1612434260, "")
	time.Sleep(time.Duration(20) * time.Second)
}

func SubEstimatedRateKLine() {
	var wsixClient = new(ws.WSIndexClient).Init("", config.Sign)
	wsixClient.SubEstimatedRateKLine("BTC-USDT", "15min", func(m *index.SubIndexKLineResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(20) * time.Second)
}

func ReqEstimatedRateKLine() {
	var wsixClient = new(ws.WSIndexClient).Init("", config.Sign)
	wsixClient.ReqEstimatedRateKLine("BTC-USDT", "15min", func(m *index.ReqIndexKLineResponse) {
		log.Info("%v", *m)
	}, 1604395758, 1604396758, "")
	time.Sleep(time.Duration(20) * time.Second)
}

func SubBasis() {
	var wsixClient = new(ws.WSIndexClient).Init("", config.Sign)
	wsixClient.SubBasis("BTC-USDT", "15min", func(m *index.SubBasiesResponse) {
		log.Info("%v", *m)
	}, "", "")
	time.Sleep(time.Duration(20) * time.Second)
}

func ReqBasis() {
	var wsixClient = new(ws.WSIndexClient).Init("", config.Sign)
	wsixClient.ReqBasis("BTC-USDT", "15min", func(m *index.ReqBasisResponse) {
		log.Info("%v", *m)
	}, 1604395758, 1604396758, "", "")
	time.Sleep(time.Duration(20) * time.Second)
}

func SubIndexKLine() {
	var wsixClient = new(ws.WSIndexClient).Init("", config.Sign)
	wsixClient.SubIndexKLine("BTC-USDT", "15min", func(m *index.SubIndexKLineResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(20) * time.Second)
}

func ReqIndexKLine() {
	var wsixClient = new(ws.WSIndexClient).Init("", config.Sign)
	wsixClient.ReqIndexKLine("BTC-USDT", "15min", func(m *index.ReqIndexKLineResponse) {
		log.Info("%v", *m)
	}, 1604395758, 1604396758, "")
	time.Sleep(time.Duration(20) * time.Second)
}
