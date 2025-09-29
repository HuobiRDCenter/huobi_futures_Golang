package restful

import (
	"encoding/json"
	"fmt"

	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap"
	requestorder "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/request/order"
	responseorder "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/response/order"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/reqbuilder"
)

type OrderClient struct {
	PUrlBuilder *reqbuilder.PrivateUrlBuilder
}

func (oc *OrderClient) Init(accessKey string, secretKey string, host string) *OrderClient {
	if host == "" {
		host = linearswap.LINEAR_SWAP_DEFAULT_HOST
	}
	oc.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return oc
}

// IsolatedPlaceOrderAsync ([Isolated] Place an Order)
func (oc *OrderClient) IsolatedPlaceOrderAsync(data chan responseorder.PlaceOrderResponse, request requestorder.PlaceOrderRequest) {
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_order", nil)

	content, err := json.Marshal(request)
	if err != nil {
		log.Error("PlaceOrderRequest to json error: %v", err)
	}
	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.PlaceOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to PlaceOrderResponse error: %s", getErr)
	}
	data <- result
}

// CrossPlaceOrderAsync ([Cross] Place An Order)
func (oc *OrderClient) CrossPlaceOrderAsync(data chan responseorder.PlaceOrderResponse, request requestorder.PlaceOrderRequest) {
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_order", nil)

	content, err := json.Marshal(request)
	if err != nil {
		log.Error("PlaceOrderRequest to json error: %v", err)
	}
	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.PlaceOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to PlaceOrderResponse error: %s", getErr)
	}
	data <- result
}

// IsolatedPlaceBatchOrderAsync ([Isolated] Place a Batch of Orders)
func (oc *OrderClient) IsolatedPlaceBatchOrderAsync(data chan responseorder.PlaceBatchOrderResponse, requests requestorder.BatchPlaceOrderRequest) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_batchorder", nil)

	// content
	bdata, err := json.Marshal(requests)
	if err != nil {
		log.Error("[] PlaceOrderRequest to json error: %v", err)
	}
	content := string(bdata)
	content = fmt.Sprintf("{\"orders_data\":%s}", content)

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.PlaceBatchOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to PlaceBatchOrderResponse error: %s", getErr)
	}
	data <- result
}

// CrossPlaceBatchOrderAsync ([Cross] Place A Batch Of Orders)
func (oc *OrderClient) CrossPlaceBatchOrderAsync(data chan responseorder.PlaceBatchOrderResponse, requests requestorder.BatchPlaceOrderRequest) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_batchorder", nil)

	// content
	bdata, err := json.Marshal(requests)
	if err != nil {
		log.Error("[] PlaceOrderRequest to json error: %v", err)
	}
	content := string(bdata)
	content = fmt.Sprintf("{\"orders_data\":%s}", content)

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.PlaceBatchOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to PlaceBatchOrderResponse error: %s", getErr)
	}
	data <- result
}

// IsolatedCancelOrderAsync ([Isolated] Cancel an Order) and ([Isolated] Cancel All Orders)
func (oc *OrderClient) IsolatedCancelOrderAsync(data chan responseorder.CancelOrderResponse, contractCode string,
	orderId string, clientOrderId string, offset string, direction string) {

	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cancel", nil)
	if orderId == "" && clientOrderId == "" {
		url = oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cancelall", nil)
	}

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if orderId != "" {
		content += fmt.Sprintf(",\"order_id\": \"%s\"", orderId)
	}
	if clientOrderId != "" {
		content += fmt.Sprintf(",\"client_order_id\": \"%s\"", clientOrderId)
	}
	if offset != "" {
		content += fmt.Sprintf(",\"offset\": \"%s\"", offset)
	}
	if direction != "" {
		content += fmt.Sprintf(",\"direction\": \"%s\"", direction)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.CancelOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to CancelOrderResponse error: %s", getErr)
	}
	data <- result
}

