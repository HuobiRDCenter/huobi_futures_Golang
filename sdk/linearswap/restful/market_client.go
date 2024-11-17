package restful

import (
	"encoding/json"
	"fmt"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/response/market"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/reqbuilder"
)

type MarketClient struct {
	PUrlBuilder *reqbuilder.PublicUrlBuilder
}

func (mc *MarketClient) Init(host string) *MarketClient {
	if host == "" {
		host = linearswap.LINEAR_SWAP_DEFAULT_HOST
	}
	mc.PUrlBuilder = new(reqbuilder.PublicUrlBuilder).Init(host)
	return mc
}

// GetContractInfoAsync [General] Query Swap Info
func (mc *MarketClient) GetContractInfoAsync(data chan market.GetContractInfoResponse, contractCode string, supportMarginMode string, pair string, contractType string, businessType string) {
	// location
	location := "/linear-swap-api/v1/swap_contract_info"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if supportMarginMode != "" {
		option += fmt.Sprintf("&support_margin_mode=%s", supportMarginMode)
	}
	if pair != "" {
		option += fmt.Sprintf("&pair=%s", pair)
	}
	if contractType != "" {
		option += fmt.Sprintf("&contract_type=%s", contractType)
	}
	if businessType != "" {
		option += fmt.Sprintf("&business_type=%s", businessType)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetContractInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetContractInfoAsync error: %s", getErr)
	}
	data <- result
}

// GetIndexAsync [General] Query Swap Index Price Information
func (mc *MarketClient) GetIndexAsync(data chan market.GetIndexResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_index"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetIndexResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetIndexResponse error: %s", getErr)
	}
	data <- result
}

// GetPriceLimitAsync ([General] Query Swap Price Limitation)
func (mc *MarketClient) GetPriceLimitAsync(data chan market.GetPriceLimitResponse, contractCode string, pair string, contractType string, businessType string) {
	// location
	location := "/linear-swap-api/v1/swap_price_limit"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if pair != "" {
		option += fmt.Sprintf("&pair=%s", pair)
	}
	if contractType != "" {
		option += fmt.Sprintf("&contract_type=%s", contractType)
	}
	if businessType != "" {
		option += fmt.Sprintf("&business_type=%s", businessType)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetPriceLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetPriceLimitResponse error: %s", getErr)
	}
	data <- result
}

// GetOpenInterestAsync [General] Get Swap Open Interest Information
func (mc *MarketClient) GetOpenInterestAsync(data chan market.GetOpenInterestResponse, contractCode string, pair string, contractType string, businessType string) {
	// location
	location := "/linear-swap-api/v1/swap_open_interest"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if pair != "" {
		option += fmt.Sprintf("&pair=%s", pair)
	}
	if contractType != "" {
		option += fmt.Sprintf("&contract_type=%s", contractType)
	}
	if businessType != "" {
		option += fmt.Sprintf("&business_type=%s", businessType)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetOpenInterestResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetOpenInterestResponse error: %s", getErr)
	}
	data <- result
}

// GetDepthAsync [General] Get Market Depth
func (mc *MarketClient) GetDepthAsync(data chan market.GetDepthResponse, contractCode string, fcType string) {
	// location
	location := fmt.Sprintf("/linear-swap-ex/market/depth?contract_code=%s&type=%s", contractCode, fcType)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetDepthResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetDepthResponse error: %s", getErr)
	}
	data <- result
}

// GetBboAsync [General]Get Market BBO Data
func (mc *MarketClient) GetBboAsync(data chan market.GetBboResponse, contractCode string, businessType string) {
	// location
	location := "/linear-swap-ex/market/bbo"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if businessType != "" {
		option += fmt.Sprintf("&business_type=%s", businessType)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetBboResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetBboResponse error: %s", getErr)
	}
	data <- result
}

// GetKLineAsync ([General] Get KLine Data)
func (mc *MarketClient) GetKLineAsync(data chan market.GetKLineResponse, contractCode string, period string, size int, from int, to int) {
	// location
	location := fmt.Sprintf("/linear-swap-ex/market/history/kline?contract_code=%s&period=%s", contractCode, period)

	// option
	option := ""
	if size != 0 {
		option += fmt.Sprintf("&size=%d", size)
	}
	if from != 0 {
		option += fmt.Sprintf("&from=%d", from)
	}
	if to != 0 {
		option += fmt.Sprintf("&to=%d", to)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetKLineResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetKLineResponse error: %s", getErr)
	}
	data <- result
}

