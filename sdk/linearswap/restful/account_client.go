package restful

import (
	"encoding/json"
	"fmt"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/response/account"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/reqbuilder"
)

type AccountClient struct {
	PUrlBuilder *reqbuilder.PrivateUrlBuilder
}

func (ac *AccountClient) Init(accessKey string, secretKey string, host string) *AccountClient {
	if host == "" {
		host = linearswap.LINEAR_SWAP_DEFAULT_HOST
	}
	ac.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return ac
}

// GetBalanceValuationAsync [General]Query Asset Valuation
func (ac *AccountClient) GetBalanceValuationAsync(data chan account.GetBalanceValuationResponse, valuationAsset string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_balance_valuation", nil)

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
	result := account.GetBalanceValuationResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetBalanceValuationResponse error: %s", jsonErr)
	}
	data <- result
}

// IsolatedGetAccountInfoAsync [Isolated] Query User’s Account Information and [Isolated] Query a single sub-account's assets information
func (ac *AccountClient) IsolatedGetAccountInfoAsync(data chan account.GetAccountInfoResponse, contractCode string, subUid int64) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_account_info", nil)
	if subUid != 0 {
		url = ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_sub_account_info", nil)
	}

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if subUid != 0 {
		content += fmt.Sprintf(",\"sub_uid\": %d", subUid)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetAccountInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAccountInfoResponse error: %s", jsonErr)
	}
	data <- result
}

// CrossGetAccountInfoAsync [Cross] Query User's Account Information and [Cross] Query A Sub-Account's Assets Information
func (ac *AccountClient) CrossGetAccountInfoAsync(data chan account.GetCrossAccountInfoResponse, marginAccount string, subUid int64) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_account_info", nil)
	if subUid != 0 {
		url = ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_sub_account_info", nil)
	}

	// content
	content := ""
	if marginAccount != "" {
		content = fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	}
	if subUid != 0 {
		content += fmt.Sprintf(",\"sub_uid\": %d", subUid)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetCrossAccountInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAccountInfoResponse error: %s", jsonErr)
	}
	data <- result
}

// IsolatedGetAccountPositionAsync [Isolated] Query User’s Position Information and [Isolated] Query a single sub-account's assets information
func (ac *AccountClient) IsolatedGetAccountPositionAsync(data chan account.GetAccountPositionResponse, contractCode string, subUid int64) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_position_info", nil)
	if subUid != 0 {
		url = ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_sub_account_info", nil)
	}

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if subUid != 0 {
		content += fmt.Sprintf(",\"sub_uid\": %d", subUid)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetAccountPositionResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAccountPositionResponse error: %s", jsonErr)
	}
	data <- result
}

// CrossGetAccountPositionAsync [Cross] Query User's Position Information and [Cross] Query A Sub-Account's Position Information
func (ac *AccountClient) CrossGetAccountPositionAsync(data chan account.GetAccountPositionResponse, contractCode string, contractType string, pair string, subUid int64) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_position_info", nil)
	if subUid != 0 {
		url = ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_sub_position_info", nil)
	}

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if contractType != "" {
		content = fmt.Sprintf(",\"contract_type\": \"%s\"", contractType)
	}
	if subUid != 0 {
		content += fmt.Sprintf(",\"sub_uid\": %d", subUid)
	}
	if pair != "" {
		content += fmt.Sprintf(",\"pair\": %s", pair)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetAccountPositionResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAccountPositionResponse error: %s", jsonErr)
	}
	data <- result
}

// IsolatedGetAssetsPositionAsync [Isolated] Query Assets And Positions
func (ac *AccountClient) IsolatedGetAssetsPositionAsync(data chan account.GetAssetsPositionResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_account_position_info", nil)

	// content
	content := fmt.Sprintf("{\"contract_code\": \"%s\"}", contractCode)

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetAssetsPositionResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAssetsPositionResponse error: %s", jsonErr)
	}
	data <- result
}

// CrossGetAssetsPositionAsync [Cross] Query Assets And Positions
func (ac *AccountClient) CrossGetAssetsPositionAsync(data chan account.GetAssetsPositionResponseSingle, marginAccount string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_account_position_info", nil)

	// content
	content := fmt.Sprintf("{\"margin_account\": \"%s\"}", marginAccount)

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetAssetsPositionResponseSingle{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAssetsPositionResponse error: %s", jsonErr)
	}
	data <- result
}

