package restful

import (
	"encoding/json"
	"fmt"

	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap"
	requestcopytrading "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/request/copytrading"
	responsecopytrading "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/response/copytrading"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/reqbuilder"
)

type CopytradingClient struct {
	PUrlBuilder *reqbuilder.PrivateUrlBuilder
}

func (oc *CopytradingClient) Init(accessKey string, secretKey string, host string) *CopytradingClient {
	if host == "" {
		host = linearswap.LINEAR_SWAP_DEFAULT_HOST
	}
	oc.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return oc
}

func (oc *CopytradingClient) CopytradingTraderPlaceOrder(data chan responsecopytrading.CopytradingTraderPlaceOrderResponse, request requestcopytrading.CopytradingTraderPlaceOrderRequest) {
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/copytrading/trader/place_order", nil)
	content, err := json.Marshal(request)
	if err != nil {
		log.Error("error: %v", err)
	}
	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsecopytrading.CopytradingTraderPlaceOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", getErr)
	}
	data <- result
}

func (ac *CopytradingClient) CopytradingTraderQueryContractLever(data chan responsecopytrading.CopytradingTraderQueryContractLeverResponse,
	contractCode string, marginMode string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.GET_METHOD, "/copytrading/trader/query_contract_lever", nil)
	// option
	option := "?"
	if contractCode != "" {
		option += fmt.Sprintf("&contract_code=%s", contractCode)
	}
	if marginMode != "" {
		option += fmt.Sprintf("&margin_mode=%s", marginMode)
	}
	if option != "" {
		url += option
	}
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsecopytrading.CopytradingTraderQueryContractLeverResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

func (oc *CopytradingClient) CopytradingTraderUpdatedContractLever(data chan responsecopytrading.CopytradingTraderUpdatedContractLeverResponse, request requestcopytrading.CopytradingTraderUpdatedContractLeverRequest) {
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/copytrading/trader/updated_contract_lever", nil)
	content, err := json.Marshal(request)
	if err != nil {
		log.Error("error: %v", err)
	}
	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsecopytrading.CopytradingTraderUpdatedContractLeverResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", getErr)
	}
	data <- result
}

func (oc *CopytradingClient) CopytradingTraderCloseOrder(data chan responsecopytrading.CopytradingTraderCloseOrderResponse, request requestcopytrading.CopytradingTraderCloseOrderRequest) {
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/copytrading/trader/close_order", nil)
	content, err := json.Marshal(request)
	if err != nil {
		log.Error("error: %v", err)
	}
	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsecopytrading.CopytradingTraderCloseOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", getErr)
	}
	data <- result
}

func (oc *CopytradingClient) CopytradingTraderCloseAllPosition(data chan responsecopytrading.CopytradingTraderCloseAllPositionResponse, request requestcopytrading.CopytradingTraderCloseAllPositionRequest) {
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/copytrading/trader/close_all_position", nil)
	content, err := json.Marshal(request)
	if err != nil {
		log.Error("error: %v", err)
	}
	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsecopytrading.CopytradingTraderCloseAllPositionResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", getErr)
	}
	data <- result
}

func (oc *CopytradingClient) CopytradingTraderAddMargin(data chan responsecopytrading.CopytradingTraderAddMarginResponse, request requestcopytrading.CopytradingTraderAddMarginRequest) {
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/copytrading/trader/add_margin", nil)
	content, err := json.Marshal(request)
	if err != nil {
		log.Error("error: %v", err)
	}
	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsecopytrading.CopytradingTraderAddMarginResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", getErr)
	}
	data <- result
}

func (ac *CopytradingClient) CopytradingTraderPositionList(data chan responsecopytrading.CopytradingTraderPositionListResponse) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.GET_METHOD, "/copytrading/trader/position_list", nil)
	// option
	option := "?"
	if option != "" {
		url += option
	}
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsecopytrading.CopytradingTraderPositionListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

func (oc *CopytradingClient) CopytradingTraderAccountTransfer(data chan responsecopytrading.CopytradingTraderAccountTransferResponse, request requestcopytrading.CopytradingTraderAccountTransferRequest) {
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/copytrading/trader/account_transfer", nil)
	content, err := json.Marshal(request)
	if err != nil {
		log.Error("error: %v", err)
	}
	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsecopytrading.CopytradingTraderAccountTransferResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", getErr)
	}
	data <- result
}

func (ac *CopytradingClient) CopytradingTraderCurrentPositions(data chan responsecopytrading.CopytradingTraderCurrentPositionsResponse,
	contractCode string, startTime int64, endTime int64, direct string, fromId int64, limit int64) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.GET_METHOD, "/copytrading/trader/current_positions", nil)
	// option
	option := "?"
	if contractCode != "" {
		option += fmt.Sprintf("&contract_code=%s", contractCode)
	}
	option += fmt.Sprintf("&startTime=%s", startTime)
	option += fmt.Sprintf("&end_time=%s", endTime)
	option += fmt.Sprintf("&direct=%s", direct)
	option += fmt.Sprintf("&from_id=%s", fromId)
	option += fmt.Sprintf("&limit=%s", limit)
	if option != "" {
		url += option
	}
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsecopytrading.CopytradingTraderCurrentPositionsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