// GetMarkPriceKLineAsync [General]Get Kline Data of Mark Price
func (mc *MarketClient) GetMarkPriceKLineAsync(data chan market.GetStrKLineResponse, contractCode string, period string, size int) {
	// location
	location := fmt.Sprintf("/index/market/history/linear_swap_mark_price_kline?contract_code=%s&period=%s&size=%d", contractCode, period, size)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetStrKLineResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetStrKLineResponse error: %s", getErr)
	}
	data <- result
}

// GetMergedAsync [General] Get Market Data Overview
func (mc *MarketClient) GetMergedAsync(data chan market.GetMergedResponse, contractCode string) {
	// location
	location := fmt.Sprintf("/linear-swap-ex/market/detail/merged?contract_code=%s", contractCode)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetMergedResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetMergedResponse error: %s", getErr)
	}
	data <- result
}

// GetBatchMergedAsync [General]Get a Batch of Market Data Overview
func (mc *MarketClient) GetBatchMergedAsync(data chan market.GetBatchMergedResponse, contractCode string, businessType string) {
	// location
	location := fmt.Sprintf("/linear-swap-ex/market/detail/batch_merged")

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if businessType != "" {
		option += fmt.Sprintf("&business_type=%s", businessType)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetBatchMergedResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetBatchMergedAsync error: %s", getErr)
	}
	data <- result
}

// GetTradeAsync [General] Query The Last Trade of a Contract
func (mc *MarketClient) GetTradeAsync(data chan market.GetTradeResponse, contractCode string, businessType string) {
	// location
	location := "/linear-swap-ex/market/trade"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if businessType != "" {
		option += fmt.Sprintf("&business_type=%s", businessType)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetTradeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetTradeResponse error: %s", getErr)
	}
	data <- result
}

// GetHisTradeAsync [General] Query a Batch of Trade Records of a Contract
func (mc *MarketClient) GetHisTradeAsync(data chan market.GetHisTradeResponse, contractCode string, size int) {
	// location
	location := fmt.Sprintf("/linear-swap-ex/market/history/trade?contract_code=%s&size=%d", contractCode, size)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetHisTradeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHisTradeResponse error: %s", getErr)
	}
	data <- result
}

// GetRiskInfoAsync [General]Get Market BBO Data
func (mc *MarketClient) GetRiskInfoAsync(data chan market.GetRiskInfoResponse, contractCode string, businessType string) {
	// location
	location := "/linear-swap-api/v1/swap_risk_info"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if businessType != "" {
		option += fmt.Sprintf("&business_type=%s", businessType)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetRiskInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetRiskInfoResponse error: %s", getErr)
	}
	data <- result
}

// GetInsuranceFundAsync [General] Query history records of insurance fund balance
func (mc *MarketClient) GetInsuranceFundAsync(data chan market.GetInsuranceFundResponse, contractCode string, pageIndex int, pageSize int) {
	// location
	location := fmt.Sprintf("/linear-swap-api/v1/swap_insurance_fund?contract_code=%s", contractCode)

	// option
	option := ""
	if pageIndex != 0 {
		option += fmt.Sprintf("&size=%d", pageIndex)
	}
	if pageSize != 0 {
		option += fmt.Sprintf("&from=%d", pageSize)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetInsuranceFundResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetInsuranceFundResponse error: %s", getErr)
	}
	data <- result
}

// IsolatedGetAdjustFactorFundAsync [Isolated] Query information on Tiered Adjustment Factor
func (mc *MarketClient) IsolatedGetAdjustFactorFundAsync(data chan market.GetAdjustFactorFundResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_adjustfactor"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetAdjustFactorFundResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAdjustFactorFundResponse error: %s", getErr)
	}
	data <- result
}

// CrossGetAdjustFactorFundAsync [Cross] Query Information On Tiered Adjustment Factor
func (mc *MarketClient) CrossGetAdjustFactorFundAsync(data chan market.GetAdjustFactorFundResponse, contractCode string, pair string, contractType string, businessType string) {
	// location
	location := "/linear-swap-api/v1/swap_cross_adjustfactor"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if pair != "" {
		option += fmt.Sprintf("&pair=%s", pair)
	}
	if contractType != "" {
		option += fmt.Sprintf("&contract_type=%s", contractType)
	}
	if businessType != "" {
		option += fmt.Sprintf("&business_type=%s", businessType)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetAdjustFactorFundResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAdjustFactorFundResponse error: %s", getErr)
	}
	data <- result
}