// SetSubAuthAsync [General]Set a Batch of Sub-Account Trading Permissions
func (ac *AccountClient) SetSubAuthAsync(data chan account.SetSubAuthResponse, subUid string, subAuth int) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_sub_auth", nil)

	// content
	content := fmt.Sprintf("{\"sub_uid\": \"%s\", \"sub_auth\": %d}", subUid, subAuth)

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.SetSubAuthResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SetSubAuthResponse error: %s", jsonErr)
	}
	data <- result
}

// IsolatedGetSubAccountListResponseAsync [Isolated] Query assets information of all sub-accounts under the master account
func (ac *AccountClient) IsolatedGetSubAccountListResponseAsync(data chan account.GetSubAccountListResponse, contractCode string, direct string, fromId int64) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_sub_account_list", nil)

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if direct != "" {
		content = fmt.Sprintf(",\"direct\": \"%s\"", direct)
	}
	if fromId != 0 {
		content = fmt.Sprintf(",\"from_id\": \"%d\"", fromId)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetSubAccountListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAllSubAssetsResponse error: %s", getErr)
	}
	data <- result
}

// CrossGetSubAccountListAsync [Cross] Query Assets Information Of All Sub-Accounts Under The Master Account
func (ac *AccountClient) CrossGetSubAccountListAsync(data chan account.GetSubAccountListResponse, marginAccount string, direct string, fromId int64) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_sub_account_list", nil)

	// content
	content := ""
	if marginAccount != "" {
		content = fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	}
	if direct != "" {
		content = fmt.Sprintf(",\"direct\": \"%s\"", direct)
	}
	if fromId != 0 {
		content = fmt.Sprintf(",\"from_id\": \"%d\"", fromId)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetSubAccountListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSubAccountListResponse error: %s", getErr)
	}
	data <- result
}

// IsolatedGetSubAccountInfoListAsync [Isolated]Query a Batch of Sub-Account's Assets Information
func (ac *AccountClient) IsolatedGetSubAccountInfoListAsync(data chan account.GetSubAccountInfoListResponse,
	contractCode string, pageIndex int, pageSize int) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_sub_account_info_list", nil)

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if pageIndex != 0 {
		content = fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content = fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetSubAccountInfoListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSubAccountInfoListResponse error: %s", getErr)
	}
	data <- result
}

// CrossGetSubAccountInfoListAsync [Cross]Query a Batch of Sub-Account's Assets Information
func (ac *AccountClient) CrossGetSubAccountInfoListAsync(data chan account.GetSubAccountInfoListResponse,
	marginAccount string, pageIndex int, pageSize int) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_sub_account_info_list", nil)

	// content
	content := ""
	if marginAccount != "" {
		content = fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	}
	if pageIndex != 0 {
		content = fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content = fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetSubAccountInfoListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSubAccountInfoListResponse error: %s", getErr)
	}
	data <- result
}

// AccountTransferAsync [General] Transfer between master and sub account and [General] Transfer between different margin accounts under the same account
func (ac *AccountClient) AccountTransferAsync(data chan account.AccountTransferResponse, asset string, fromMarginAccount string, toMarginAccount string, amount float64,
	subUid int64, fcType string, clientOrderId int64) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_master_sub_transfer", nil)
	if subUid == 0 {
		url = ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_transfer_inner", nil)
	}

	// content
	content := fmt.Sprintf(",\"asset\":\"%s\", \"from_margin_account\":\"%s\", \"to_margin_account\":\"%s\", \"amount\":%f",
		asset, fromMarginAccount, toMarginAccount, amount)
	if subUid != 0 {
		content += fmt.Sprintf(",\"sub_uid\": %d,\"type\": \"%s\"", subUid, fcType)
	}
	if clientOrderId != 0 {
		content += fmt.Sprintf(",\"client_order_id\": %d", clientOrderId)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.AccountTransferResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to AccountTransferResponse error: %s", jsonErr)
	}
	data <- result
}

