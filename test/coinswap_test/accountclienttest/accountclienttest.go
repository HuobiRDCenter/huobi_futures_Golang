package accountclienttest

import (
	"github.com/HuobiRDCenter/huobi_futures_Golang/config"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinswap/restful"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinswap/restful/response/account"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
)

func RunAllExamples() {
	SwapBalanceValuationAsync()
	SwapAccountInfoAsync()
	SwapPositionInfoAsync()
	SwapAccountPositionInfoAsync()
	SwapSubAuthAsync()
	SwapSubAuthListAsync()
	SwapSubAccountListAsync()
	SwapSubAccountInfoListAsync()
	SwapSubAccountInfoAsync()
	SwapSubPositionInfoAsync()
	SwapFinancialRecordAsync()
	SwapFinancialRecordExactAsync()
	SwapUserSettlementRecordsAsync()
	SwapAvailableLevelRateAsync()
	SwapOrderLimitAsync()
	SwapFeeAsync()
	SwapTransferLimitAsync()
	SwapPositionLimitAsync()
	SwapMasterSubTransferAsync()
	SwapMasterSubTransferRecordAsync()
	SwapApiTradingStatusAsync()
}

func SwapBalanceValuationAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SwapBalanceValuationResponse)
	go client.SwapBalanceValuationAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func SwapAccountInfoAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SwapAccountInfoResponse)
	go client.SwapAccountInfoAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func SwapPositionInfoAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SwapPositionInfoResponse)
	go client.SwapPositionInfoAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func SwapAccountPositionInfoAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SwapAccountPositionInfoResponse)
	go client.SwapAccountPositionInfoAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func SwapSubAuthAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SwapSubAuthResponse)
	go client.SwapSubAuthAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapSubAuthListAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SwapSubAuthListResponse)
	go client.SwapSubAuthListAsync(resp, "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapSubAccountListAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SwapSubAccountListResponse)
	go client.SwapSubAccountListAsync(resp, "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapSubAccountInfoListAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SwapSubAccountInfoListResponse)
	go client.SwapSubAccountInfoListAsync(resp, "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapSubAccountInfoAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SwapSubAccountInfoResponse)
	go client.SwapSubAccountInfoAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapSubPositionInfoAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SwapSubPositionInfoResponse)
	go client.SwapSubPositionInfoAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapFinancialRecordAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SwapFinancialRecordResponse)
	go client.SwapFinancialRecordAsync(resp, "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapFinancialRecordExactAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SwapFinancialRecordExactResponse)
	go client.SwapFinancialRecordExactAsync(resp, "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapUserSettlementRecordsAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SwapUserSettlementRecordsResponse)
	go client.SwapUserSettlementRecordsAsync(resp, "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapAvailableLevelRateAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SwapAvailableLevelRateResponse)
	go client.SwapAvailableLevelRateAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func SwapOrderLimitAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SwapOrderLimitResponse)
	go client.SwapOrderLimitAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapFeeAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SwapFeeResponse)
	go client.SwapFeeAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func SwapTransferLimitAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SwapTransferLimitResponse)
	go client.SwapTransferLimitAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func SwapPositionLimitAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SwapPositionLimitResponse)
	go client.SwapPositionLimitAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func SwapMasterSubTransferAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SwapMasterSubTransferResponse)
	go client.SwapMasterSubTransferAsync(resp, "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapMasterSubTransferRecordAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SwapMasterSubTransferRecordResponse)
	go client.SwapMasterSubTransferRecordAsync(resp, "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapApiTradingStatusAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SwapApiTradingStatusResponse)
	go client.SwapApiTradingStatusAsync(resp)
	x := <-resp
	log.Info("%v", x)
}