// GetHisOpenInterestAsync [General] Query information on open interest
func (mc *MarketClient) GetHisOpenInterestAsync(data chan market.GetHisOpenInterestResponse, contractCode string, period string, amountType int, size int, pair string, contractType string) {
	// location
	location := fmt.Sprintf("/linear-swap-api/v1/swap_his_open_interest?contract_code=%s&period=%s&amount_type=%d", contractCode, period, amountType)

	// option
	option := ""
	if size != 0 {
		option += fmt.Sprintf("&size=%d", size)
	}
	if pair != "" {
		option += fmt.Sprintf("&pair=%s", pair)
	}
	if contractType != "" {
		option += fmt.Sprintf("&contract_type=%s", contractType)
	}

	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetHisOpenInterestResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHisOpenInterestResponse error: %s", getErr)
	}
	data <- result
}

// IsolatedGetLadderMarginAsync [Isolated]Query information on Tiered Margin
func (mc *MarketClient) IsolatedGetLadderMarginAsync(data chan market.GetLadderMarginResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_ladder_margin"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetLadderMarginResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetLadderMarginResponse error: %s", getErr)
	}
	data <- result
}

// CrossGetLadderMarginAsync ([Cross]Query information on Tiered Margin)
func (mc *MarketClient) CrossGetLadderMarginAsync(data chan market.GetLadderMarginResponse, contractCode string, pair string, contractType string, businessType string) {
	// location
	location := "/linear-swap-api/v1/swap_cross_ladder_margin"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if pair != "" {
		option += fmt.Sprintf("&pair=%s", pair)
	}
	if contractType != "" {
		option += fmt.Sprintf("&contract_type=%s", contractType)
	}
	if businessType != "" {
		option += fmt.Sprintf("&business_type=%s", businessType)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetLadderMarginResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetLadderMarginResponse error: %s", getErr)
	}
	data <- result
}

// GetEliteAccountRatioAsync [General] Query Top Trader Sentiment Index Function-Account)
func (mc *MarketClient) GetEliteAccountRatioAsync(data chan market.GetEliteRatioResponse, contractCode string, period string) {
	// location
	location := fmt.Sprintf("/linear-swap-api/v1/swap_elite_account_ratio?contract_code=%s&period=%s", contractCode, period)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetEliteRatioResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetElitePositionRatioResponse error: %s", getErr)
	}
	data <- result
}

// GetElitePositionRatioAsync ([General] Query Top Trader Sentiment Index Function-Position)
func (mc *MarketClient) GetElitePositionRatioAsync(data chan market.GetEliteRatioResponse, contractCode string, period string) {
	// location
	location := fmt.Sprintf("/linear-swap-api/v1/swap_elite_position_ratio?contract_code=%s&period=%s", contractCode, period)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetEliteRatioResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetElitePositionRatioResponse error: %s", getErr)
	}
	data <- result
}

// IsolatedGetApiStateAsync ([Isolated] Query information on system status)
func (mc *MarketClient) IsolatedGetApiStateAsync(data chan market.GetApiStateResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_api_state"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetApiStateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetApiStateResponse error: %s", getErr)
	}
	data <- result
}

// CrossGetTransferStateAsync ([Cross] Query Information On Transfer State)
func (mc *MarketClient) CrossGetTransferStateAsync(data chan market.GetTransferStateResponse, marginAccount string) {
	// location
	location := "/linear-swap-api/v1/swap_cross_transfer_state"

	// option
	if marginAccount == "" {
		marginAccount = "USDT"
	}
	location += fmt.Sprintf("?margin_account=%s", marginAccount)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetTransferStateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetTransferStateResponse error: %s", getErr)
	}
	data <- result
}

// CrossGetTradeStateAsync ([Cross] Query Information On Trade State)
func (mc *MarketClient) CrossGetTradeStateAsync(data chan market.GetTradeStateResponse, contractCode string, pair string, contractType string, businessType string) {
	// location
	location := "/linear-swap-api/v1/swap_cross_trade_state"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if pair != "" {
		option += fmt.Sprintf("&pair=%s", pair)
	}
	if contractType != "" {
		option += fmt.Sprintf("&contract_type=%s", contractType)
	}
	if businessType != "" {
		option += fmt.Sprintf("&business_type=%s", businessType)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetTradeStateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetTradeStatusResponse error: %s", getErr)
	}
	data <- result
}

