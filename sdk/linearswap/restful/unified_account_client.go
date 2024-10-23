package restful

import (
	"encoding/json"
	"fmt"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/response/unifiedaccount"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/reqbuilder"
)

type UnifiedAccountClient struct {
	PUrlBuilder *reqbuilder.PrivateUrlBuilder
}

func (uac *UnifiedAccountClient) Init(accessKey string, secretKey string, host string, sign string) *UnifiedAccountClient {
	if host == "" {
		host = linearswap.LINEAR_SWAP_DEFAULT_HOST
	}
	uac.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host, sign)
	return uac
}

// GetUnifiedAccountInfoAsync (Query unified account assets)
func (uac *UnifiedAccountClient) GetUnifiedAccountInfoAsync(data chan unifiedaccount.GetUnifiedAccountInfoResponse, contractCode string) {
	// ulr
	url := uac.PUrlBuilder.Build(linearswap.GET_METHOD, "/linear-swap-api/v3/unified_account_info", nil)
	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if option != "" {
		url += fmt.Sprintf("?%s", option)
	}
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := unifiedaccount.GetUnifiedAccountInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetUnifiedAccountInfoResponse error: %s", jsonErr)
	}
	data <- result
}

// GetLinearSwapOverviewAccountInfoAsync (Deductible asset inquiry )
func (uac *UnifiedAccountClient) GetLinearSwapOverviewAccountInfoAsync(data chan unifiedaccount.GetLinearSwapOverviewAccountInfoResponse,
	tradePartition string) {
	// ulr
	url := uac.PUrlBuilder.Build(linearswap.GET_METHOD, "/linear-swap-api/v3/linear_swap_overview_account_info", nil)
	// option
	option := ""
	if tradePartition != "" {
		option += fmt.Sprintf("trade_partition=%s", tradePartition)
	}
	if option != "" {
		url += fmt.Sprintf("?%s", option)
	}
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := unifiedaccount.GetLinearSwapOverviewAccountInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetLinearSwapOverviewAccountInfoResponse error: %s", jsonErr)
	}
	data <- result
}

// LinearSwapFeeSwitchAsync (Deductible asset inquiry )
func (uac *UnifiedAccountClient) LinearSwapFeeSwitchAsync(data chan unifiedaccount.LinearSwapFeeSwitchResponse,
	feeOption string, deductionCurrency string) {
	// ulr
	url := uac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v3/linear_swap_fee_switch", nil)

	// content
	content := ""
	if feeOption != "" {
		content += fmt.Sprintf(",\"fee_option\": \"%s\"", feeOption)
	}
	if deductionCurrency != "" {
		content += fmt.Sprintf(",\"deduction_currency\": \"%s\"", deductionCurrency)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := unifiedaccount.LinearSwapFeeSwitchResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to LinearSwapFeeSwitchResponse error: %s", jsonErr)
	}
	data <- result
}

// FixPositionMarginChangeAsync (Adjust margin for isolated positions)
func (uac *UnifiedAccountClient) FixPositionMarginChangeAsync(data chan unifiedaccount.FixPositionMarginChangeResponse,
	amount string, asset string, contractCode string, fcType string, direction string, clientOrderId string) {
	// ulr
	url := uac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v3/fix_position_margin_change", nil)

	// content
	content := ""
	if amount != "" {
		content += fmt.Sprintf(",\"amount\": \"%s\"", amount)
	}
	if asset != "" {
		content += fmt.Sprintf(",\"asset\": \"%s\"", asset)
	}
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if fcType != "" {
		content += fmt.Sprintf(",\"type\": \"%s\"", fcType)
	}
	if direction != "" {
		content += fmt.Sprintf(",\"direction\": \"%s\"", direction)
	}
	if clientOrderId != "" {
		content += fmt.Sprintf(",\"client_order_id\": \"%s\"", clientOrderId)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := unifiedaccount.FixPositionMarginChangeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to FixPositionMarginChangeResponse error: %s", jsonErr)
	}
	data <- result
}

// FixPositionMarginChangeRecordAsync (Query the margin adjustment records of isolated positions )
func (uac *UnifiedAccountClient) FixPositionMarginChangeRecordAsync(data chan unifiedaccount.FixPositionMarginChangeRecordResponse,
	asset string, contractCode string, startTime string, endTime string, direct string, fromId string) {
	// ulr
	url := uac.PUrlBuilder.Build(linearswap.GET_METHOD, "/linear-swap-api/v3/fix_position_margin_change_record", nil)
	// option
	option := ""
	if asset != "" {
		option += fmt.Sprintf("asset=%s", asset)
	}
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if startTime != "" {
		option += fmt.Sprintf("start_time=%s", startTime)
	}
	if endTime != "" {
		option += fmt.Sprintf("end_time=%s", endTime)
	}
	if direct != "" {
		option += fmt.Sprintf("direct=%s", direct)
	}
	if fromId != "" {
		option += fmt.Sprintf("from_id=%s", fromId)
	}
	if option != "" {
		url += fmt.Sprintf("?%s", option)
	}
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := unifiedaccount.FixPositionMarginChangeRecordResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to FixPositionMarginChangeRecordResponse error: %s", jsonErr)
	}
	data <- result
}

// SwapUnifiedAccountTypeAsync (Account type query)
func (uac *UnifiedAccountClient) SwapUnifiedAccountTypeAsync(data chan unifiedaccount.SwapUnifiedAccountTypeResponse) {
	// ulr
	url := uac.PUrlBuilder.Build(linearswap.GET_METHOD, "/linear-swap-api/v3/swap_unified_account_type", nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := unifiedaccount.SwapUnifiedAccountTypeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapUnifiedAccountTypeResponse error: %s", jsonErr)
	}
	data <- result
}

// SwapSwitchAccountTypeAsync (Account Type Change)
func (uac *UnifiedAccountClient) SwapSwitchAccountTypeAsync(data chan unifiedaccount.SwapSwitchAccountTypeResponse,
	accountType string) {
	// ulr
	url := uac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v3/swap_switch_account_type", nil)

	// content
	content := ""
	if accountType != "" {
		content += fmt.Sprintf(",\"account_type\": \"%s\"", accountType)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := unifiedaccount.SwapSwitchAccountTypeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapSwitchAccountTypeResponse error: %s", jsonErr)
	}
	data <- result
}