// CrossCancelOrderAsync ([Cross] Cancel An Order) and  ([Cross] Cancel All Orders)
func (oc *OrderClient) CrossCancelOrderAsync(data chan responseorder.CancelOrderResponse, contractCode string,
	orderId string, clientOrderId string, offset string, direction string, pair string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_cancel", nil)
	if orderId == "" && clientOrderId == "" {
		url = oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_cancelall", nil)
	}

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if orderId != "" {
		content += fmt.Sprintf(",\"order_id\": \"%s\"", orderId)
	}
	if clientOrderId != "" {
		content += fmt.Sprintf(",\"client_order_id\": \"%s\"", clientOrderId)
	}
	if offset != "" {
		content += fmt.Sprintf(",\"offset\": \"%s\"", offset)
	}
	if direction != "" {
		content += fmt.Sprintf(",\"direction\": \"%s\"", direction)
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
	result := responseorder.CancelOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to CancelOrderResponse error: %s", getErr)
	}
	data <- result
}

// IsolatedSwitchLeverRateAsync ([Isolated] Switch Leverage)
func (oc *OrderClient) IsolatedSwitchLeverRateAsync(data chan responseorder.SwitchLeverRateResponse, contractCode string, leverRate int) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_switch_lever_rate", nil)

	// content
	content := fmt.Sprintf("{\"contract_code\": \"%s\", \"lever_rate\": %d}", contractCode, leverRate)

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.SwitchLeverRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwitchLeverRateResponse error: %s", getErr)
	}
	data <- result
}

// CrossSwitchLeverRateAsync ([Cross] Switch Leverage)
func (oc *OrderClient) CrossSwitchLeverRateAsync(data chan responseorder.SwitchLeverRateResponse, contractCode string, leverRate int, pair string, contractType string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_switch_lever_rate", nil)

	// content
	content := fmt.Sprintf("\"contract_code\": \"%s\", \"lever_rate\": %d", contractCode, leverRate)
	if pair != "" {
		content += fmt.Sprintf(",\"pair\": %s", pair)
	}
	if contractType != "" {
		content += fmt.Sprintf(",\"contract_type\": %s", contractType)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content)
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.SwitchLeverRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwitchLeverRateResponse error: %s", getErr)
	}
	data <- result
}

// IsolatedGetOrderInfoAsync ([Isolated] Get Information of an Order)
func (oc *OrderClient) IsolatedGetOrderInfoAsync(data chan responseorder.GetOrderInfoResponse, contractCode string,
	orderId string, clientOrderId string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_order_info", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\":\"%s\"", contractCode)
	if orderId != "" {
		content += fmt.Sprintf(",\"order_id\":\"%s\"", orderId)
	}
	if clientOrderId != "" {
		content += fmt.Sprintf(",\"client_order_id\":\"%s\"", clientOrderId)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.GetOrderInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetOrderInfoResponse error: %s", getErr)
	}
	data <- result
}

// CrossGetOrderInfoAsync ([Cross] Get Information of order)
func (oc *OrderClient) CrossGetOrderInfoAsync(data chan responseorder.GetOrderInfoResponse, contractCode string, orderId string, clientOrderId string, pair string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_order_info", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\":\"%s\"", contractCode)
	if orderId != "" {
		content += fmt.Sprintf(",\"order_id\":\"%s\"", orderId)
	}
	if clientOrderId != "" {
		content += fmt.Sprintf(",\"client_order_id\":\"%s\"", clientOrderId)
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
	result := responseorder.GetOrderInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetOrderInfoResponse error: %s", getErr)
	}
	data <- result
}

// IsolatedGetOrderDetailAsync ([Isolated] Order details acquisition)
func (oc *OrderClient) IsolatedGetOrderDetailAsync(data chan responseorder.GetOrderDetailResponse, contractCode string, orderId int64, createdAt int64,
	orderType int, pageIndex int, pageSize int) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_order_detail", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\", \"order_id\": %d", contractCode, orderId)
	if createdAt != 0 {
		content += fmt.Sprintf(",\"created_at\": %d", createdAt)
	}
	if orderType != 0 {
		content += fmt.Sprintf(",\"order_type\": %d", orderType)
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
	result := responseorder.GetOrderDetailResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetOrderDetailResponse error: %s", getErr)
	}
	data <- result
}