// GetFundingRateAsync ([General] Query funding rate)
func (mc *MarketClient) GetFundingRateAsync(data chan market.GetFundingRateResponse, contractCode string) {
	// location
	location := fmt.Sprintf("/linear-swap-api/v1/swap_funding_rate?contract_code=%s", contractCode)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetFundingRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetFundingRateResponse error: %s", getErr)
	}
	data <- result
}

// GetBatchFundingRateAsync ([General]Query a Batch of Funding Rate)
func (mc *MarketClient) GetBatchFundingRateAsync(data chan market.GetBatchFundingRateResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_batch_funding_rate"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetBatchFundingRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetBatchFundingRateResponse error: %s", getErr)
	}
	data <- result
}

// GetHisFundingRateAsync ([General] Query historical funding rate)
func (mc *MarketClient) GetHisFundingRateAsync(data chan market.GetHisFundingRateResponse, contractCode string, pageIndex int, pageSize int) {
	// location
	location := fmt.Sprintf("/linear-swap-api/v1/swap_historical_funding_rate?contract_code=%s", contractCode)

	// option
	option := ""
	if pageIndex != 0 {
		option += fmt.Sprintf("&page_index=%d", pageIndex)
	}
	if pageSize != 0 {
		option += fmt.Sprintf("&page_size=%d", pageSize)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetHisFundingRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHisFundingRateResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetLiquidationOrdersAsync(data chan market.GetLiquidationOrdersResponse, contractCode string, tradeType int, createDate int,
	pageIndex int, pageSize int) {
	// location
	location := fmt.Sprintf("/linear-swap-api/v1/swap_liquidation_orders?contract_code=%s&trade_type=%d&create_date=%d", contractCode, tradeType, createDate)

	// option
	option := ""
	if pageIndex != 0 {
		option += fmt.Sprintf("&page_index=%d", pageIndex)
	}
	if pageSize != 0 {
		option += fmt.Sprintf("&page_size=%d", pageSize)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetLiquidationOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetLiquidationOrdersResponse error: %s", getErr)
	}
	data <- result
}

// GetPremiumIndexKLineAsync ([General] Query Premium Index Kline Data)
func (mc *MarketClient) GetPremiumIndexKLineAsync(data chan market.GetStrKLineResponse, contractCode string, period string, size int) {
	// location
	location := fmt.Sprintf("/index/market/history/linear_swap_premium_index_kline?contract_code=%s&period=%s&size=%d", contractCode, period, size)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetStrKLineResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetPriceLimitResponse error: %s", getErr)
	}
	data <- result
}

// GetEstimatedRateKLineAsync ([General] Query Estimated Funding Rate Kline Data)
func (mc *MarketClient) GetEstimatedRateKLineAsync(data chan market.GetStrKLineResponse, contractCode string, period string, size int) {
	// location
	location := fmt.Sprintf("/index/market/history/linear_swap_estimated_rate_kline?contract_code=%s&period=%s&size=%d", contractCode, period, size)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetStrKLineResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetStrKLineResponse error: %s", getErr)
	}
	data <- result
}

// GetBasisAsync ([General] Query Basis Data)
func (mc *MarketClient) GetBasisAsync(data chan market.GetBasisResponse, contractCode string, period string, size int, basisPriceType string) {
	// location
	location := fmt.Sprintf("/index/market/history/linear_swap_basis?contract_code=%s&period=%s&size=%d", contractCode, period, size)

	// option
	option := ""
	if basisPriceType != "" {
		option += fmt.Sprintf("&basis_price_type=%s", basisPriceType)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetBasisResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetBasisResponse error: %s", getErr)
	}
	data <- result
}

// GetEstimatedSettlementPriceAsync ([General]Get the estimated settlement price)
func (mc *MarketClient) GetEstimatedSettlementPriceAsync(data chan market.GetEstimatedSettlementPriceResponse, contractCode string, pair string, contractType string, businessType string) {
	// location
	location := "/linear-swap-api/v1/swap_estimated_settlement_price"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if pair != "" {
		option += fmt.Sprintf("&pair=%s", pair)
	}
	if contractType != "" {
		option += fmt.Sprintf("&contract_type=%s", contractType)
	}
	if businessType != "" {
		option += fmt.Sprintf("&business_type=%s", businessType)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetEstimatedSettlementPriceResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetEstimatedSettlementPriceResponse error: %s", getErr)
	}
	data <- result
}

