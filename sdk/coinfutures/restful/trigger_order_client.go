package restful

import (
	"encoding/json"
	"fmt"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinfutures"
	requesttiggerorder "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinfutures/restful/request/triggerorder"
	responsetriggerorder "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinfutures/restful/response/triggerorder"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/reqbuilder"
)

type TriggerOrderClient struct {
	PUrlBuilder *reqbuilder.PrivateUrlBuilder
}

func (toc *TriggerOrderClient) Init(accessKey string, secretKey string, host string, sign string) *TriggerOrderClient {
	if host == "" {
		host = coinfutures.COIN_FUTURES_DEFAULT_HOST
	}
	toc.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host, sign)
	return toc
}

// ContractTriggerOrderAsync (Place Trigger Order)
func (toc *TriggerOrderClient) ContractTriggerOrderAsync(data chan responsetriggerorder.ContractTriggerOrderResponse, request requesttiggerorder.ContractTriggerOrderRequest) {
	url := toc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_trigger_order", nil)

	content, err := json.Marshal(request)
	if err != nil {
		log.Error("PlaceOrderRequest to json error: %v", err)
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.ContractTriggerOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractTriggerOrderResponse error: %s", getErr)
	}
	data <- result
}

// SwapTriggerCancelAsync (Cancel Trigger Order)
func (toc *TriggerOrderClient) SwapTriggerCancelAsync(data chan responsetriggerorder.SwapTriggerCancelResponse,
	symbol string, orderId string) {
	// url
	url := toc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_trigger_cancel", nil)

	// content
	content := ""
	if symbol != "" {
		content += fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if orderId != "" {
		content += fmt.Sprintf(",\"order_id\": \"%s\"", orderId)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.SwapTriggerCancelResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapTriggerCancelResponse error: %s", getErr)
	}
	data <- result
}

// ContractTriggerCancelAllAsync (Cancel All Trigger Orders)
func (toc *TriggerOrderClient) ContractTriggerCancelAllAsync(data chan responsetriggerorder.ContractTriggerCancelAllResponse,
	symbol string, contractCode string, contractType string, direction string, offset string) {
	// url
	url := toc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_trigger_cancelall", nil)

	// content
	content := ""
	if symbol != "" {
		content += fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if contractType != "" {
		content += fmt.Sprintf(",\"contract_type\": \"%s\"", contractType)
	}
	if direction != "" {
		content += fmt.Sprintf(",\"direction\": \"%s\"", direction)
	}
	if offset != "" {
		content += fmt.Sprintf(",\"offset\": \"%s\"", offset)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.ContractTriggerCancelAllResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractTriggerCancelAllResponse error: %s", getErr)
	}
	data <- result
}

// ContractTriggerOpenOrdersAsync (Query Trigger Order Open Orders)
func (toc *TriggerOrderClient) ContractTriggerOpenOrdersAsync(data chan responsetriggerorder.ContractTriggerOpenOrdersResponse,
	symbol string, contractCode string, pageIndex string, pageSize string, tradeType string) {
	// url
	url := toc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_trigger_openorders", nil)

	// content
	content := ""
	if symbol != "" {
		content += fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if pageIndex != "" {
		content += fmt.Sprintf(",\"page_index\": \"%s\"", pageIndex)
	}
	if pageSize != "" {
		content += fmt.Sprintf(",\"page_size\": \"%s\"", pageSize)
	}
	if tradeType != "" {
		content += fmt.Sprintf(",\"trade_type\": \"%s\"", tradeType)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.ContractTriggerOpenOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractTriggerOpenOrdersResponse error: %s", getErr)
	}
	data <- result
}

// ContractTriggerHisOrdersAsync (Query Trigger Order History)
func (toc *TriggerOrderClient) ContractTriggerHisOrdersAsync(data chan responsetriggerorder.ContractTriggerHisOrdersResponse,
	symbol string, contractCode string, tradeType string, status string, createDate string, pageIndex string,
	pageSize string, sortBy string) {
	// url
	url := toc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_trigger_hisorders", nil)

	// content
	content := ""
	if symbol != "" {
		content += fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if pageIndex != "" {
		content += fmt.Sprintf(",\"page_index\": \"%s\"", pageIndex)
	}
	if pageSize != "" {
		content += fmt.Sprintf(",\"page_size\": \"%s\"", pageSize)
	}
	if tradeType != "" {
		content += fmt.Sprintf(",\"trade_type\": \"%s\"", tradeType)
	}
	if status != "" {
		content += fmt.Sprintf(",\"status\": \"%s\"", status)
	}
	if createDate != "" {
		content += fmt.Sprintf(",\"create_date\": \"%s\"", createDate)
	}
	if sortBy != "" {
		content += fmt.Sprintf(",\"sort_by\": \"%s\"", sortBy)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.ContractTriggerHisOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractTriggerHisOrdersResponse error: %s", getErr)
	}
	data <- result
}

// ContractTpslOrderAsync (Set a Take-profit and Stop-loss Order for an Existing Position)
func (toc *TriggerOrderClient) ContractTpslOrderAsync(data chan responsetriggerorder.ContractTpslOrderResponse, request requesttiggerorder.ContractTriggerOrderRequest) {
	url := toc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_tpsl_order", nil)

	content, err := json.Marshal(request)
	if err != nil {
		log.Error("PlaceOrderRequest to json error: %v", err)
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.ContractTpslOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractTpslOrderResponse error: %s", getErr)
	}
	data <- result
}

// ContractTpslCancelAsync (Cancel a Take-profit and Stop-loss Order)
func (toc *TriggerOrderClient) ContractTpslCancelAsync(data chan responsetriggerorder.ContractTpslCancelResponse,
	symbol string, orderId string) {
	// url
	url := toc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_tpsl_cancel", nil)

	// content
	content := ""
	if symbol != "" {
		content += fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if orderId != "" {
		content += fmt.Sprintf(",\"order_id\": \"%s\"", orderId)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.ContractTpslCancelResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractTpslCancelResponse error: %s", getErr)
	}
	data <- result
}

// ContractTpslCancelAllAsync (Cancel all Take-profit and Stop-loss Orders)
func (toc *TriggerOrderClient) ContractTpslCancelAllAsync(data chan responsetriggerorder.ContractTpslCancelAllResponse,
	symbol string, contractCode string, contractType string, direction string) {
	// url
	url := toc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_tpsl_cancelall", nil)

	// content
	content := ""
	if symbol != "" {
		content += fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if contractType != "" {
		content += fmt.Sprintf(",\"contract_type\": \"%s\"", contractType)
	}
	if direction != "" {
		content += fmt.Sprintf(",\"direction\": \"%s\"", direction)
	}

	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.ContractTpslCancelAllResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractTpslCancelAllResponse error: %s", getErr)
	}
	data <- result
}

// ContractTpslOpenOrdersAsync (Query Open Take-profit and Stop-loss Orders)
func (toc *TriggerOrderClient) ContractTpslOpenOrdersAsync(data chan responsetriggerorder.ContractTpslOpenOrdersResponse,
	symbol string, contractCode string, pageIndex string, pageSize string, tradeType string) {
	// url
	url := toc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_tpsl_openorders", nil)

	// content
	content := ""
	if symbol != "" {
		content += fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if pageIndex != "" {
		content += fmt.Sprintf(",\"page_index\": \"%s\"", pageIndex)
	}
	if pageSize != "" {
		content += fmt.Sprintf(",\"page_size\": \"%s\"", pageSize)
	}
	if tradeType != "" {
		content += fmt.Sprintf(",\"trade_type\": \"%s\"", tradeType)
	}

	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.ContractTpslOpenOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractTpslOpenOrdersResponse error: %s", getErr)
	}
	data <- result
}

// ContractTpslHisOrdersAsync (Query Take-profit and Stop-loss History Orders)
func (toc *TriggerOrderClient) ContractTpslHisOrdersAsync(data chan responsetriggerorder.ContractTpslHisOrdersResponse,
	symbol string, contractCode string, status string, createDate string, pageIndex string, pageSize string, sortBy string) {
	// url
	url := toc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_tpsl_hisorders", nil)

	// content
	content := ""
	if symbol != "" {
		content += fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if status != "" {
		content += fmt.Sprintf(",\"status\": \"%s\"", status)
	}
	if createDate != "" {
		content += fmt.Sprintf(",\"create_date\": \"%s\"", createDate)
	}
	if pageIndex != "" {
		content += fmt.Sprintf(",\"page_index\": \"%s\"", pageIndex)
	}
	if pageSize != "" {
		content += fmt.Sprintf(",\"page_size\": \"%s\"", pageSize)
	}
	if sortBy != "" {
		content += fmt.Sprintf(",\"sort_by\": \"%s\"", sortBy)
	}

	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.ContractTpslHisOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractTpslHisOrdersResponse error: %s", getErr)
	}
	data <- result
}

// ContractRelationTpslOrderAsync (Query Info Of Take-profit and Stop-loss Order That Related To Position Opening Order)
func (toc *TriggerOrderClient) ContractRelationTpslOrderAsync(data chan responsetriggerorder.ContractRelationTpslOrderResponse,
	symbol string, orderId string) {
	// url
	url := toc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_relation_tpsl_order", nil)

	// content
	content := ""
	if symbol != "" {
		content += fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if orderId != "" {
		content += fmt.Sprintf(",\"order_id\": \"%s\"", orderId)
	}

	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.ContractRelationTpslOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractRelationTpslOrderResponse error: %s", getErr)
	}
	data <- result
}

// ContractTrackOrderAsync (Place a Trailing Order)
func (toc *TriggerOrderClient) ContractTrackOrderAsync(data chan responsetriggerorder.ContractTrackOrderResponse,
	symbol string, contractType string, contractCode string, direction string, offset string, leverRate string, volume string,
	callbackRate string, activePrice string, orderPriceType string) {
	// url
	url := toc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_track_order", nil)

	// content
	content := ""
	if symbol != "" {
		content += fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if contractType != "" {
		content += fmt.Sprintf(",\"contract_type\": \"%s\"", contractType)
	}
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if direction != "" {
		content += fmt.Sprintf(",\"direction\": \"%s\"", direction)
	}
	if offset != "" {
		content += fmt.Sprintf(",\"offset\": \"%s\"", offset)
	}
	if leverRate != "" {
		content += fmt.Sprintf(",\"lever_rate\": \"%s\"", leverRate)
	}
	if volume != "" {
		content += fmt.Sprintf(",\"volume\": \"%s\"", volume)
	}
	if callbackRate != "" {
		content += fmt.Sprintf(",\"callback_rate\": \"%s\"", callbackRate)
	}
	if activePrice != "" {
		content += fmt.Sprintf(",\"active_price\": \"%s\"", activePrice)
	}
	if orderPriceType != "" {
		content += fmt.Sprintf(",\"order_price_type\": \"%s\"", orderPriceType)
	}

	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.ContractTrackOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractTrackOrderResponse error: %s", getErr)
	}
	data <- result
}

// ContractTrackCancelAsync (Cancel a Trailing Order)
func (toc *TriggerOrderClient) ContractTrackCancelAsync(data chan responsetriggerorder.ContractTrackCancelResponse,
	symbol string, orderId string) {
	// url
	url := toc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_track_cancel", nil)

	// content
	content := ""
	if symbol != "" {
		content += fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if orderId != "" {
		content += fmt.Sprintf(",\"order_id\": \"%s\"", orderId)
	}

	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.ContractTrackCancelResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractTrackCancelResponse error: %s", getErr)
	}
	data <- result
}

// ContractTrackCancelAllAsync (Cancel All Trailing Orders)
func (toc *TriggerOrderClient) ContractTrackCancelAllAsync(data chan responsetriggerorder.ContractTrackCancelAllResponse,
	symbol string, contractCode string, contractType string, direction string, offset string) {
	// url
	url := toc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_track_cancelall", nil)

	// content
	content := ""
	if symbol != "" {
		content += fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if contractType != "" {
		content += fmt.Sprintf(",\"contract_type\": \"%s\"", contractType)
	}
	if direction != "" {
		content += fmt.Sprintf(",\"direction\": \"%s\"", direction)
	}
	if offset != "" {
		content += fmt.Sprintf(",\"offset\": \"%s\"", offset)
	}

	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.ContractTrackCancelAllResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractTrackCancelAllResponse error: %s", getErr)
	}
	data <- result
}

// ContractTrackOpenOrdersAsync (Current unfilled trailing order acquisition)
func (toc *TriggerOrderClient) ContractTrackOpenOrdersAsync(data chan responsetriggerorder.ContractTrackOpenOrdersResponse,
	symbol string, contractCode string, tradeType string, pageIndex string, pageSize string) {
	// url
	url := toc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_track_openorders", nil)

	// content
	content := ""
	if symbol != "" {
		content += fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if tradeType != "" {
		content += fmt.Sprintf(",\"trade_type\": \"%s\"", tradeType)
	}
	if pageIndex != "" {
		content += fmt.Sprintf(",\"page_index\": \"%s\"", pageIndex)
	}
	if pageSize != "" {
		content += fmt.Sprintf(",\"page_size\": \"%s\"", pageSize)
	}

	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.ContractTrackOpenOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractTrackOpenOrdersResponse error: %s", getErr)
	}
	data <- result
}

// ContractTrackHisOrdersAsync (Get History Trailing Orders)
func (toc *TriggerOrderClient) ContractTrackHisOrdersAsync(data chan responsetriggerorder.ContractTrackHisOrdersResponse,
	symbol string, contractCode string, status string, tradeType string, createDate string, pageIndex string,
	pageSize string, sortBy string) {
	// url
	url := toc.PUrlBuilder.Build(coinfutures.POST_METHOD, "/api/v1/contract_track_hisorders", nil)

	// content
	content := ""
	if symbol != "" {
		content += fmt.Sprintf(",\"symbol\": \"%s\"", symbol)
	}
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if status != "" {
		content += fmt.Sprintf(",\"status\": \"%s\"", status)
	}
	if tradeType != "" {
		content += fmt.Sprintf(",\"trade_type\": \"%s\"", tradeType)
	}
	if createDate != "" {
		content += fmt.Sprintf(",\"create_date\": \"%s\"", createDate)
	}
	if pageIndex != "" {
		content += fmt.Sprintf(",\"page_index\": \"%s\"", pageIndex)
	}
	if pageSize != "" {
		content += fmt.Sprintf(",\"page_size\": \"%s\"", pageSize)
	}
	if sortBy != "" {
		content += fmt.Sprintf(",\"sort_by\": \"%s\"", sortBy)
	}

	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.ContractTrackHisOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to ContractTrackHisOrdersResponse error: %s", getErr)
	}
	data <- result
}