// CrossGetOrderDetailAsync ([Cross] Get Detail Information of order)
func (oc *OrderClient) CrossGetOrderDetailAsync(data chan responseorder.GetOrderDetailResponse, contractCode string, orderId int64, createdAt int64,
	orderType int, pageIndex int, pageSize int, pair string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_order_detail", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\", \"order_id\": %d", contractCode, orderId)
	if createdAt != 0 {
		content += fmt.Sprintf(",\"created_at\": %d", createdAt)
	}
	if orderType != 0 {
		content += fmt.Sprintf(",\"order_type\": %d", orderType)
	}
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
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
	result := responseorder.GetOrderDetailResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetOrderDetailResponse error: %s", getErr)
	}
	data <- result
}

// IsolatedGetOpenOrderAsync ([Isolated] Current unfilled order acquisition)
func (oc *OrderClient) IsolatedGetOpenOrderAsync(data chan responseorder.GetOpenOrderResponse, contractCode string,
	pageIndex int, pageSize int, sortBy string, tradeType int) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_openorders", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if sortBy != "" {
		content += fmt.Sprintf(",\"sort_by\": \"%s\"", sortBy)
	}
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"trade_type\": %d", tradeType)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.GetOpenOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetOpenOrderResponse error: %s", getErr)
	}
	data <- result
}

// CrossGetOpenOrderAsync ([Cross] Current unfilled order acquisition)
func (oc *OrderClient) CrossGetOpenOrderAsync(data chan responseorder.GetOpenOrderResponse, contractCode string,
	pageIndex int, pageSize int, sortBy string, tradeType int) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_openorders", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if sortBy != "" {
		content += fmt.Sprintf(",\"sort_by\": \"%s\"", sortBy)
	}
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"trade_type\": %d", tradeType)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.GetOpenOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetOpenOrderResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) IsolatedGetHisOrderAsync(data chan responseorder.GetHisOrderResponse, contractCode string, tradeType int, fcType int, status string,
	createDate int, pageIndex int, pageSize int, sortBy string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_hisorders", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\",\"trade_type\": \"%d\",\"type\": \"%d\",\"status\": \"%s\",\"create_date\": %d", contractCode, tradeType, fcType, status, createDate)
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if sortBy != "" {
		content += fmt.Sprintf(",\"sort_by\": \"%s\"", sortBy)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.GetHisOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHisOrderResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) CrossGetHisOrderAsync(data chan responseorder.GetHisOrderResponse, contractCode string, tradeType int, fcType int, status string,
	createDate int, pageIndex int, pageSize int, sortBy string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_hisorders", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\",\"trade_type\": \"%d\",\"type\": \"%d\",\"status\": \"%s\",\"create_date\": %d", contractCode, tradeType, fcType, status, createDate)
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if sortBy != "" {
		content += fmt.Sprintf(",\"sort_by\": \"%s\"", sortBy)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.GetHisOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHisOrderResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) IsolatedGetHisOrderExactAsync(data chan responseorder.GetHisOrderExactResponse, contractCode string,
	tradeType int, fcType int, status string, orderPriceType string, startTime int64, endTime int64,
	fromId int64, size int, direct string) {

	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_hisorders_exact", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\",\"trade_type\": \"%d\",\"type\": \"%d\",\"status\": \"%s\"", contractCode, tradeType, fcType, status)
	if orderPriceType != "" {
		content += fmt.Sprintf(",\"order_price_type\": \"%s\"", orderPriceType)
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
	result := responseorder.GetHisOrderExactResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHisOrderExactResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) CrossGetHisOrderExactAsync(data chan responseorder.GetHisOrderExactResponse, contractCode string,
	tradeType int, fcType int, status string, orderPriceType string, startTime int64, endTime int64,
	fromId int64, size int, direct string) {

	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_hisorders_exact", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\",\"trade_type\": \"%d\",\"type\": \"%d\",\"status\": \"%s\"", contractCode, tradeType, fcType, status)
	if orderPriceType != "" {
		content += fmt.Sprintf(",\"order_price_type\": \"%s\"", orderPriceType)
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
	result := responseorder.GetHisOrderExactResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHisOrderExactResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) IsolatedGetHisMatchAsync(data chan responseorder.GetHisMatchResponse, contractCode string, tradeType int, createDate int,
	pageIndex int, pageSize int) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_matchresults", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\",\"trade_type\": %d,\"create_date\": %d", contractCode, tradeType, createDate)
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
	result := responseorder.GetHisMatchResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHisMatchResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) CrossGetHisMatchAsync(data chan responseorder.GetHisMatchResponse, contractCode string, tradeType int, createDate int,
	pageIndex int, pageSize int) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_matchresults", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\",\"trade_type\": %d,\"create_date\": %d", contractCode, tradeType, createDate)
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
	result := responseorder.GetHisMatchResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHisMatchResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) IsolatedGetHisMatchExactAsync(data chan responseorder.GetHisMatchExactResponse, contractCode string,
	tradeType int, startTime int64, endTime int64, fromId int64, size int, direct string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_matchresults_exact", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\",\"trade_type\": %d", contractCode, tradeType)
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
	result := responseorder.GetHisMatchExactResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHisMatchExactResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) CrossGetHisMatchExactAsync(data chan responseorder.GetHisMatchExactResponse, contractCode string,
	tradeType int, startTime int64, endTime int64, fromId int64, size int, direct string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_matchresults_exact", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\",\"trade_type\": %d", contractCode, tradeType)
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
	result := responseorder.GetHisMatchExactResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHisMatchExactResponse error: %s", getErr)
	}
	data <- result
}