// GetSwapLiquidationOrdersAsync ([General] Query Liquidation Orders(New))
func (mc *MarketClient) GetSwapLiquidationOrdersAsync(data chan market.GetSwapLiquidationOrdersResponse, contract string,
	pair string, tradeType string, startTime string, endTime string, direct string, fromId string) {
	// location
	location := "/linear-swap-api/v3/swap_liquidation_orders"

	// option
	option := ""
	if contract != "" {
		option += fmt.Sprintf("?contract=%s", contract)
	}
	if pair != "" {
		option += fmt.Sprintf("&pair=%s", pair)
	}
	if tradeType != "" {
		option += fmt.Sprintf("&trade_type=%s", tradeType)
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
	if fromId != "" {
		option += fmt.Sprintf("&from_id=%s", fromId)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetSwapLiquidationOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapLiquidationOrdersResponse error: %s", getErr)
	}
	data <- result
}

// GetSwapSettlementRecordsAsync ([General] Query historical settlement records of the platform interface)
func (mc *MarketClient) GetSwapSettlementRecordsAsync(data chan market.GetSwapSettlementRecordsResponse, contractCode string,
	startTime string, endTime string, pageIndex string, pageSize string) {
	// location
	location := "/linear-swap-api/v1/swap_settlement_records"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if startTime != "" {
		option += fmt.Sprintf("&start_time=%s", startTime)
	}
	if endTime != "" {
		option += fmt.Sprintf("&end_time=%s", endTime)
	}
	if pageIndex != "" {
		option += fmt.Sprintf("&page_index=%s", pageIndex)
	}
	if pageSize != "" {
		option += fmt.Sprintf("&page_size=%s", pageSize)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetSwapSettlementRecordsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapSettlementRecordsResponse error: %s", getErr)
	}
	data <- result
}

// GetSwapQueryElementsAsync (Contract Elements)
func (mc *MarketClient) GetSwapQueryElementsAsync(data chan market.GetSwapQueryElementsResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_query_elements"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetSwapQueryElementsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapQueryElementsResponse error: %s", getErr)
	}
	data <- result
}

// GetTimestampAsync (Get current system timestamp)
func (mc *MarketClient) GetTimestampAsync(data chan market.GetTimestampResponse) {
	// location
	location := "/api/v1/timestamp"
	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetTimestampResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetTimestampResponse error: %s", getErr)
	}
	data <- result
}

// GetHeartbeatAsync (Query whether the system is available)
func (mc *MarketClient) GetHeartbeatAsync(data chan market.GetHeartbeatResponse) {
	// location
	location := "/heartbeat/"
	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetHeartbeatResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHeartbeatResponse error: %s", getErr)
	}
	data <- result
}

// GetSummaryAsync (Get system status)
func (mc *MarketClient) GetSummaryAsync(data chan market.GetSummaryResponse) {
	// location
	url := "https://status-linear-swap.huobigroup.com/api/v2/summary.json"
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetSummaryResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSummaryResponse error: %s", getErr)
	}
	data <- result
}

// GetBatchMergedV2Async ([General]Get a Batch of Market Data Overview(V2))
func (mc *MarketClient) GetBatchMergedV2Async(data chan market.GetBatchMergedV2Response, contractCode string, businessType string) {
	// location
	location := "/v2/linear-swap-ex/market/detail/batch_merged"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if businessType != "" {
		option += fmt.Sprintf("&business_type=%s", businessType)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetBatchMergedV2Response{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetBatchMergedV2Response error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetMarketRiskLimitAsync(data chan market.GetMarketRiskLimitResponse, contractCode string,
	contractType string, marginMode string, tier string) {
	// location
	location := "/v5/market/risk_limit"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if contractType != "" {
		option += fmt.Sprintf("&contract_type=%s", contractType)
	}
	if marginMode != "" {
		option += fmt.Sprintf("&margin_mode=%s", marginMode)
	}
	if tier != "" {
		option += fmt.Sprintf("&tier=%s", tier)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetMarketRiskLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetMarketAssetsDeductionCurrencyAsync(data chan market.GetMarketAssetsDeductionCurrencyResponse) {
	// location
	location := "/v5/assets_deduction_currency"

	// option
	option := ""
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetMarketAssetsDeductionCurrencyResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetMarketMultiAssetsMarginListAsync(data chan market.GetMarketMultiAssetsMarginListResponse) {
	// location
	location := "/v5/market/multi_assets_margin/list"

	// option
	option := ""
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetMarketMultiAssetsMarginListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", getErr)
	}
	data <- result
}