// GetAccountTransHisAsync [General] Query transfer records between master and sub account
func (ac *AccountClient) GetAccountTransHisAsync(data chan account.GetAccountTransHisResponse, marginAccount string,
	beMasterSub bool, fcType string, createDate int, pageIndex int, pageSize int) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_financial_record", nil)
	if beMasterSub {
		url = ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_master_sub_transfer_record", nil)
	}

	// content
	content := fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	if fcType != "" {
		content += fmt.Sprintf(",\"type\": \"%s\"", fcType)
		if beMasterSub {
			content += fmt.Sprintf(",\"transfer_type\": \"%s\"", fcType)
		}
	}
	if createDate != 0 {
		content += fmt.Sprintf(",\"create_date\": %d", createDate)
	} else {
		if beMasterSub {
			createDate = 7
			content += fmt.Sprintf(",\"create_date\": %d", createDate)
		}
	}
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetAccountTransHisResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAccountTransHisResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) GetFinancialRecordExactAsync(data chan account.GetFinancialRecordExactResponse, marginAccount string,
	contractCode string, fcType string, startTime int64, endTime int64, fromId int64, size int, direct string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_financial_record_exact", nil)

	// content
	content := fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if fcType != "" {
		content += fmt.Sprintf(",\"type\": \"%s\"", fcType)
	}
	if startTime != 0 {
		content += fmt.Sprintf(",\"start_time\": %d", startTime)
	}
	if endTime != 0 {
		content += fmt.Sprintf(",\"end_time\": %d", endTime)
	}
	if fromId != 0 {
		content += fmt.Sprintf(",\"from_id\": %d", fromId)
	}
	if size != 0 {
		content += fmt.Sprintf(",\"size\": %d", size)
	}
	if direct != "" {
		content += fmt.Sprintf(",\"direct\": \"%s\"", direct)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetFinancialRecordExactResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetFinancialRecordExactResponse error: %s", jsonErr)
	}
	data <- result
}

// IsolatedGetSettlementRecordsAsync [Isolated]Query Settlement Records of Users
func (ac *AccountClient) IsolatedGetSettlementRecordsAsync(data chan account.IsolatedGetSettlementRecordsResponse, contractCode string,
	startTime int64, endTime int64, pageIndex int, pageSize int) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_user_settlement_records", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if startTime != 0 {
		content += fmt.Sprintf(",\"start_time\": %d", startTime)
	}
	if endTime != 0 {
		content += fmt.Sprintf(",\"end_time\": %d", endTime)
	}
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.IsolatedGetSettlementRecordsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to IsolatedGetSettlementRecordsResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) CrossGetSettlementRecordsAsync(data chan account.CrossGetSettlementRecordsResponse, marginAccount string,
	startTime int64, endTime int64, pageIndex int, pageSize int) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_user_settlement_records", nil)

	// content
	content := fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	if startTime != 0 {
		content += fmt.Sprintf(",\"start_time\": %d", startTime)
	}
	if endTime != 0 {
		content += fmt.Sprintf(",\"end_time\": %d", endTime)
	}
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.CrossGetSettlementRecordsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to CrossGetSettlementRecordsResponse error: %s", jsonErr)
	}
	data <- result
}

// IsolatedGetValidLeverRateAsync [Isolated] Query user’s available leverage
func (ac *AccountClient) IsolatedGetValidLeverRateAsync(data chan account.GetValidLeverRateResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_available_level_rate", nil)

	// content
	content := fmt.Sprintf("{ \"contract_code\": \"%s\" }", contractCode)
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetValidLeverRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetValidLeverRateResponse error: %s", jsonErr)
	}
	data <- result
}

// CrossGetValidLeverRateAsync [Cross] Query User’s Available Leverage
func (ac *AccountClient) CrossGetValidLeverRateAsync(data chan account.GetValidLeverRateResponse, contractCode string, pair string, contractType string, businessType string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_available_level_rate", nil)

	// content
	content := fmt.Sprintf("{ \"contract_code\": \"%s\" }", contractCode)
	if pair != "" {
		content += fmt.Sprintf(",\"pair\": %s", pair)
	}
	if contractType != "" {
		content += fmt.Sprintf(",\"contract_type\": %s", contractType)
	}
	if businessType != "" {
		content += fmt.Sprintf(",\"business_type\": %s", businessType)
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetValidLeverRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetValidLeverRateResponse error: %s", jsonErr)
	}
	data <- result
}

// GetOrderLimitAsync [General] Query swap information on order limit
func (ac *AccountClient) GetOrderLimitAsync(data chan account.GetOrderLimitResponse, orderPriceType string, contractCode string, pair string, contractType string, businessType string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_order_limit", nil)

	// content
	content := fmt.Sprintf(",\"order_price_type\":\"%s\"", orderPriceType)
	if pair != "" {
		content += fmt.Sprintf(",\"pair\": \"%s\"", pair)
	}
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if contractType != "" {
		content += fmt.Sprintf(",\"contract_type\": \"%s\"", contractType)
	}
	if businessType != "" {
		content += fmt.Sprintf(",\"business_type\": \"%s\"", businessType)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetOrderLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetOrderLimitResponse error: %s", jsonErr)
	}
	data <- result
}

