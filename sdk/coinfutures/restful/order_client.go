package restful

import (
	"encoding/json"
	"fmt"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinfutures"
	requestorder "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinfutures/restful/request/order"
	responseorder "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinfutures/restful/response/order"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/reqbuilder"
)

type OrderClient struct {
	PUrlBuilder *reqbuilder.PrivateUrlBuilder
}

func (oc *OrderClient) Init(accessKey string, secretKey string, host string, sign string) *OrderClient {
	if host == "" {
		host = coinfutures.COIN_FUTURES_DEFAULT_HOST
	}
	oc.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host, sign)
	return oc
}

// ContractCancelAfterAsync ( Automatic Order Cancellation)
func (oc *OrderClient) ContractCancelAfterAsync(data chan responseorder.ContractCancelAfterResponse, onOff string,
	timeOut string) {
	url := oc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract-cancel-after", nil)

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
	result := responseorder.ContractCancelAfterResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractCancelAfterResponse error: %s", getErr)
	}
	data <- result
}

// ContractOrderAsync (Place an Order)
func (oc *OrderClient) ContractOrderAsync(data chan responseorder.ContractOrderResponse, request requestorder.ContractOrderRequest) {
	url := oc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_order", nil)

	content, err := json.Marshal(request)
	if err != nil {
		log.Error("PlaceOrderRequest to json error: %v", err)
	}
	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.ContractOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractOrderResponse error: %s", getErr)
	}
	data <- result
}

// ContractBatchOrderAsync (Place a Batch of Orders)
func (oc *OrderClient) ContractBatchOrderAsync(data chan responseorder.ContractBatchOrderResponse, request requestorder.BatchContractOrderRequest) {
	url := oc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_batchorder", nil)

	content, err := json.Marshal(request)
	if err != nil {
		log.Error("PlaceOrderRequest to json error: %v", err)
	}
	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.ContractBatchOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractBatchOrderResponse error: %s", getErr)
	}
	data <- result
}

// ContractCancelAsync (Cancel an Order)
func (oc *OrderClient) ContractCancelAsync(data chan responseorder.ContractCancelResponse, orderId string,
	clientOrderId string, symbol string) {
	url := oc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_cancel", nil)

	// content
	content := ""
	if orderId != "" {
		content = fmt.Sprintf(",\"order_id\": \"%s\"", orderId)
	}
	if clientOrderId != "" {
		content = fmt.Sprintf(",\"client_order_id\": \"%s\"", clientOrderId)
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
	result := responseorder.ContractCancelResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractCancelResponse error: %s", getErr)
	}
	data <- result
}