// IsolatedLightningCloseAsync ([Isolated] Place Lightning Close Order)
func (oc *OrderClient) IsolatedLightningCloseAsync(data chan responseorder.LightningCloseResponse, contractCode string, volume int, direction string,
	clientOrderId int64, orderPriceType string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_lightning_close_position", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\",\"volume\": %d,\"direction\": \"%s\"", contractCode, volume, direction)
	if clientOrderId != 0 {
		content += fmt.Sprintf(",\"client_order_id\": %d", clientOrderId)
	}
	if orderPriceType != "" {
		content += fmt.Sprintf(",\"order_price_type\": \"%s\"", orderPriceType)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.LightningCloseResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to LightningCloseResponse error: %s", getErr)
	}
	data <- result
}

// CrossLightningCloseAsync ([Cross] Place Lightning Close Position)
func (oc *OrderClient) CrossLightningCloseAsync(data chan responseorder.LightningCloseResponse, contractCode string, volume int, direction string,
	clientOrderId int64, orderPriceType string, pair string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_lightning_close_position", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\",\"volume\": %d,\"direction\": \"%s\"", contractCode, volume, direction)
	if clientOrderId != 0 {
		content += fmt.Sprintf(",\"client_order_id\": %d", clientOrderId)
	}
	if orderPriceType != "" {
		content += fmt.Sprintf(",\"order_price_type\": \"%s\"", orderPriceType)
	}
	if pair != "" {
		content += fmt.Sprintf(",\"pair\": \"%s\"", pair)
	}

	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.LightningCloseResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to LightningCloseResponse error: %s", getErr)
	}
	data <- result
}

// LinearCancelAfterAsync ([General] Automatic Order Cancellation)
func (oc *OrderClient) LinearCancelAfterAsync(data chan responseorder.LinearCancelAfterResponse, onOff string, timeOut string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/linear-cancel-after", nil)

	// content
	content := ""
	if onOff != "" {
		content += fmt.Sprintf(",\"on_off\": \"%s\"", onOff)
	}
	if timeOut != "" {
		content += fmt.Sprintf(",\"time_out\": \"%s\"", timeOut)
	}

	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.LinearCancelAfterResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to LinearCancelAfterResponse error: %s", getErr)
	}
	data <- result
}