// GetFeeAsync [General] Query information on swap trading fee
func (ac *AccountClient) GetFeeAsync(data chan account.GetFeeResponse, contractCode string, pair string, contractType string, businessType string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_fee", nil)

	// content
	content := fmt.Sprintf("{ \"contract_code\": \"%s\" }", contractCode)
	if pair != "" {
		content += fmt.Sprintf(",\"pair\": \"%s\"", pair)
	}
	if contractType != "" {
		content += fmt.Sprintf(",\"contract_type\": \"%s\"", contractType)
	}
	if businessType != "" {
		content += fmt.Sprintf(",\"business_type\": \"%s\"", businessType)
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetFeeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetFeeResponse error: %s", jsonErr)
	}
	data <- result
}

// IsolatedGetTransferLimitAsync [Isolated] Query information on Transfer Limit
func (ac *AccountClient) IsolatedGetTransferLimitAsync(data chan account.GetTransferLimitResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_transfer_limit", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetTransferLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetTransferLimitResponse error: %s", jsonErr)
	}
	data <- result
}

// CrossGetTransferLimitAsync [Cross] Query Information On Transfer Limit
func (ac *AccountClient) CrossGetTransferLimitAsync(data chan account.GetTransferLimitResponse, marginAccount string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_transfer_limit", nil)

	// content
	content := ""
	if marginAccount != "" {
		content += fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetTransferLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetTransferLimitResponse error: %s", jsonErr)
	}
	data <- result
}

// IsolatedGetPositionLimitAsync [Isolated] Query information on position limit
func (ac *AccountClient) IsolatedGetPositionLimitAsync(data chan account.GetPositionLimitResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_position_limit", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetPositionLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetPositionLimitResponse error: %s", jsonErr)
	}
	data <- result
}

// CrossGetPositionLimitAsync [Cross] Query Information On Position Limit
func (ac *AccountClient) CrossGetPositionLimitAsync(data chan account.GetPositionLimitResponse, contractCode string, pair string, contractType string, businessType string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_position_limit", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if pair != "" {
		content += fmt.Sprintf(",\"pair\": \"%s\"", pair)
	}
	if contractType != "" {
		content += fmt.Sprintf(",\"contract_type\": \"%s\"", contractType)
	}
	if businessType != "" {
		content += fmt.Sprintf(",\"business_type\": \"%s\"", businessType)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetPositionLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetPositionLimitResponse error: %s", jsonErr)
	}
	data <- result
}

// GetApiTradingStatusAsync [General] Query user's API indicator disable information
func (ac *AccountClient) GetApiTradingStatusAsync(data chan account.GetApiTradingStatusResponse) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.GET_METHOD, "/linear-swap-api/v1/swap_api_trading_status", nil)

	// content is nil
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetApiTradingStatusResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetApiTradingStatusResponse error: %s", jsonErr)
	}
	data <- result
}

