package main

import (
	"github.com/HuobiRDCenter/huobi_futures_Golang/test/linearswap_test/accountclienttest"
	"github.com/HuobiRDCenter/huobi_futures_Golang/test/linearswap_test/commonclienttest"
	"github.com/HuobiRDCenter/huobi_futures_Golang/test/linearswap_test/marketclienttest"
	"github.com/HuobiRDCenter/huobi_futures_Golang/test/linearswap_test/orderclienttest"
	"github.com/HuobiRDCenter/huobi_futures_Golang/test/linearswap_test/transferclienttest"
	"github.com/HuobiRDCenter/huobi_futures_Golang/test/linearswap_test/triggerorderclienttest"
	"github.com/HuobiRDCenter/huobi_futures_Golang/test/linearswap_test/unifiedaccountclienttest"
	"github.com/HuobiRDCenter/huobi_futures_Golang/test/linearswap_test/wscenternotifyclienttest"
	"github.com/HuobiRDCenter/huobi_futures_Golang/test/linearswap_test/wsindexclienttest"
	"github.com/HuobiRDCenter/huobi_futures_Golang/test/linearswap_test/wsmarketclienttest"
	"github.com/HuobiRDCenter/huobi_futures_Golang/test/linearswap_test/wsnotifyclienttest"
)

func main() {
	runAll()
}

// Run all examples
func runAll() {
	accountclienttest.RunAllExamples()
	commonclienttest.RunAllExamples()
	marketclienttest.RunAllExamples()
	orderclienttest.RunAllExamples()
	transferclienttest.RunAllExamples()
	triggerorderclienttest.RunAllExamples()
	unifiedaccountclienttest.RunAllExamples()
	wsindexclienttest.RunAllExamples()
	wsmarketclienttest.RunAllExamples()
	wsnotifyclienttest.RunAllExamples()
	wscenternotifyclienttest.RunAllExamples()
}