// SwapSwitchPositionModeAsync ([Isolated]Switch Position Mode)
func (oc *OrderClient) SwapSwitchPositionModeAsync(data chan responseorder.SwapSwitchPositionModeResponse, marginAccount string, positionMode string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_switch_position_mode", nil)

	// content
	content := ""
	if marginAccount != "" {
		content += fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	}
	if positionMode != "" {
		content += fmt.Sprintf(",\"position_mode\": \"%s\"", positionMode)
	}

	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.SwapSwitchPositionModeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapSwitchPositionModeResponse error: %s", getErr)
	}
	data <- result
}

// SwapCrossSwitchPositionModeAsync ([Cross]Switch Position Mode)
func (oc *OrderClient) SwapCrossSwitchPositionModeAsync(data chan responseorder.SwapSwitchPositionModeResponse, marginAccount string, positionMode string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_switch_position_mode", nil)

	// content
	content := ""
	if marginAccount != "" {
		content += fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	}
	if positionMode != "" {
		content += fmt.Sprintf(",\"position_mode\": \"%s\"", positionMode)
	}

	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.SwapSwitchPositionModeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapSwitchPositionModeResponse error: %s", getErr)
	}
	data <- result
}

// SwapHisordersAsync ([Isolated] Get History Orders(New))
func (oc *OrderClient) SwapHisordersAsync(data chan responseorder.SwapHisordersResponse, contract string, tradeType string, startTime string,
	endTime string, direct string, fromID string, orderType string, status string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v3/swap_hisorders", nil)

	// content
	content := ""
	if contract != "" {
		content += fmt.Sprintf(",\"contract\": \"%s\"", contract)
	}
	if tradeType != "" {
		content += fmt.Sprintf(",\"trade_type\": \"%s\"", tradeType)
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
	if fromID != "" {
		content += fmt.Sprintf(",\"from_id\": \"%s\"", fromID)
	}
	if orderType != "" {
		content += fmt.Sprintf(",\"type\": \"%s\"", orderType)
	}
	if status != "" {
		content += fmt.Sprintf(",\"status\": \"%s\"", status)
	}

	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.SwapHisordersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapHisordersResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) SwapTradeOrderAsync(data chan responseorder.SwapTradeOrderResponse, contractCode string, marginMode string, positionSide string,
	side string, type_ string, priceMatch string, clientOrderId string, price string, volume string, reduceOnly int, timeInForce string,
	tpTriggerPrice string, tpOrderPrice string, tpType string, tpTriggerPriceType string, slTriggerPrice string, slOrderPrice string,
	slType string, slTriggerPriceType string, priceProtect bool, triggerProtect bool) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/v5/trade/order", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if marginMode != "" {
		content += fmt.Sprintf(",\"margin_mode\": \"%s\"", marginMode)
	}
	if positionSide != "" {
		content += fmt.Sprintf(",\"position_side\": \"%s\"", positionSide)
	}
	if side != "" {
		content += fmt.Sprintf(",\"side\": \"%s\"", side)
	}
	if type_ != "" {
		content += fmt.Sprintf(",\"type\": \"%s\"", type_)
	}
	if priceMatch != "" {
		content += fmt.Sprintf(",\"price_match\": \"%s\"", priceMatch)
	}
	if clientOrderId != "" {
		content += fmt.Sprintf(",\"client_order_id\": \"%s\"", clientOrderId)
	}
	if price != "" {
		content += fmt.Sprintf(",\"price\": \"%s\"", price)
	}
	if volume != "" {
		content += fmt.Sprintf(",\"volume\": \"%s\"", volume)
	}
	content += fmt.Sprintf(",\"reduce_only\": \"%s\"", reduceOnly)
	if timeInForce != "" {
		content += fmt.Sprintf(",\"time_in_force\": \"%s\"", timeInForce)
	}
	if tpTriggerPrice != "" {
		content += fmt.Sprintf(",\"tp_trigger_price\": \"%s\"", tpTriggerPrice)
	}
	if tpOrderPrice != "" {
		content += fmt.Sprintf(",\"tp_order_price\": \"%s\"", tpOrderPrice)
	}
	if tpType != "" {
		content += fmt.Sprintf(",\"tp_type\": \"%s\"", tpType)
	}
	if tpTriggerPriceType != "" {
		content += fmt.Sprintf(",\"tp_trigger_price_type\": \"%s\"", tpTriggerPriceType)
	}
	if slTriggerPrice != "" {
		content += fmt.Sprintf(",\"sl_trigger_price\": \"%s\"", slTriggerPrice)
	}
	if slOrderPrice != "" {
		content += fmt.Sprintf(",\"sl_order_price\": \"%s\"", slOrderPrice)
	}
	if slType != "" {
		content += fmt.Sprintf(",\"sl_type\": \"%s\"", slType)
	}
	if slTriggerPriceType != "" {
		content += fmt.Sprintf(",\"sl_trigger_price_type\": \"%s\"", slTriggerPriceType)
	}
	content += fmt.Sprintf(",\"price_protect\": \"%s\"", priceProtect)
	content += fmt.Sprintf(",\"trigger_protect\": \"%s\"", triggerProtect)
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.SwapTradeOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) SwapTradeBatchOrderAsync(data chan responseorder.SwapTradeBatchOrderResponse, contractCode string, marginMode string,
	side string, type_ string, priceMatch string, clientOrderId string, price string, volume string, reduceOnly int, timeInForce string,
	tpTriggerPrice string, tpOrderPrice string, tpType string, tpTriggerPriceType string, slTriggerPrice string, slOrderPrice string,
	slType string, slTriggerPriceType string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/v5/trade/batch_orders", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if marginMode != "" {
		content += fmt.Sprintf(",\"margin_mode\": \"%s\"", marginMode)
	}
	if side != "" {
		content += fmt.Sprintf(",\"side\": \"%s\"", side)
	}
	if type_ != "" {
		content += fmt.Sprintf(",\"type\": \"%s\"", type_)
	}
	if priceMatch != "" {
		content += fmt.Sprintf(",\"price_match\": \"%s\"", priceMatch)
	}
	if clientOrderId != "" {
		content += fmt.Sprintf(",\"client_order_id\": \"%s\"", clientOrderId)
	}
	if price != "" {
		content += fmt.Sprintf(",\"price\": \"%s\"", price)
	}
	if volume != "" {
		content += fmt.Sprintf(",\"volume\": \"%s\"", volume)
	}
	content += fmt.Sprintf(",\"reduce_only\": \"%s\"", reduceOnly)
	if timeInForce != "" {
		content += fmt.Sprintf(",\"time_in_force\": \"%s\"", timeInForce)
	}
	if tpTriggerPrice != "" {
		content += fmt.Sprintf(",\"tp_trigger_price\": \"%s\"", tpTriggerPrice)
	}
	if tpOrderPrice != "" {
		content += fmt.Sprintf(",\"tp_order_price\": \"%s\"", tpOrderPrice)
	}
	if tpType != "" {
		content += fmt.Sprintf(",\"tp_type\": \"%s\"", tpType)
	}
	if tpTriggerPriceType != "" {
		content += fmt.Sprintf(",\"tp_trigger_price_type\": \"%s\"", tpTriggerPriceType)
	}
	if slTriggerPrice != "" {
		content += fmt.Sprintf(",\"sl_trigger_price\": \"%s\"", slTriggerPrice)
	}
	if slOrderPrice != "" {
		content += fmt.Sprintf(",\"sl_order_price\": \"%s\"", slOrderPrice)
	}
	if slType != "" {
		content += fmt.Sprintf(",\"sl_type\": \"%s\"", slType)
	}
	if slTriggerPriceType != "" {
		content += fmt.Sprintf(",\"sl_trigger_price_type\": \"%s\"", slTriggerPriceType)
	}

	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.SwapTradeBatchOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) SwapTradeCancelOrderAsync(data chan responseorder.SwapTradeCancelOrderResponse, contractCode string,
	orderId string, clientOrderId string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/v5/trade/cancel_order", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if orderId != "" {
		content += fmt.Sprintf(",\"order_id\": \"%s\"", orderId)
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
	result := responseorder.SwapTradeCancelOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) SwapTradeCancelBatchOrdersAsync(data chan responseorder.SwapTradeCancelBatchOrdersResponse, contractCode string,
	orderId string, clientOrderId string, priceProtect bool, triggerProtect bool) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/v5/trade/cancel_batch_orders", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if orderId != "" {
		content += fmt.Sprintf(",\"order_id\": \"%s\"", orderId)
	}
	if clientOrderId != "" {
		content += fmt.Sprintf(",\"client_order_id\": \"%s\"", clientOrderId)
	}
	content += fmt.Sprintf(",\"price_protect\": \"%s\"", priceProtect)
	content += fmt.Sprintf(",\"trigger_protect\": \"%s\"", triggerProtect)
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.SwapTradeCancelBatchOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) SwapTradeAllOrdersAsync(data chan responseorder.SwapTradeAllOrdersResponse, contractCode string,
	side string, positionSide string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/v5/trade/cancel_all_orders", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if side != "" {
		content += fmt.Sprintf(",\"side\": \"%s\"", side)
	}
	if positionSide != "" {
		content += fmt.Sprintf(",\"position_side\": \"%s\"", positionSide)
	}

	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.SwapTradeAllOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) SwapTradePositionAsync(data chan responseorder.SwapTradePositionResponse, contractCode string,
	marginMode string, positionSide string, clientOrderId string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/v5/trade/position", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if marginMode != "" {
		content += fmt.Sprintf(",\"margin_mode\": \"%s\"", marginMode)
	}
	if positionSide != "" {
		content += fmt.Sprintf(",\"position_side\": \"%s\"", positionSide)
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
	result := responseorder.SwapTradePositionResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) SwapTradePositionAllAsync(data chan responseorder.SwapTradePositionAllResponse) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/v5/trade/position_all", nil)

	// content
	content := ""

	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.SwapTradePositionAllResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", getErr)
	}
	data <- result
}

