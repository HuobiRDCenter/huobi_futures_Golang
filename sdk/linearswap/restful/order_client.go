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

func (oc *OrderClient) Init(accessKey string, secretKey string, host string, sign string) *OrderClient {
	if host == "" {
		host = linearswap.LINEAR_SWAP_DEFAULT_HOST
	}
	oc.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host, sign)
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
