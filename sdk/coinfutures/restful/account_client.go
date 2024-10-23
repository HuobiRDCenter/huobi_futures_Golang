package restful

import (
	"encoding/json"
	"fmt"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinfutures"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinfutures/restful/response/account"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/reqbuilder"
)

type AccountClient struct {
	PUrlBuilder *reqbuilder.PrivateUrlBuilder
}

func (ac *AccountClient) Init(accessKey string, secretKey string, host string, sign string) *AccountClient {
	if host == "" {
		host = coinfutures.COIN_FUTURES_DEFAULT_HOST
	}
	ac.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host, sign)
	return ac
}

// ContractBalanceValuationAsync (Query Asset Valuation)
func (ac *AccountClient) ContractBalanceValuationAsync(data chan account.ContractBalanceValuationResponse, valuationAsset string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_balance_valuation", nil)
	print(url)

	// content
	content := ""
	if valuationAsset != "" {
		content = fmt.Sprintf(",\"valuation_asset\": \"%s\"", valuationAsset)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.ContractBalanceValuationResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractBalanceValuationResponse error: %s", jsonErr)
	}
	data <- result
}

// ContractAccountInfoAsync (Query User’s Account Information)
func (ac *AccountClient) ContractAccountInfoAsync(data chan account.ContractAccountInfoResponse, symbol string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_account_info", nil)

	// content
	content := ""
	if symbol != "" {
		content = fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.ContractAccountInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractAccountInfoResponse error: %s", jsonErr)
	}
	data <- result
}

// ContractPositionInfoAsync (Query User’s Position Information)
func (ac *AccountClient) ContractPositionInfoAsync(data chan account.ContractPositionInfoResponse, symbol string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_position_info", nil)

	// content
	content := ""
	if symbol != "" {
		content = fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.ContractPositionInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractPositionInfoResponse error: %s", jsonErr)
	}
	data <- result
}

// ContractSubAuthAsync (Set a Batch of Sub-Account Trading Permissions)
func (ac *AccountClient) ContractSubAuthAsync(data chan account.ContractSubAuthResponse, subUid string, subAuth string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_sub_auth", nil)

	// content
	content := ""
	if subUid != "" {
		content = fmt.Sprintf(",\"sub_uid\": \"%s\"", subUid)
	}
	if subAuth != "" {
		content = fmt.Sprintf(",\"sub_auth\": \"%s\"", subAuth)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.ContractSubAuthResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractSubAuthResponse error: %s", jsonErr)
	}
	data <- result
}

// ContractSubAuthListAsync (Query sub-account transaction permissions)
func (ac *AccountClient) ContractSubAuthListAsync(data chan account.ContractSubAuthListResponse, subUid string, startTime string,
	endTime string, direct string, fromId string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_sub_auth_list", nil)

	// content
	content := ""
	if subUid != "" {
		content = fmt.Sprintf(",\"sub_uid\": \"%s\"", subUid)
	}
	if startTime != "" {
		content = fmt.Sprintf(",\"start_time\": \"%s\"", startTime)
	}
	if endTime != "" {
		content = fmt.Sprintf(",\"end_time\": \"%s\"", endTime)
	}
	if direct != "" {
		content = fmt.Sprintf(",\"direct\": \"%s\"", direct)
	}
	if fromId != "" {
		content = fmt.Sprintf(",\"from_id\": \"%s\"", fromId)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.ContractSubAuthListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractSubAuthListResponse error: %s", jsonErr)
	}
	data <- result
}