func (ac *AccountClient) GetTradeOpensAsync(data chan responseorder.GetTradeOpensResponse, contractCode string,
	marginMode string, orderId string, clientOrderId string, from int, limit int, direct string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.GET_METHOD, "/v5/trade/order/opens", nil)
	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if marginMode != "" {
		option += fmt.Sprintf("&margin_mode=%s", marginMode)
	}
	if orderId != "" {
		option += fmt.Sprintf("&order_id=%s", orderId)
	}
	if clientOrderId != "" {
		option += fmt.Sprintf("&client_order_id=%s", clientOrderId)
	}
	option += fmt.Sprintf("&from=%s", from)
	option += fmt.Sprintf("&limit=%s", limit)
	if direct != "" {
		option += fmt.Sprintf("&direct=%s", direct)
	}
	if option != "" {
		url += option
	}
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.GetTradeOpensResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) GetTradeOrderTradesAsync(data chan responseorder.GetTradeOpensResponse, contractCode string, side string,
	marginMode string, orderId string, clientOrderId string, from int, limit int, direct string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.GET_METHOD, "/api/V5/trade/order/details", nil)
	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if side != "" {
		option += fmt.Sprintf("&side=%s", side)
	}
	if marginMode != "" {
		option += fmt.Sprintf("&margin_mode=%s", marginMode)
	}
	if orderId != "" {
		option += fmt.Sprintf("&order_id=%s", orderId)
	}
	if clientOrderId != "" {
		option += fmt.Sprintf("&client_order_id=%s", clientOrderId)
	}
	option += fmt.Sprintf("&from=%s", from)
	option += fmt.Sprintf("&limit=%s", limit)
	if direct != "" {
		option += fmt.Sprintf("&direct=%s", direct)
	}
	if option != "" {
		url += option
	}
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.GetTradeOpensResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) GetTradeOrderHistoryAsync(data chan responseorder.GetTradeOrderHistoryResponse,
	contractCode string, state string, type_ string,
	priceMatch string, startTime string, endTime string, from int, limit int, direct string, marginMode string, timeInForce string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.GET_METHOD, "/api/v5/trade/order/history", nil)
	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if timeInForce != "" {
		option += fmt.Sprintf("?time_in_force=%s", timeInForce)
	}
	if state != "" {
		option += fmt.Sprintf("&state=%s", state)
	}
	if type_ != "" {
		option += fmt.Sprintf("&type=%s", type_)
	}
	if priceMatch != "" {
		option += fmt.Sprintf("&price_match=%s", priceMatch)
	}
	if startTime != "" {
		option += fmt.Sprintf("&start_time=%s", startTime)
	}
	if endTime != "" {
		option += fmt.Sprintf("&end_time=%s", endTime)
	}
	if marginMode != "" {
		option += fmt.Sprintf("&margin_mode=%s", marginMode)
	}
	option += fmt.Sprintf("&from=%s", from)
	option += fmt.Sprintf("&limit=%s", limit)
	if direct != "" {
		option += fmt.Sprintf("&direct=%s", direct)
	}
	if option != "" {
		url += option
	}
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.GetTradeOrderHistoryResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) GetTradePositionOpensAsync(data chan responseorder.GetTradePositionOpensResponse,
	contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.GET_METHOD, "/v5/trade/position/opens", nil)
	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if option != "" {
		url += option
	}
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.GetTradePositionOpensResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) GetTradePositionHistoryAsync(data chan responseorder.GetTradePositionHistoryResponse,
	contractCode string, contractType string, marginMode string, startTime string, endTime string, from int,
	limit int, direct string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.GET_METHOD, "/v5/trade/position/history", nil)
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
	if startTime != "" {
		option += fmt.Sprintf("&start_time=%s", startTime)
	}
	if endTime != "" {
		option += fmt.Sprintf("&end_time=%s", endTime)
	}
	option += fmt.Sprintf("&from=%s", from)
	option += fmt.Sprintf("&limit=%s", limit)
	if direct != "" {
		option += fmt.Sprintf("&direct=%s", direct)
	}
	if option != "" {
		url += option
	}
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.GetTradePositionHistoryResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) GetTradePositionLeverAsync(data chan responseorder.GetTradePositionLeverResponse,
	contractCode string, marginMode string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.GET_METHOD, "/v5/position/lever", nil)
	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
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
	result := responseorder.GetTradePositionLeverResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