// ContractCancelAllAsync (Cancel All Orders)
func (oc *OrderClient) ContractCancelAllAsync(data chan responseorder.ContractCancelAllResponse, symbol string,
	contractCode string, contractType string, direction string, offset string) {
	url := oc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_cancelall", nil)

	// content
	content := ""
	if symbol != "" {
		content = fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if contractType != "" {
		content = fmt.Sprintf(",\"contract_type\": \"%s\"", contractType)
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
	result := responseorder.ContractCancelAllResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractCancelAllResponse error: %s", getErr)
	}
	data <- result
}

// ContractSwitchLeverRateAsync (Switch Leverage)
func (oc *OrderClient) ContractSwitchLeverRateAsync(data chan responseorder.ContractSwitchLeverRateResponse, symbol string,
	leverRate string) {
	url := oc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_switch_lever_rate", nil)

	// content
	content := ""
	if symbol != "" {
		content = fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
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
	result := responseorder.ContractSwitchLeverRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractSwitchLeverRateResponse error: %s", getErr)
	}
	data <- result
}

// ContractOrderInfoAsync (Get Information of an Order)
func (oc *OrderClient) ContractOrderInfoAsync(data chan responseorder.ContractOrderInfoResponse, orderId string,
	clientOrderId string, symbol string) {
	url := oc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_order_info", nil)

	// content
	content := ""
	if orderId != "" {
		content = fmt.Sprintf(",\"order_id\": \"%s\"", orderId)
	}
	if clientOrderId != "" {
		content = fmt.Sprintf(",\"client_order_id\": \"%s\"", clientOrderId)
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
	result := responseorder.ContractOrderInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractOrderInfoResponse error: %s", getErr)
	}
	data <- result
}

// ContractOrderDetailAsync (Order Details Acquisition)
func (oc *OrderClient) ContractOrderDetailAsync(data chan responseorder.ContractOrderDetailResponse, symbol string,
	orderId string, createAt string, orderType string, pageIndex string, pageSize string) {
	url := oc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_order_detail", nil)

	// content
	content := ""
	if symbol != "" {
		content = fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if orderId != "" {
		content = fmt.Sprintf(",\"order_id\": \"%s\"", orderId)
	}
	if createAt != "" {
		content = fmt.Sprintf(",\"created_at\": \"%s\"", createAt)
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
	result := responseorder.ContractOrderDetailResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractOrderDetailResponse error: %s", getErr)
	}
	data <- result
}

// ContractOpenOrdersAsync (Query Open Orders)
func (oc *OrderClient) ContractOpenOrdersAsync(data chan responseorder.ContractOpenOrdersResponse, symbol string,
	pageIndex string, pageSize string, sortBy string, tradeType string) {
	url := oc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_openorders", nil)

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
	result := responseorder.ContractOpenOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractOpenOrdersResponse error: %s", getErr)
	}
	data <- result
}

// ContractHisordersAsync (Get History Orders(New))
func (oc *OrderClient) ContractHisordersAsync(data chan responseorder.ContractHisordersResponse, symbol string,
	tradeType string, fcType string, status string, orderType string, sortBy string, startTime string, endTime string,
	direct string, fromId string) {
	url := oc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v3/contract_hisorders", nil)

	// content
	content := ""
	if symbol != "" {
		content = fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
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
	if orderType != "" {
		content = fmt.Sprintf(",\"order_type\": \"%s\"", orderType)
	}
	if sortBy != "" {
		content = fmt.Sprintf(",\"sort_by\": \"%s\"", sortBy)
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
	result := responseorder.ContractHisordersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractHisordersResponse error: %s", getErr)
	}
	data <- result
}

// ContractHisOrdersExactAsync (Query History Orders via Multiple fields(New))
func (oc *OrderClient) ContractHisOrdersExactAsync(data chan responseorder.ContractHisOrdersExactResponse, symbol string,
	tradeType string, fcType string, status string, contract string, orderType string, sortBy string, startTime string,
	endTime string, direct string, fromId string) {
	url := oc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v3/contract_hisorders_exact", nil)

	// content
	content := ""
	if symbol != "" {
		content = fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
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
	if contract != "" {
		content = fmt.Sprintf(",\"contract\": \"%s\"", contract)
	}
	if orderType != "" {
		content = fmt.Sprintf(",\"order_type\": \"%s\"", orderType)
	}
	if sortBy != "" {
		content = fmt.Sprintf(",\"sort_by\": \"%s\"", sortBy)
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
	result := responseorder.ContractHisOrdersExactResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractHisOrdersExactResponse error: %s", getErr)
	}
	data <- result
}

// ContractMatchResultsAsync (Get History Match Results(New))
func (oc *OrderClient) ContractMatchResultsAsync(data chan responseorder.ContractMatchResultsResponse, symbol string,
	tradeType string, startTime string, endTime string, direct string, fromId string) {
	url := oc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v3/contract_matchresults", nil)

	// content
	content := ""
	if symbol != "" {
		content = fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
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
	result := responseorder.ContractMatchResultsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractMatchResultsResponse error: %s", getErr)
	}
	data <- result
}

// ContractMatchResultsExactAsync (Query history transactions via multiple fields(New))
func (oc *OrderClient) ContractMatchResultsExactAsync(data chan responseorder.ContractMatchResultsExactResponse, symbol string,
	contract string, tradeType string, startTime string, endTime string, direct string, fromId string) {
	url := oc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v3/contract_matchresults_exact", nil)

	// content
	content := ""
	if symbol != "" {
		content = fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
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
	result := responseorder.ContractMatchResultsExactResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractMatchResultsExactResponse error: %s", getErr)
	}
	data <- result
}

// LightningClosePositionAsync (Place Flash Close Order)
func (oc *OrderClient) LightningClosePositionAsync(data chan responseorder.LightningClosePositionResponse, symbol string,
	contractType string, contractCode string, volume string, direction string, clientOrderId string, orderPriceType string) {
	url := oc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/lightning_close_position", nil)

	// content
	content := ""
	if symbol != "" {
		content = fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if contractType != "" {
		content = fmt.Sprintf(",\"contract_type\": \"%s\"", contractType)
	}
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
	result := responseorder.LightningClosePositionResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to LightningClosePositionResponse error: %s", getErr)
	}
	data <- result
}
