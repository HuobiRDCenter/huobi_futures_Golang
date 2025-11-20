package restful

import (
	"encoding/json"
	"fmt"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinswap"
	requestorder "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinswap/restful/request/order"
	responseorder "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinswap/restful/response/order"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/reqbuilder"
)

type OrderClient struct {
	PUrlBuilder *reqbuilder.PrivateUrlBuilder
}

func (oc *OrderClient) Init(accessKey string, secretKey string, host string, sign string) *OrderClient {
	if host == "" {
		host = coinswap.COIN_SWAP_DEFAULT_HOST
	}
	oc.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host, sign)
	return oc
}

// SwapCancelAfterAsync (Automatic Order Cancellation)
func (oc *OrderClient) SwapCancelAfterAsync(data chan responseorder.SwapCancelAfterResponse, onOff string,
	timeOut string) {
	url := oc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap-cancel-after", nil)

	// content
	content := ""
	if onOff != "" {
		content = fmt.Sprintf(",\"on_off\": \"%s\"", onOff)
	}
	if timeOut != "" {
		content = fmt.Sprintf(",\"time_out\": \"%s\"", timeOut)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.SwapCancelAfterResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapCancelAfterResponse error: %s", getErr)
	}
	data <- result
}

// SwapOrderAsync (Place an Order)
func (oc *OrderClient) SwapOrderAsync(data chan responseorder.SwapOrderResponse, request requestorder.SwapOrderRequest) {
	url := oc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_order", nil)

	content, err := json.Marshal(request)
	if err != nil {
		log.Error("PlaceOrderRequest to json error: %v", err)
	}
	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.SwapOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapOrderResponse error: %s", getErr)
	}
	data <- result
}

// SwapBatchOrderAsync (Place a Batch of Orders)
func (oc *OrderClient) SwapBatchOrderAsync(data chan responseorder.SwapBatchOrderResponse, request requestorder.BatchSwapOrderRequest) {
	url := oc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_batchorder", nil)

	content, err := json.Marshal(request)
	if err != nil {
		log.Error("PlaceOrderRequest to json error: %v", err)
	}
	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.SwapBatchOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapBatchOrderResponse error: %s", getErr)
	}
	data <- result
}

// SwapCancelAsync (Cancel an Order)
func (oc *OrderClient) SwapCancelAsync(data chan responseorder.SwapCancelResponse, orderId string,
	clientOrderId string, contractCode string) {
	url := oc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_cancel", nil)

	// content
	content := ""
	if orderId != "" {
		content = fmt.Sprintf(",\"order_id\": \"%s\"", orderId)
	}
	if clientOrderId != "" {
		content = fmt.Sprintf(",\"client_order_id\": \"%s\"", clientOrderId)
	}
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
	result := responseorder.SwapCancelResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapCancelResponse error: %s", getErr)
	}
	data <- result
}

// SwapCancelAllAsync (Cancel All Orders)
func (oc *OrderClient) SwapCancelAllAsync(data chan responseorder.SwapCancelAllResponse, contractCode string,
	direction string, offset string) {
	url := oc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_cancelall", nil)

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if direction != "" {
		content = fmt.Sprintf(",\"direction\": \"%s\"", direction)
	}
	if offset != "" {
		content = fmt.Sprintf(",\"offset\": \"%s\"", offset)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.SwapCancelAllResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapCancelAllResponse error: %s", getErr)
	}
	data <- result
}

// SwapSwitchLeverRateAsync (Switch Leverage)
func (oc *OrderClient) SwapSwitchLeverRateAsync(data chan responseorder.SwapSwitchLeverRateResponse, contractCode string,
	leverRate string) {
	url := oc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_switch_lever_rate", nil)

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if leverRate != "" {
		content = fmt.Sprintf(",\"lever_rate\": \"%s\"", leverRate)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.SwapSwitchLeverRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapSwitchLeverRateResponse error: %s", getErr)
	}
	data <- result
}

// SwapOrderInfoAsync (Get Information of an Order)
func (oc *OrderClient) SwapOrderInfoAsync(data chan responseorder.SwapOrderInfoResponse, orderId string,
	clientOrderId string, contractCode string) {
	url := oc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_order_info", nil)

	// content
	content := ""
	if orderId != "" {
		content = fmt.Sprintf(",\"order_id\": \"%s\"", orderId)
	}
	if clientOrderId != "" {
		content = fmt.Sprintf(",\"client_order_id\": \"%s\"", clientOrderId)
	}
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
	result := responseorder.SwapOrderInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapOrderInfoResponse error: %s", getErr)
	}
	data <- result
}

// SwapOrderDetailAsync (Order details acquisition)
func (oc *OrderClient) SwapOrderDetailAsync(data chan responseorder.SwapOrderDetailResponse, contractCode string,
	orderId string, createdAt string, orderType string, pageIndex string, pageSize string) {
	url := oc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_order_detail", nil)

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if orderId != "" {
		content = fmt.Sprintf(",\"order_id\": \"%s\"", orderId)
	}
	if createdAt != "" {
		content = fmt.Sprintf(",\"created_at\": \"%s\"", createdAt)
	}
	if orderType != "" {
		content = fmt.Sprintf(",\"order_type\": \"%s\"", orderType)
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
	result := responseorder.SwapOrderDetailResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapOrderDetailResponse error: %s", getErr)
	}
	data <- result
}