func (oc *OrderClient) SwapTradePositionLeverAsync(data chan responseorder.SwapTradePositionLeverResponse, contractCode string,
	marginMode string, leverRate string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/v5/position/lever", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if marginMode != "" {
		content += fmt.Sprintf(",\"margin_mode\": \"%s\"", marginMode)
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
	result := responseorder.SwapTradePositionLeverResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", getErr)
	}
	data <- result
}

func (ac *AccountClient) GetTradePositionModeAsync(data chan responseorder.GetTradePositionModeResponse) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.GET_METHOD, "/api/v5/position/mode", nil)
	// option
	option := ""
	if option != "" {
		url += option
	}
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.GetTradePositionModeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

func (oc *OrderClient) SwapTradePositionModeAsync(data chan responseorder.GetTradePositionModeResponse, positionMode string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/api/v5/position/mode", nil)

	// content
	content := ""
	if positionMode != "" {
		content += fmt.Sprintf(",\"position_mode\": \"%s\"", positionMode)
	}

	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.GetTradePositionModeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", getErr)
	}
	data <- result
}

func (ac *AccountClient) GetTradePositionRiskLimitAsync(data chan responseorder.GetTradePositionRiskLimitResponse,
	contractCode string, marginMode string, positionSide string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.GET_METHOD, "/v5/position/riskLimit", nil)
	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if marginMode != "" {
		option += fmt.Sprintf("&margin_mode=%s", marginMode)
	}
	if positionSide != "" {
		option += fmt.Sprintf("&position_side=%s", positionSide)
	}
	if option != "" {
		url += option
	}
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.GetTradePositionRiskLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) GetTradeOrderAsync(data chan responseorder.GetTradeOrderResponse,
	contractCode string, marginMode string, orderId string, clientOrderId string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.GET_METHOD, "/api/v5/trade/order/history", nil)
	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if orderId != "" {
		option += fmt.Sprintf("&order_id=%s", orderId)
	}
	if clientOrderId != "" {
		option += fmt.Sprintf("&client_order_id=%s", clientOrderId)
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
	result := responseorder.GetTradeOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}
