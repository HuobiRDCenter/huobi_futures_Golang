package restful

import (
	"encoding/json"
	"fmt"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinswap"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinswap/restful/response/account"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/reqbuilder"
)

type AccountClient struct {
	PUrlBuilder *reqbuilder.PrivateUrlBuilder
}

func (ac *AccountClient) Init(accessKey string, secretKey string, host string, sign string) *AccountClient {
	if host == "" {
		host = coinswap.COIN_SWAP_DEFAULT_HOST
	}
	ac.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host, sign)
	return ac
}

// SwapBalanceValuationAsync (Query Asset Valuation)
func (ac *AccountClient) SwapBalanceValuationAsync(data chan account.SwapBalanceValuationResponse, valuationAsset string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_balance_valuation", nil)

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
	result := account.SwapBalanceValuationResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapBalanceValuationResponse error: %s", jsonErr)
	}
	data <- result
}

// SwapAccountInfoAsync (Query User’s Account Information)
func (ac *AccountClient) SwapAccountInfoAsync(data chan account.SwapAccountInfoResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_account_info", nil)

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.SwapAccountInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapAccountInfoResponse error: %s", jsonErr)
	}
	data <- result
}

// SwapPositionInfoAsync (Query User’s Position Information)
func (ac *AccountClient) SwapPositionInfoAsync(data chan account.SwapPositionInfoResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_position_info", nil)

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.SwapPositionInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapPositionInfoResponse error: %s", jsonErr)
	}
	data <- result
}

// SwapAccountPositionInfoAsync (Query Assets And Positions)
func (ac *AccountClient) SwapAccountPositionInfoAsync(data chan account.SwapAccountPositionInfoResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_account_position_info", nil)

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.SwapAccountPositionInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapAccountPositionInfoResponse error: %s", jsonErr)
	}
	data <- result
}

// SwapSubAuthAsync (Query sub-account transaction permissions)
func (ac *AccountClient) SwapSubAuthAsync(data chan account.SwapSubAuthResponse, subUid string, subAuth string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_sub_auth", nil)

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
	result := account.SwapSubAuthResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapSubAuthResponse error: %s", jsonErr)
	}
	data <- result
}

// SwapSubAuthListAsync (Set a Batch of Sub-Account Trading Permissions)
func (ac *AccountClient) SwapSubAuthListAsync(data chan account.SwapSubAuthListResponse, subUid string, startTime string,
	endTime string, direct string, fromId string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_sub_auth_list", nil)

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
	result := account.SwapSubAuthListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapSubAuthListResponse error: %s", jsonErr)
	}
	data <- result
}

// SwapSubAccountListAsync (Query assets information of all sub-accounts under the master account)
func (ac *AccountClient) SwapSubAccountListAsync(data chan account.SwapSubAccountListResponse, contractCode string,
	direct string, fromId string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_sub_account_list", nil)

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if direct != "" {
		content = fmt.Sprintf(",\"direct\": \"%s\"", direct)
	}
	if fromId != "" {
		content = fmt.Sprintf(",\"fromId\": \"%s\"", fromId)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.SwapSubAccountListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapSubAccountListResponse error: %s", jsonErr)
	}
	data <- result
}

// SwapSubAccountInfoListAsync (Query a Batch of Sub-Account's Assets Information)
func (ac *AccountClient) SwapSubAccountInfoListAsync(data chan account.SwapSubAccountInfoListResponse, contractCode string,
	pageIndex string, pageSize string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_sub_account_info_list", nil)

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
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
	result := account.SwapSubAccountInfoListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapSubAccountInfoListResponse error: %s", jsonErr)
	}
	data <- result
}

// SwapSubAccountInfoAsync (Query a single sub-account's assets information)
func (ac *AccountClient) SwapSubAccountInfoAsync(data chan account.SwapSubAccountInfoResponse, contractCode string,
	subUid string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_sub_account_info", nil)

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
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
	result := account.SwapSubAccountInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapSubAccountInfoResponse error: %s", jsonErr)
	}
	data <- result
}

// SwapSubPositionInfoAsync (Query a single sub-account's position information)
func (ac *AccountClient) SwapSubPositionInfoAsync(data chan account.SwapSubPositionInfoResponse, contractCode string,
	subUid string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_sub_position_info", nil)

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
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
	result := account.SwapSubPositionInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapSubPositionInfoResponse error: %s", jsonErr)
	}
	data <- result
}