// SwapSubPositionInfoAsync [Isolated] Query a single sub-account's position information
func (ac *AccountClient) SwapSubPositionInfoAsync(data chan account.SwapSubPositionInfoResponse, contractCode string, subUid string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_sub_position_info", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if subUid != "" {
		content += fmt.Sprintf(",\"sub_uid\": \"%s\"", subUid)
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

// SwapFinancialRecordAsync [General] Query account financial records(New)
func (ac *AccountClient) SwapFinancialRecordAsync(data chan account.SwapFinancialRecordResponse, contract string,
	marAcct string, fcType string, startTime string, endTime string, direct string, fromId string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v3/swap_financial_record", nil)

	// content
	content := ""
	if contract != "" {
		content += fmt.Sprintf(",\"contract\": \"%s\"", contract)
	}
	if marAcct != "" {
		content += fmt.Sprintf(",\"mar_acct\": \"%s\"", marAcct)
	}
	if fcType != "" {
		content += fmt.Sprintf(",\"type\": \"%s\"", fcType)
	}
	if startTime != "" {
		content += fmt.Sprintf(",\"start_time\": \"%s\"", startTime)
	}
	if endTime != "" {
		content += fmt.Sprintf(",\"end_time\": \"%s\"", endTime)
	}
	if direct != "" {
		content += fmt.Sprintf(",\"direct\": \"%s\"", direct)
	}
	if fromId != "" {
		content += fmt.Sprintf(",\"from_id\": \"%s\"", fromId)
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

// SwapFinancialRecordExactAsync [General]Query account financial records via Multiple Fields(New)
func (ac *AccountClient) SwapFinancialRecordExactAsync(data chan account.SwapFinancialRecordResponse, contract string,
	marAcct string, fcType string, startTime string, endTime string, direct string, fromId string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v3/swap_financial_record_exact", nil)

	// content
	content := ""
	if contract != "" {
		content += fmt.Sprintf(",\"contract\": \"%s\"", contract)
	}
	if marAcct != "" {
		content += fmt.Sprintf(",\"mar_acct\": \"%s\"", marAcct)
	}
	if fcType != "" {
		content += fmt.Sprintf(",\"type\": \"%s\"", fcType)
	}
	if startTime != "" {
		content += fmt.Sprintf(",\"start_time\": \"%s\"", startTime)
	}
	if endTime != "" {
		content += fmt.Sprintf(",\"end_time\": \"%s\"", endTime)
	}
	if direct != "" {
		content += fmt.Sprintf(",\"direct\": \"%s\"", direct)
	}
	if fromId != "" {
		content += fmt.Sprintf(",\"from_id\": \"%s\"", fromId)
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

// SwapLeverPositionLimitAsync [Isolated]Query Users' Position Limit for All Leverages
func (ac *AccountClient) SwapLeverPositionLimitAsync(data chan account.SwapLeverPositionLimitResponse, contractCode string,
	leverRate string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_lever_position_limit", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if leverRate != "" {
		content += fmt.Sprintf(",\"lever_rate\": \"%s\"", leverRate)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.SwapLeverPositionLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapLeverPositionLimitResponse error: %s", jsonErr)
	}
	data <- result
}

// SwapCrossLeverPositionLimitAsync [Cross]Query Users' Position Limit for All Leverages
func (ac *AccountClient) SwapCrossLeverPositionLimitAsync(data chan account.SwapCrossLeverPositionLimitResponse,
	businessType string, contractType string, pair string, contractCode string, leverRate string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_lever_position_limit", nil)

	// content
	content := ""
	if businessType != "" {
		content += fmt.Sprintf(",\"business_type\": \"%s\"", businessType)
	}
	if contractType != "" {
		content += fmt.Sprintf(",\"contract_type\": \"%s\"", contractType)
	}
	if pair != "" {
		content += fmt.Sprintf(",\"pair\t\": \"%s\"", pair)
	}
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if leverRate != "" {
		content += fmt.Sprintf(",\"lever_rate\": \"%s\"", leverRate)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.SwapCrossLeverPositionLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapCrossLeverPositionLimitResponse error: %s", jsonErr)
	}
	data <- result
}

// GetSwapSubAuthListAsync [General] Query sub-account transaction permissions
func (ac *AccountClient) GetSwapSubAuthListAsync(data chan account.GetSwapSubAuthListResponse, subUID string, startTime string,
	endTime string, direct string, fromID string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.GET_METHOD, "/linear-swap-api/v1/swap_sub_auth_list", nil)

	// option
	option := ""
	if subUID != "" {
		option += fmt.Sprintf("?sub_uid=%s", subUID)
	}
	if startTime != "" {
		option += fmt.Sprintf("&start_time=%s", startTime)
	}
	if endTime != "" {
		option += fmt.Sprintf("&end_time=%s", endTime)
	}
	if direct != "" {
		option += fmt.Sprintf("&direct=%s", direct)
	}
	if fromID != "" {
		option += fmt.Sprintf("&from_id=%s", fromID)
	}
	if option != "" {
		url += option
	}
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetSwapSubAuthListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapSubAuthListResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) GetAccountBalanceAsync(data chan account.GetAccountBalanceResponse) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.GET_METHOD, "/v5/account/balance", nil)

	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetAccountBalanceResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) SwapMultiAssetsMarginAsync(data chan account.SwapMultiAssetsMarginResponse, assetsMode int) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/v5/account/multi_assets_margin", nil)

	// content
	content := ""
	content += fmt.Sprintf(",\"assets_mode\": \"%s\"", assetsMode)
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.SwapMultiAssetsMarginResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}
