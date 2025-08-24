package accountclienttest

import (
	"github.com/HuobiRDCenter/huobi_futures_Golang/config"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinfutures/restful"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinfutures/restful/response/account"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
)

func RunAllExamples() {
	ContractBalanceValuationAsync()
	ContractAccountInfoAsync()
	ContractPositionInfoAsync()
	ContractSubAuthAsync()
	ContractSubAuthListAsync()
	ContractSubAccountListAsync()
	ContractSubAccountInfoListAsync()
	ContractSubAccountInfoAsync()
	ContractSubPositionInfoAsync()
	ContractFinancialRecordAsync()
	ContractFinancialRecordExactAsync()
	ContractUserSettlementRecordsAsync()
	ContractOrderLimitAsync()
	ContractFeeAsync()
	ContractTransferLimitAsync()
	ContractPositionLimitAsync()
	ContractAccountPositionInfoAsync()
	ContractMasterSubTransferAsync()
	ContractMasterSubTransferRecordAsync()
	GetContractApiTradingStatusAsync()
	ContractAvailableLevelRateAsync()
}

func ContractBalanceValuationAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.ContractBalanceValuationResponse)
	go client.ContractBalanceValuationAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func ContractAccountInfoAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.ContractAccountInfoResponse)
	go client.ContractAccountInfoAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func ContractPositionInfoAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.ContractPositionInfoResponse)
	go client.ContractPositionInfoAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func ContractSubAuthAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.ContractSubAuthResponse)
	go client.ContractSubAuthAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractSubAuthListAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.ContractSubAuthListResponse)
	go client.ContractSubAuthListAsync(resp, "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractSubAccountListAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.ContractSubAccountListResponse)
	go client.ContractSubAccountListAsync(resp, "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractSubAccountInfoListAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.ContractSubAccountInfoListResponse)
	go client.ContractSubAccountInfoListAsync(resp, "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractSubAccountInfoAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.ContractSubAccountInfoResponse)
	go client.ContractSubAccountInfoAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractSubPositionInfoAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.ContractSubPositionInfoResponse)
	go client.ContractSubPositionInfoAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractFinancialRecordAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.ContractFinancialRecordResponse)
	go client.ContractFinancialRecordAsync(resp, "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractFinancialRecordExactAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.ContractFinancialRecordExactResponse)
	go client.ContractFinancialRecordExactAsync(resp, "", "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractUserSettlementRecordsAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.ContractUserSettlementRecordsResponse)
	go client.ContractUserSettlementRecordsAsync(resp, "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractOrderLimitAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.ContractOrderLimitResponse)
	go client.ContractOrderLimitAsync(resp, "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractFeeAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.ContractFeeResponse)
	go client.ContractFeeAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func ContractTransferLimitAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.ContractTransferLimitResponse)
	go client.ContractTransferLimitAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func ContractPositionLimitAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.ContractPositionLimitResponse)
	go client.ContractPositionLimitAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func ContractAccountPositionInfoAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.ContractAccountPositionInfoResponse)
	go client.ContractAccountPositionInfoAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}

func ContractMasterSubTransferAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.ContractMasterSubTransferResponse)
	go client.ContractMasterSubTransferAsync(resp, "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func ContractMasterSubTransferRecordAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.ContractMasterSubTransferRecordResponse)
	go client.ContractMasterSubTransferRecordAsync(resp, "", "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}

func GetContractApiTradingStatusAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.GetContractApiTradingStatusResponse)
	go client.GetContractApiTradingStatusAsync(resp)
	x := <-resp
	log.Info("%v", x)
}

func ContractAvailableLevelRateAsync() {
	client := new(restful.AccountClient).Init(config.AccessKey, config.SecretKey, config.Host, config.Sign)
	resp := make(chan account.ContractAvailableLevelRateResponse)
	go client.ContractAvailableLevelRateAsync(resp, "")
	x := <-resp
	log.Info("%v", x)
}