// SwapOpenOrdersAsync (Current unfilled order acquisition)
func (oc *OrderClient) SwapOpenOrdersAsync(data chan responseorder.SwapOpenOrdersResponse, contractCode string,
	pageIndex string, pageSize string, sortBy string, tradeType string) {
	url := oc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_openorders", nil)

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
	if sortBy != "" {
		content = fmt.Sprintf(",\"sort_by\": \"%s\"", sortBy)
	}
	if tradeType != "" {
		content = fmt.Sprintf(",\"trade_type\": \"%s\"", tradeType)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.SwapOpenOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapOpenOrdersResponse error: %s", getErr)
	}
	data <- result
}

// SwapHisOrdersAsync (Get History Orders(New))
func (oc *OrderClient) SwapHisOrdersAsync(data chan responseorder.SwapHisOrdersResponse, startTime string,
	endTime string, direct string, fromId string, contract string, tradeType string, fcType string, status string) {
	url := oc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v3/swap_hisorders", nil)

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
	if contract != "" {
		content = fmt.Sprintf(",\"contract\": \"%s\"", contract)
	}
	if tradeType != "" {
		content = fmt.Sprintf(",\"trade_type\": \"%s\"", tradeType)
	}
	if fcType != "" {
		content = fmt.Sprintf(",\"type\": \"%s\"", fcType)
	}
	if status != "" {
		content = fmt.Sprintf(",\"status\": \"%s\"", status)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.SwapHisOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapHisOrdersResponse error: %s", getErr)
	}
	data <- result
}

// SwapHisOrdersExactAsync (Query history orders via multiple fields(New))
func (oc *OrderClient) SwapHisOrdersExactAsync(data chan responseorder.SwapHisOrdersExactResponse, startTime string,
	endTime string, direct string, fromId string, contract string, tradeType string, fcType string, status string, priceType string) {
	url := oc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v3/swap_hisorders_exact", nil)

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
	if contract != "" {
		content = fmt.Sprintf(",\"contract\": \"%s\"", contract)
	}
	if tradeType != "" {
		content = fmt.Sprintf(",\"trade_type\": \"%s\"", tradeType)
	}
	if fcType != "" {
		content = fmt.Sprintf(",\"type\": \"%s\"", fcType)
	}
	if status != "" {
		content = fmt.Sprintf(",\"status\": \"%s\"", status)
	}
	if priceType != "" {
		content = fmt.Sprintf(",\"price_type\": \"%s\"", priceType)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.SwapHisOrdersExactResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapHisOrdersExactResponse error: %s", getErr)
	}
	data <- result
}

// SwapMatchResultsAsync (Acquire History Match Results(New))
func (oc *OrderClient) SwapMatchResultsAsync(data chan responseorder.SwapMatchResultsResponse, contract string, tradeType string, startTime string,
	endTime string, direct string, fromId string) {
	url := oc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v3/swap_matchresults", nil)

	// content
	content := ""
	if contract != "" {
		content = fmt.Sprintf(",\"contract\": \"%s\"", contract)
	}
	if tradeType != "" {
		content = fmt.Sprintf(",\"trade_type\": \"%s\"", tradeType)
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
	result := responseorder.SwapMatchResultsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapMatchResultsResponse error: %s", getErr)
	}
	data <- result
}

// SwapMatchResultsExactAsync (Query history transactions via multiple fields(New))
func (oc *OrderClient) SwapMatchResultsExactAsync(data chan responseorder.SwapMatchResultsExactResponse, contract string, tradeType string, startTime string,
	endTime string, direct string, fromId string) {
	url := oc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v3/swap_matchresults_exact", nil)

	// content
	content := ""
	if contract != "" {
		content = fmt.Sprintf(",\"contract\": \"%s\"", contract)
	}
	if tradeType != "" {
		content = fmt.Sprintf(",\"trade_type\": \"%s\"", tradeType)
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
	result := responseorder.SwapMatchResultsExactResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapMatchResultsExactResponse error: %s", getErr)
	}
	data <- result
}

// SwapLightningClosePositionAsync (Place Lightning Close Order)
func (oc *OrderClient) SwapLightningClosePositionAsync(data chan responseorder.SwapLightningClosePositionResponse, contractCode string,
	volume string, direction string, clientOrderId string, orderPriceType string) {
	url := oc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_lightning_close_position", nil)

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if volume != "" {
		content = fmt.Sprintf(",\"volume\": \"%s\"", volume)
	}
	if direction != "" {
		content = fmt.Sprintf(",\"direction\": \"%s\"", direction)
	}
	if clientOrderId != "" {
		content = fmt.Sprintf(",\"client_order_id\": \"%s\"", clientOrderId)
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
	result := responseorder.SwapLightningClosePositionResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapLightningClosePositionResponse error: %s", getErr)
	}
	data <- result
}
