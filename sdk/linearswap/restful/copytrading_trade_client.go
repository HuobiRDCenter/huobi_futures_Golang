package restful

import (
    "encoding/json"
    "fmt"
    "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap"
    "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/response/copytradingtrade"
    "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
    "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/reqbuilder"
)

type CopyTradingTradeClient struct {
    PUrlBuilder *reqbuilder.PrivateUrlBuilder
}

func (ctc *CopyTradingTradeClient) Init(accessKey string, secretKey string, host string) *CopyTradingTradeClient {
    if host == "" {
        host = linearswap.LINEAR_SWAP_DEFAULT_HOST
    }
    ctc.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
    return ctc
}

// QueryTraderInstrumentsAsync 查询交易员可带单合约
func (ctc *CopyTradingTradeClient) QueryTraderInstrumentsAsync(data chan copytradingtrade.BaseArrayResponse[copytradingtrade.InstrumentResponse], instType string) {
	// ulr
	url := ctc.PUrlBuilder.Build(linearswap.GET_METHOD, "/api/v6/copyTrading/trader/instruments", nil)

	// option
	option := "?"
	if instType != "" {
		option += fmt.Sprintf("&instType=%s", instType)
	}
	if option != "" {
		url += option
	}
	
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := copytradingtrade.BaseArrayResponse[copytradingtrade.InstrumentResponse]{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

// QueryTraderStatisticsAsync 查询交易员统计数据
func (ctc *CopyTradingTradeClient) QueryTraderStatisticsAsync(data chan copytradingtrade.BaseArrayResponse[copytradingtrade.StatisticsResponse], instType string) {
	// ulr
	url := ctc.PUrlBuilder.Build(linearswap.GET_METHOD, "/api/v6/copyTrading/trader/statistics", nil)

	// option
	option := "?"
	if instType != "" {
		option += fmt.Sprintf("&instType=%s", instType)
	}
	if option != "" {
		url += option
	}
	
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := copytradingtrade.BaseArrayResponse[copytradingtrade.StatisticsResponse]{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

// QueryTraderProfitSharingHistoryAsync 查询交易员分润历史
func (ctc *CopyTradingTradeClient) QueryTraderProfitSharingHistoryAsync(data chan copytradingtrade.BaseArrayResponse[copytradingtrade.ProfitSharingHistoryResponse], 
    instId string, instType string, begin string, end string, after string, before string, limit string) {
	// ulr
	url := ctc.PUrlBuilder.Build(linearswap.GET_METHOD, "/api/v6/copyTrading/trader/profit-sharing-history", nil)

	// option
	option := "?"
	if instId != "" {
		option += fmt.Sprintf("&instId=%s", instId)
	}
	if instType != "" {
		option += fmt.Sprintf("&instType=%s", instType)
	}
	if begin != "" {
		option += fmt.Sprintf("&begin=%s", begin)
	}
	if end != "" {
		option += fmt.Sprintf("&end=%s", end)
	}
	if after != "" {
		option += fmt.Sprintf("&after=%s", after)
	}
	if before != "" {
		option += fmt.Sprintf("&before=%s", before)
	}
	if limit != "" {
		option += fmt.Sprintf("&limit=%s", limit)
	}
	if option != "" {
		url += option
	}
	
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := copytradingtrade.BaseArrayResponse[copytradingtrade.ProfitSharingHistoryResponse]{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

// QueryTraderProfitSharingHistorySummaryAsync 查询交易员分润历史汇总
func (ctc *CopyTradingTradeClient) QueryTraderProfitSharingHistorySummaryAsync(data chan copytradingtrade.BaseArrayResponse[copytradingtrade.ProfitSharingHistorySummaryResponse], instType string) {
	// ulr
	url := ctc.PUrlBuilder.Build(linearswap.GET_METHOD, "/api/v6/copyTrading/trader/profit-sharing-history-summary", nil)

	// option
	option := "?"
	if instType != "" {
		option += fmt.Sprintf("&instType=%s", instType)
	}
	if option != "" {
		url += option
	}
	
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := copytradingtrade.BaseArrayResponse[copytradingtrade.ProfitSharingHistorySummaryResponse]{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

// QueryTraderUnrealizedProfitSharingSummaryAsync 查询交易员待分润汇总
func (ctc *CopyTradingTradeClient) QueryTraderUnrealizedProfitSharingSummaryAsync(data chan copytradingtrade.BaseArrayResponse[copytradingtrade.UnrealizedProfitSharingSummaryResponse], 
    instType string, ccy string) {
	// ulr
	url := ctc.PUrlBuilder.Build(linearswap.GET_METHOD, "/api/v6/copyTrading/trader/unrealized-profit-sharing-summary", nil)

	// option
	option := "?"
	if instType != "" {
		option += fmt.Sprintf("&instType=%s", instType)
	}
	if ccy != "" {
		option += fmt.Sprintf("&ccy=%s", ccy)
	}
	if option != "" {
		url += option
	}
	
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := copytradingtrade.BaseArrayResponse[copytradingtrade.UnrealizedProfitSharingSummaryResponse]{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

// QueryTraderFollowersAsync 查询交易员跟随者列表
func (ctc *CopyTradingTradeClient) QueryTraderFollowersAsync(data chan copytradingtrade.BaseArrayResponse[copytradingtrade.FollowersResponse], 
    instType string, begin string, end string, after string, before string, limit string) {
	// ulr
	url := ctc.PUrlBuilder.Build(linearswap.GET_METHOD, "/api/v6/copyTrading/trader/followers", nil)

	// option
	option := "?"
	if instType != "" {
		option += fmt.Sprintf("&instType=%s", instType)
	}
	if begin != "" {
		option += fmt.Sprintf("&begin=%s", begin)
	}
	if end != "" {
		option += fmt.Sprintf("&end=%s", end)
	}
	if after != "" {
		option += fmt.Sprintf("&after=%s", after)
	}
	if before != "" {
		option += fmt.Sprintf("&before=%s", before)
	}
	if limit != "" {
		option += fmt.Sprintf("&limit=%s", limit)
	}
	if option != "" {
		url += option
	}
	
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := copytradingtrade.BaseArrayResponse[copytradingtrade.FollowersResponse]{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

// DeleteTraderFollowerAsync 移除交易员跟随者
func (ctc *CopyTradingTradeClient) DeleteTraderFollowerAsync(data chan copytradingtrade.BaseArrayResponse[copytradingtrade.DeleteFollowerResponse], 
    instType string, followerUids []string) {
	// ulr
	url := ctc.PUrlBuilder.Build(linearswap.POST_METHOD, "/api/v6/copyTrading/trader/follower", nil)

	// 构建请求内容
	content := ""
	if instType != "" {
		content += fmt.Sprintf(",\"instType\":\"%s\"", instType)
	}
	if followerUids != nil && len(followerUids) > 0 {
		followerUidsJSON, _ := json.Marshal(followerUids)
		content += fmt.Sprintf(",\"followerUids\":%s", string(followerUidsJSON))
	}
	
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http post error: %s", getErr)
	}
	result := copytradingtrade.BaseArrayResponse[copytradingtrade.DeleteFollowerResponse]{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

// TraderTransferAsync 交易员账户划转
func (ctc *CopyTradingTradeClient) TraderTransferAsync(data chan copytradingtrade.BaseArrayResponse[copytradingtrade.TransferResponse], 
    amt string, from string, to string, ccy string) {
	// ulr
	url := ctc.PUrlBuilder.Build(linearswap.POST_METHOD, "/api/v6/copyTrading/trader/transfer", nil)

	// 构建请求内容
	content := ""
	if amt != "" {
		content += fmt.Sprintf(",\"amt\":\"%s\"", amt)
	}
	if from != "" {
		content += fmt.Sprintf(",\"from\":\"%s\"", from)
	}
	if to != "" {
		content += fmt.Sprintf(",\"to\":\"%s\"", to)
	}
	if ccy != "" {
		content += fmt.Sprintf(",\"ccy\":\"%s\"", ccy)
	}
	
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http post error: %s", getErr)
	}
	result := copytradingtrade.BaseArrayResponse[copytradingtrade.TransferResponse]{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

// TraderFollowerSettingsAsync 交易员跟单设置
func (ctc *CopyTradingTradeClient) TraderFollowerSettingsAsync(data chan copytradingtrade.BaseArrayResponse[copytradingtrade.FollowerSettingsResponse], 
    instType string, enable *bool, profitSharingRatio string, maxFollowers string) {
	// ulr
	url := ctc.PUrlBuilder.Build(linearswap.POST_METHOD, "/api/v6/copyTrading/trader/follower-settings", nil)

	// 构建请求内容
	content := ""
	if instType != "" {
		content += fmt.Sprintf(",\"instType\":\"%s\"", instType)
	}
	if enable != nil {
		content += fmt.Sprintf(",\"enable\":%v", *enable)
	}
	if profitSharingRatio != "" {
		content += fmt.Sprintf(",\"profitSharingRatio\":\"%s\"", profitSharingRatio)
	}
	if maxFollowers != "" {
		content += fmt.Sprintf(",\"maxFollowers\":\"%s\"", maxFollowers)
	}
	
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http post error: %s", getErr)
	}
	result := copytradingtrade.BaseArrayResponse[copytradingtrade.FollowerSettingsResponse]{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

// TraderConfigAsync 查询交易员配置
func (ctc *CopyTradingTradeClient) TraderConfigAsync(data chan copytradingtrade.BaseArrayResponse[copytradingtrade.ConfigResponse], instType string) {
	// ulr
	url := ctc.PUrlBuilder.Build(linearswap.GET_METHOD, "/api/v6/copyTrading/trader/config", nil)

	// option
	option := "?"
	if instType != "" {
		option += fmt.Sprintf("&instType=%s", instType)
	}
	if option != "" {
		url += option
	}
	
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := copytradingtrade.BaseArrayResponse[copytradingtrade.ConfigResponse]{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

// TraderApikeyAsync 生成交易员API密钥
func (ctc *CopyTradingTradeClient) TraderApikeyAsync(data chan copytradingtrade.BaseArrayResponse[copytradingtrade.ApikeyResponse], instType string) {
	// ulr
	url := ctc.PUrlBuilder.Build(linearswap.POST_METHOD, "/api/v6/copyTrading/trader/apikey", nil)

	// 构建请求内容
	content := ""
	if instType != "" {
		content += fmt.Sprintf(",\"instType\":\"%s\"", instType)
	}
	
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http post error: %s", getErr)
	}
	result := copytradingtrade.BaseArrayResponse[copytradingtrade.ApikeyResponse]{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}