// SwapFinancialRecordAsync (Query account financial records(New))
func (ac *AccountClient) SwapFinancialRecordAsync(data chan account.SwapFinancialRecordResponse, contract string,
	fcType string, startTime string, endTime string, direct string, fromId string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v3/swap_financial_record", nil)

	// content
	content := ""
	if contract != "" {
		content = fmt.Sprintf(",\"contract\": \"%s\"", contract)
	}
	if fcType != "" {
		content = fmt.Sprintf(",\"type\": \"%s\"", fcType)
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
	result := account.SwapFinancialRecordResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapFinancialRecordResponse error: %s", jsonErr)
	}
	data <- result
}

// SwapFinancialRecordExactAsync (Query financial records via multiple fields(New))
func (ac *AccountClient) SwapFinancialRecordExactAsync(data chan account.SwapFinancialRecordExactResponse, contract string,
	fcType string, startTime string, endTime string, direct string, fromId string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v3/swap_financial_record_exact", nil)

	// content
	content := ""
	if contract != "" {
		content = fmt.Sprintf(",\"contract\": \"%s\"", contract)
	}
	if fcType != "" {
		content = fmt.Sprintf(",\"type\": \"%s\"", fcType)
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
	result := account.SwapFinancialRecordExactResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapFinancialRecordExactResponse error: %s", jsonErr)
	}
	data <- result
}

// SwapUserSettlementRecordsAsync (Query user’s settlement records)
func (ac *AccountClient) SwapUserSettlementRecordsAsync(data chan account.SwapUserSettlementRecordsResponse, contractCode string,
	startTime string, endTime string, pageIndex string, pageSize string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_user_settlement_records", nil)

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
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
	result := account.SwapUserSettlementRecordsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapUserSettlementRecordsResponse error: %s", jsonErr)
	}
	data <- result
}

// SwapAvailableLevelRateAsync (Query user’s available leverage)
func (ac *AccountClient) SwapAvailableLevelRateAsync(data chan account.SwapAvailableLevelRateResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_available_level_rate", nil)

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.SwapAvailableLevelRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapAvailableLevelRateResponse error: %s", jsonErr)
	}
	data <- result
}

// SwapOrderLimitAsync (Query swap information on order limit)
func (ac *AccountClient) SwapOrderLimitAsync(data chan account.SwapOrderLimitResponse, contractCode string, orderPriceType string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_order_limit", nil)

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
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
	result := account.SwapOrderLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapOrderLimitResponse error: %s", jsonErr)
	}
	data <- result
}

// SwapFeeAsync (Query information on swap trading fee)
func (ac *AccountClient) SwapFeeAsync(data chan account.SwapFeeResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_fee", nil)

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.SwapFeeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapFeeResponse error: %s", jsonErr)
	}
	data <- result
}

// SwapTransferLimitAsync (Query information on Transfer Limit)
func (ac *AccountClient) SwapTransferLimitAsync(data chan account.SwapTransferLimitResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_transfer_limit", nil)

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.SwapTransferLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapTransferLimitResponse error: %s", jsonErr)
	}
	data <- result
}

// SwapPositionLimitAsync (Query information on position limit)
func (ac *AccountClient) SwapPositionLimitAsync(data chan account.SwapPositionLimitResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_position_limit", nil)

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.SwapPositionLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapPositionLimitResponse error: %s", jsonErr)
	}
	data <- result
}

// SwapMasterSubTransferAsync (Transfer between master and sub account)
func (ac *AccountClient) SwapMasterSubTransferAsync(data chan account.SwapMasterSubTransferResponse, subUid string,
	contractCode string, amount string, fcType string, clientOrderId string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_master_sub_transfer", nil)

	// content
	content := ""
	if subUid != "" {
		content = fmt.Sprintf(",\"sub_uid\": \"%s\"", subUid)
	}
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
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
	result := account.SwapMasterSubTransferResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapMasterSubTransferResponse error: %s", jsonErr)
	}
	data <- result
}

// SwapMasterSubTransferRecordAsync (Query transfer records between master and sub account)
func (ac *AccountClient) SwapMasterSubTransferRecordAsync(data chan account.SwapMasterSubTransferRecordResponse, contractCode string,
	transferType string, createDate string, pageIndex string, pageSize string) {
	// ulr
	url := ac.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_master_sub_transfer_record", nil)

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
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
	result := account.SwapMasterSubTransferRecordResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapMasterSubTransferRecordResponse error: %s", jsonErr)
	}
	data <- result
}

// SwapApiTradingStatusAsync (Query user's API indicator disable information)
func (ac *AccountClient) SwapApiTradingStatusAsync(data chan account.SwapApiTradingStatusResponse) {
	// ulr
	url := ac.PUrlBuilder.Build(coinswap.GET_METHOD, "/swap-api/v1/swap_api_trading_status", nil)

	// content
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.SwapApiTradingStatusResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapApiTradingStatusResponse error: %s", jsonErr)
	}
	data <- result
}
