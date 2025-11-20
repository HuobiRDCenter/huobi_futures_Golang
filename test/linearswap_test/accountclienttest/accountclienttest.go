package accountclienttest

import (
	"github.com/HuobiRDCenter/huobi_futures_Golang/config"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/response/account"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
)

func RunAllExamples() {
	GetBalanceValuationAsync()
	IsolatedGetAccountInfoAsync()
	CrossGetAccountInfoAsync()
	IsolatedGetAccountPositionAsync()
	CrossGetAccountPositionAsync()
	IsolatedGetAssetsPositionAsync()
	CrossGetAssetsPositionAsync()
	SetSubAuthAsync()
	IsolatedGetSubAccountListResponseAsync()
	CrossGetSubAccountListAsync()
	IsolatedGetSubAccountInfoListAsync()
	CrossGetSubAccountInfoListAsync()
	AccountTransferAsync()
	GetAccountTransHisAsync()
	GetFinancialRecordExactAsync()
	IsolatedGetSettlementRecordsAsync()
	CrossGetSettlementRecordsAsync()
	IsolatedGetValidLeverRateAsync()
	CrossGetValidLeverRateAsync()
	GetOrderLimitAsync()
	GetFeeAsync()
	IsolatedGetTransferLimitAsync()
	CrossGetTransferLimitAsync()
	IsolatedGetPositionLimitAsync()
	CrossGetPositionLimitAsync()
	GetApiTradingStatusAsync()
	SwapSubPositionInfoAsync()
	SwapFinancialRecordAsync()
	SwapFinancialRecordExactAsync()
	SwapLeverPositionLimitAsync()
	SwapCrossLeverPositionLimitAsync()
	GetSwapSubAuthListAsync()
}

func GetBalanceValuationAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.GetBalanceValuationResponse)
	go client.GetBalanceValuationAsync(resp, "BTC")
	x := <-resp
	log.Info("%v", x)
}

func IsolatedGetAccountInfoAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.GetAccountInfoResponse)
	go client.IsolatedGetAccountInfoAsync(resp, "BTC-USDT", 0)
	x := <-resp
	log.Info("%v", x)
}

func CrossGetAccountInfoAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.GetCrossAccountInfoResponse)
	go client.CrossGetAccountInfoAsync(resp, "", 0)
	x := <-resp
	log.Info("%v", x)
}

func IsolatedGetAccountPositionAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.GetAccountPositionResponse)
	go client.IsolatedGetAccountPositionAsync(resp, "", 0)
	x := <-resp
	log.Info("%v", x)
}

func CrossGetAccountPositionAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.GetAccountPositionResponse)
	go client.CrossGetAccountPositionAsync(resp, "BTC-USDT", "swap", "BTC-USDT", 0)
	x := <-resp
	log.Info("%v", x)
}

func IsolatedGetAssetsPositionAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.GetAssetsPositionResponse)
	go client.IsolatedGetAssetsPositionAsync(resp, "BTC-USDT")
	x := <-resp
	log.Info("%v", x)
}

func CrossGetAssetsPositionAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.GetAssetsPositionResponseSingle)
	go client.CrossGetAssetsPositionAsync(resp, "USDT")
	x := <-resp
	log.Info("%v", x)
}

func SetSubAuthAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SetSubAuthResponse)
	go client.SetSubAuthAsync(resp, "123456", 1)
	x := <-resp
	log.Info("%v", x)
}

func IsolatedGetSubAccountListResponseAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.GetSubAccountListResponse)
	go client.IsolatedGetSubAccountListResponseAsync(resp, "BTC-USDT", "", 0)
	x := <-resp
	log.Info("%v", x)
}

func CrossGetSubAccountListAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.GetSubAccountListResponse)
	go client.CrossGetSubAccountListAsync(resp, "USDT", "", 0)
	x := <-resp
	log.Info("%v", x)
}

func IsolatedGetSubAccountInfoListAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.GetSubAccountInfoListResponse)
	go client.IsolatedGetSubAccountInfoListAsync(resp, "BTC-USDT", 1, 100)
	x := <-resp
	log.Info("%v", x)
}

func CrossGetSubAccountInfoListAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.GetSubAccountInfoListResponse)
	go client.CrossGetSubAccountInfoListAsync(resp, "BTC-USDT", 1, 100)
	x := <-resp
	log.Info("%v", x)
}

func AccountTransferAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.AccountTransferResponse)
	go client.AccountTransferAsync(resp, "USDT", "BTC-USDT", "USDT", 20, 123456, "master_to_sub", 456321)
	x := <-resp
	log.Info("%v", x)
}

func GetAccountTransHisAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.GetAccountTransHisResponse)
	go client.GetAccountTransHisAsync(resp, "BTC-USDT", true, "", 0, 0, 100)
	x := <-resp
	log.Info("%v", x)
}

func GetFinancialRecordExactAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.GetFinancialRecordExactResponse)
	go client.GetFinancialRecordExactAsync(resp, "", "", "", 0, 100, 0, 100, "")
	x := <-resp
	log.Info("%v", x)
}

func IsolatedGetSettlementRecordsAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.IsolatedGetSettlementRecordsResponse)
	go client.IsolatedGetSettlementRecordsAsync(resp, "", 0, 100, 0, 100)
	x := <-resp
	log.Info("%v", x)
}

func CrossGetSettlementRecordsAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.CrossGetSettlementRecordsResponse)
	go client.CrossGetSettlementRecordsAsync(resp, "", 0, 100, 0, 10)
	x := <-resp
	log.Info("%v", x)
}

func IsolatedGetValidLeverRateAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.GetValidLeverRateResponse)
	go client.IsolatedGetValidLeverRateAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func CrossGetValidLeverRateAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.GetValidLeverRateResponse)
	go client.CrossGetValidLeverRateAsync(resp, "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetOrderLimitAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.GetOrderLimitResponse)
	go client.GetOrderLimitAsync(resp, "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetFeeAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.GetFeeResponse)
	go client.GetFeeAsync(resp, "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func IsolatedGetTransferLimitAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.GetTransferLimitResponse)
	go client.IsolatedGetTransferLimitAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func CrossGetTransferLimitAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.GetTransferLimitResponse)
	go client.CrossGetTransferLimitAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func IsolatedGetPositionLimitAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.GetPositionLimitResponse)
	go client.IsolatedGetPositionLimitAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func CrossGetPositionLimitAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.GetPositionLimitResponse)
	go client.CrossGetPositionLimitAsync(resp, "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetApiTradingStatusAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.GetApiTradingStatusResponse)
	go client.GetApiTradingStatusAsync(resp)
	x := <-resp
	log.Info("%v", x)
}

func SwapSubPositionInfoAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SwapSubPositionInfoResponse)
	go client.SwapSubPositionInfoAsync(resp, "BTC-USDT", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapFinancialRecordAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SwapFinancialRecordResponse)
	go client.SwapFinancialRecordAsync(resp, "", "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapFinancialRecordExactAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SwapFinancialRecordResponse)
	go client.SwapFinancialRecordExactAsync(resp, "", "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapLeverPositionLimitAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SwapLeverPositionLimitResponse)
	go client.SwapLeverPositionLimitAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func SwapCrossLeverPositionLimitAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.SwapCrossLeverPositionLimitResponse)
	go client.SwapCrossLeverPositionLimitAsync(resp, "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetSwapSubAuthListAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.GetSwapSubAuthListResponse)
	go client.GetSwapSubAuthListAsync(resp, "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}