// ContractSubAccountListAsync (Query assets information of all sub-accounts under the master account)
func (ac *AccountClient) ContractSubAccountListAsync(data chan account.ContractSubAccountListResponse, symbol string,
	direct string, fromId string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_sub_account_list", nil)

	// content
	content := ""
	if symbol != "" {
		content = fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if direct != "" {
		content = fmt.Sprintf(",\"direct\": \"%s\"", direct)
	}
	if fromId != "" {
		content = fmt.Sprintf(",\"from_id\": \"%s\"", fromId)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.ContractSubAccountListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractSubAccountListResponse error: %s", jsonErr)
	}
	data <- result
}

// ContractSubAccountInfoListAsync (Query a Batch of Sub-Account's Assets Information)
func (ac *AccountClient) ContractSubAccountInfoListAsync(data chan account.ContractSubAccountInfoListResponse, symbol string,
	pageIndex string, pageSize string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_sub_account_info_list", nil)

	// content
	content := ""
	if symbol != "" {
		content = fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if pageIndex != "" {
		content = fmt.Sprintf(",\"page_index\": \"%s\"", pageIndex)
	}
	if pageSize != "" {
		content = fmt.Sprintf(",\"page_size\": \"%s\"", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.ContractSubAccountInfoListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractSubAccountInfoListResponse error: %s", jsonErr)
	}
	data <- result
}

// ContractSubAccountInfoAsync (Query a single sub-account's assets information)
func (ac *AccountClient) ContractSubAccountInfoAsync(data chan account.ContractSubAccountInfoResponse, symbol string,
	subUid string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_sub_account_info", nil)

	// content
	content := ""
	if symbol != "" {
		content = fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if subUid != "" {
		content = fmt.Sprintf(",\"sub_uid\": \"%s\"", subUid)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.ContractSubAccountInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractSubAccountInfoResponse error: %s", jsonErr)
	}
	data <- result
}

// ContractSubPositionInfoAsync (Query a single sub-account's position information)
func (ac *AccountClient) ContractSubPositionInfoAsync(data chan account.ContractSubPositionInfoResponse, symbol string,
	subUid string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_sub_position_info", nil)

	// content
	content := ""
	if symbol != "" {
		content = fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if subUid != "" {
		content = fmt.Sprintf(",\"sub_uid\": \"%s\"", subUid)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.ContractSubPositionInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractSubPositionInfoResponse error: %s", jsonErr)
	}
	data <- result
}

// ContractFinancialRecordAsync (Query account financial records(New))
func (ac *AccountClient) ContractFinancialRecordAsync(data chan account.ContractFinancialRecordResponse, startTime string,
	endTime string, direct string, fromId string, fcType string, symbol string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v3/contract_financial_record", nil)

	// content
	content := ""
	if startTime != "" {
		content = fmt.Sprintf(",\"start_time\": \"%s\"", startTime)
	}
	if endTime != "" {
		content = fmt.Sprintf(",\"end_time\": \"%s\"", endTime)
	}
	if direct != "" {
		content = fmt.Sprintf(",\"direct\": \"%s\"", direct)
	}
	if fromId != "" {
		content = fmt.Sprintf(",\"from_id\": \"%s\"", fromId)
	}
	if fcType != "" {
		content = fmt.Sprintf(",\"type\": \"%s\"", fcType)
	}
	if symbol != "" {
		content = fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.ContractFinancialRecordResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractFinancialRecordResponse error: %s", jsonErr)
	}
	data <- result
}

// ContractFinancialRecordExactAsync (Query financial records via multiple fields(New))
func (ac *AccountClient) ContractFinancialRecordExactAsync(data chan account.ContractFinancialRecordExactResponse, startTime string,
	endTime string, direct string, fromId string, fcType string, symbol string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v3/contract_financial_record_exact", nil)

	// content
	content := ""
	if startTime != "" {
		content = fmt.Sprintf(",\"start_time\": \"%s\"", startTime)
	}
	if endTime != "" {
		content = fmt.Sprintf(",\"end_time\": \"%s\"", endTime)
	}
	if direct != "" {
		content = fmt.Sprintf(",\"direct\": \"%s\"", direct)
	}
	if fromId != "" {
		content = fmt.Sprintf(",\"from_id\": \"%s\"", fromId)
	}
	if fcType != "" {
		content = fmt.Sprintf(",\"type\": \"%s\"", fcType)
	}
	if symbol != "" {
		content = fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.ContractFinancialRecordExactResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractFinancialRecordExactResponse error: %s", jsonErr)
	}
	data <- result
}

// ContractUserSettlementRecordsAsync (Query user’s settlement records)
func (ac *AccountClient) ContractUserSettlementRecordsAsync(data chan account.ContractUserSettlementRecordsResponse, symbol string,
	startTime string, endTime string, pageIndex string, pageSize string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_user_settlement_records", nil)

	// content
	content := ""
	if symbol != "" {
		content = fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if startTime != "" {
		content = fmt.Sprintf(",\"start_time\": \"%s\"", startTime)
	}
	if endTime != "" {
		content = fmt.Sprintf(",\"end_time\": \"%s\"", endTime)
	}
	if pageIndex != "" {
		content = fmt.Sprintf(",\"page_index\": \"%s\"", pageIndex)
	}
	if pageSize != "" {
		content = fmt.Sprintf(",\"page_size\": \"%s\"", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.ContractUserSettlementRecordsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractUserSettlementRecordsResponse error: %s", jsonErr)
	}
	data <- result
}

// ContractOrderLimitAsync (Query contract information on order limit)
func (ac *AccountClient) ContractOrderLimitAsync(data chan account.ContractOrderLimitResponse, symbol string,
	orderPriceType string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_order_limit", nil)

	// content
	content := ""
	if symbol != "" {
		content = fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if orderPriceType != "" {
		content = fmt.Sprintf(",\"order_price_type\": \"%s\"", orderPriceType)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.ContractOrderLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractOrderLimitResponse error: %s", jsonErr)
	}
	data <- result
}

// ContractFeeAsync (Query information on contract trading fee)
func (ac *AccountClient) ContractFeeAsync(data chan account.ContractFeeResponse, symbol string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_fee", nil)

	// content
	content := ""
	if symbol != "" {
		content = fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.ContractFeeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractFeeResponse error: %s", jsonErr)
	}
	data <- result
}

// ContractTransferLimitAsync (Query information on Transfer Limit)
func (ac *AccountClient) ContractTransferLimitAsync(data chan account.ContractTransferLimitResponse, symbol string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_transfer_limit", nil)

	// content
	content := ""
	if symbol != "" {
		content = fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.ContractTransferLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractTransferLimitResponse error: %s", jsonErr)
	}
	data <- result
}

// ContractPositionLimitAsync (Query information on position limit)
func (ac *AccountClient) ContractPositionLimitAsync(data chan account.ContractPositionLimitResponse, symbol string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_position_limit", nil)

	// content
	content := ""
	if symbol != "" {
		content = fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.ContractPositionLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractPositionLimitResponse error: %s", jsonErr)
	}
	data <- result
}

// ContractAccountPositionInfoAsync (Query Assets And Positions)
func (ac *AccountClient) ContractAccountPositionInfoAsync(data chan account.ContractAccountPositionInfoResponse, symbol string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_account_position_info", nil)

	// content
	content := ""
	if symbol != "" {
		content = fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.ContractAccountPositionInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractAccountPositionInfoResponse error: %s", jsonErr)
	}
	data <- result
}

// ContractMasterSubTransferAsync (Transfer between master and sub account)
func (ac *AccountClient) ContractMasterSubTransferAsync(data chan account.ContractMasterSubTransferResponse, subUid string,
	symbol string, amount string, fcType string, clientOrderId string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_master_sub_transfer", nil)

	// content
	content := ""
	if subUid != "" {
		content = fmt.Sprintf(",\"sub_uid\": \"%s\"", subUid)
	}
	if symbol != "" {
		content = fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if amount != "" {
		content = fmt.Sprintf(",\"amount\": \"%s\"", amount)
	}
	if fcType != "" {
		content = fmt.Sprintf(",\"type\": \"%s\"", fcType)
	}
	if clientOrderId != "" {
		content = fmt.Sprintf(",\"client_order_id\": \"%s\"", clientOrderId)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.ContractMasterSubTransferResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractMasterSubTransferResponse error: %s", jsonErr)
	}
	data <- result
}

// ContractMasterSubTransferRecordAsync (Get transfer records between master and sub account)
func (ac *AccountClient) ContractMasterSubTransferRecordAsync(data chan account.ContractMasterSubTransferRecordResponse, symbol string,
	transferType string, createDate string, pageIndex string, pageSize string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_master_sub_transfer_record", nil)

	// content
	content := ""
	if symbol != "" {
		content = fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if transferType != "" {
		content = fmt.Sprintf(",\"transfer_type\": \"%s\"", transferType)
	}
	if createDate != "" {
		content = fmt.Sprintf(",\"create_date\": \"%s\"", createDate)
	}
	if pageIndex != "" {
		content = fmt.Sprintf(",\"page_index\": \"%s\"", pageIndex)
	}
	if pageSize != "" {
		content = fmt.Sprintf(",\"page_size\": \"%s\"", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.ContractMasterSubTransferRecordResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractMasterSubTransferRecordResponse error: %s", jsonErr)
	}
	data <- result
}

// GetContractApiTradingStatusAsync (Query user's API indicator disable information)
func (ac *AccountClient) GetContractApiTradingStatusAsync(data chan account.GetContractApiTradingStatusResponse) {
	// ulr
	url := ac.PUrlBuilder.Build(coinfutures.GET_METHOD, "/api/v1/contract_api_trading_status", nil)

	// content is nil
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetContractApiTradingStatusResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetContractApiTradingStatusResponse error: %s", jsonErr)
	}
	data <- result
}

// ContractAvailableLevelRateAsync (Query Available Leverage Rate)
func (ac *AccountClient) ContractAvailableLevelRateAsync(data chan account.ContractAvailableLevelRateResponse, symbol string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_available_level_rate", nil)

	// content
	content := ""
	if symbol != "" {
		content = fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.ContractAvailableLevelRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractAvailableLevelRateResponse error: %s", jsonErr)
	}
	data <- result
}