func (ac *CopytradingClient) CopytradingTraderPositionPositions(data chan responsecopytrading.CopytradingTraderPositionPositionsResponse,
	contractCode string, startTime int64, endTime int64, direct string, fromId int64, limit int64) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.GET_METHOD, "/copytrading/trader/position_positions", nil)
	// option
	option := "?"
	if contractCode != "" {
		option += fmt.Sprintf("&contract_code=%s", contractCode)
	}
	option += fmt.Sprintf("&startTime=%s", startTime)
	option += fmt.Sprintf("&end_time=%s", endTime)
	option += fmt.Sprintf("&direct=%s", direct)
	option += fmt.Sprintf("&from_id=%s", fromId)
	option += fmt.Sprintf("&limit=%s", limit)
	if option != "" {
		url += option
	}
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsecopytrading.CopytradingTraderPositionPositionsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

func (oc *CopytradingClient) CopytradingTraderTpslOrder(data chan responsecopytrading.CopytradingTraderTpslOrderResponse, request requestcopytrading.CopytradingTraderTpslOrderRequest) {
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/copytrading/trader/tpsl_order", nil)
	content, err := json.Marshal(request)
	if err != nil {
		log.Error("error: %v", err)
	}
	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsecopytrading.CopytradingTraderTpslOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", getErr)
	}
	data <- result
}

func (ac *CopytradingClient) CopytradingTraderQueryContract(data chan responsecopytrading.CopytradingTraderQueryContractResponse) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.GET_METHOD, "/copytrading/trader/query_contract", nil)
	// option
	option := "?"
	if option != "" {
		url += option
	}
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsecopytrading.CopytradingTraderQueryContractResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

func (ac *CopytradingClient) CopytradingTraderOrderTotalDetail(data chan responsecopytrading.CopytradingTraderOrderTotalDetailResponse) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.GET_METHOD, "/copytrading/trader/order_total_detail", nil)
	// option
	option := "?"
	if option != "" {
		url += option
	}
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsecopytrading.CopytradingTraderOrderTotalDetailResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

func (ac *CopytradingClient) CopytradingTraderProfitHistoryDetails(data chan responsecopytrading.CopytradingTraderProfitHistoryDetailsResponse,
	startTime int64, endTime int64, direct string, fromId int64, limit int64) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.GET_METHOD, "/copytrading/trader/profit_history_details", nil)
	// option
	option := "?"
	option += fmt.Sprintf("&startTime=%s", startTime)
	option += fmt.Sprintf("&end_time=%s", endTime)
	option += fmt.Sprintf("&direct=%s", direct)
	option += fmt.Sprintf("&from_id=%s", fromId)
	option += fmt.Sprintf("&limit=%s", limit)
	if option != "" {
		url += option
	}
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsecopytrading.CopytradingTraderProfitHistoryDetailsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

func (ac *CopytradingClient) CopytradingTraderTotalProfitHistory(data chan responsecopytrading.CopytradingTraderTotalProfitHistoryResponse) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.GET_METHOD, "/copytrading/trader/total_profit_history", nil)
	// option
	option := "?"
	if option != "" {
		url += option
	}
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsecopytrading.CopytradingTraderTotalProfitHistoryResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

func (ac *CopytradingClient) CopytradingTraderQueryFollowers(data chan responsecopytrading.CopytradingTraderQueryFollowersResponse,
	startTime int64, endTime int64, direct string, fromId int64, limit int64) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.GET_METHOD, "/copytrading/trader/query_followers", nil)
	// option
	option := "?"
	option += fmt.Sprintf("&startTime=%s", startTime)
	option += fmt.Sprintf("&end_time=%s", endTime)
	option += fmt.Sprintf("&direct=%s", direct)
	option += fmt.Sprintf("&from_id=%s", fromId)
	option += fmt.Sprintf("&limit=%s", limit)
	if option != "" {
		url += option
	}
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsecopytrading.CopytradingTraderQueryFollowersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

func (oc *CopytradingClient) CopytradingTraderRemoveFollower(data chan responsecopytrading.CopytradingTraderRemoveFollowerResponse, request requestcopytrading.CopytradingTraderRemoveFollowerRequest) {
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/copytrading/trader/remove_follower", nil)
	content, err := json.Marshal(request)
	if err != nil {
		log.Error("error: %v", err)
	}
	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsecopytrading.CopytradingTraderRemoveFollowerResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", getErr)
	}
	data <- result
}

func (ac *CopytradingClient) CopytradingTraderOpenOrders(data chan responsecopytrading.CopytradingTraderOpenOrdersResponse,
	contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.GET_METHOD, "/copytrading/trader/open_orders", nil)
	// option
	option := "?"
	option += fmt.Sprintf("&contract_code=%s", contractCode)
	if option != "" {
		url += option
	}
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsecopytrading.CopytradingTraderOpenOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

func (ac *CopytradingClient) CopytradingTraderTpslOpenOrders(data chan responsecopytrading.CopytradingTraderTpslOpenOrdersResponse,
	contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.GET_METHOD, "/copytrading/trader/tpsl_open_orders", nil)
	// option
	option := "?"
	option += fmt.Sprintf("&contract_code=%s", contractCode)
	if option != "" {
		url += option
	}
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsecopytrading.CopytradingTraderTpslOpenOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}
