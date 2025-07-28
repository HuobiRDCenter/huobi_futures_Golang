package restful

import (
	"encoding/json"
	"fmt"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinswap"
	responsetriggerorder "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinswap/restful/response/triggerorder"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/reqbuilder"
)

type TriggerOrderClient struct {
	PUrlBuilder *reqbuilder.PrivateUrlBuilder
}

func (toc *TriggerOrderClient) Init(accessKey string, secretKey string, host string, sign string) *TriggerOrderClient {
	if host == "" {
		host = coinswap.COIN_SWAP_DEFAULT_HOST
	}
	toc.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host, sign)
	return toc
}

// SwapTriggerOrderAsync (Place Trigger Order)
func (toc *TriggerOrderClient) SwapTriggerOrderAsync(data chan responsetriggerorder.SwapTriggerOrderResponse,
	contractCode string, triggerType string, triggerPrice string, orderPrice string, orderPriceType string,
	volume string, direction string, offset string, leverRate string) {
	// url
	url := toc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_trigger_order", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if triggerType != "" {
		content += fmt.Sprintf(",\"trigger_type\": \"%s\"", triggerType)
	}
	if triggerPrice != "" {
		content += fmt.Sprintf(",\"trigger_price\": \"%s\"", triggerPrice)
	}
	if orderPrice != "" {
		content += fmt.Sprintf(",\"order_price\": \"%s\"", orderPrice)
	}
	if orderPriceType != "" {
		content += fmt.Sprintf(",\"order_price_type\": \"%s\"", orderPriceType)
	}
	if volume != "" {
		content += fmt.Sprintf(",\"volume\": \"%s\"", volume)
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
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.SwapTriggerOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapTriggerOrderResponse error: %s", getErr)
	}
	data <- result
}

// SwapTriggerCancelAsync (Cancel Trigger Order)
func (toc *TriggerOrderClient) SwapTriggerCancelAsync(data chan responsetriggerorder.SwapTriggerCancelResponse,
	contractCode string, orderId string) {
	// url
	url := toc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_trigger_cancel", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
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

// SwapTriggerCancelAllAsync (Cancel All Trigger Orders)
func (toc *TriggerOrderClient) SwapTriggerCancelAllAsync(data chan responsetriggerorder.SwapTriggerCancelAllResponse,
	contractCode string, direction string, offset string) {
	// url
	url := toc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_trigger_cancelall", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
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
	result := responsetriggerorder.SwapTriggerCancelAllResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapTriggerCancelAllResponse error: %s", getErr)
	}
	data <- result
}

// SwapTriggerOpenOrdersAsync (Query Trigger Order Open Orders)
func (toc *TriggerOrderClient) SwapTriggerOpenOrdersAsync(data chan responsetriggerorder.SwapTriggerOpenOrdersResponse,
	contractCode string, pageIndex string, pageSize string, tradeType string) {
	// url
	url := toc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_trigger_openorders", nil)

	// content
	content := ""
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
	result := responsetriggerorder.SwapTriggerOpenOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapTriggerOpenOrdersResponse error: %s", getErr)
	}
	data <- result
}

// SwapTriggerHisOrdersAsync (Query Trigger Order History)
func (toc *TriggerOrderClient) SwapTriggerHisOrdersAsync(data chan responsetriggerorder.SwapTriggerHisOrdersResponse,
	contractCode string, tradeType string, status string, createDate string, pageIndex string, pageSize string, sortBy string) {
	// url
	url := toc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_trigger_hisorders", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
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
	result := responsetriggerorder.SwapTriggerHisOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapTriggerHisOrdersResponse error: %s", getErr)
	}
	data <- result
}

// SwapTpslOrderAsync (Set a Take-profit and Stop-loss Order for an Existing Position)
func (toc *TriggerOrderClient) SwapTpslOrderAsync(data chan responsetriggerorder.SwapTpslOrderResponse,
	contractCode string, direction string, volume string, tpTriggerPrice string, tpOrderPrice string,
	tpOrderPriceType string, slTriggerPrice string, slOrderPrice string, slOrderPriceType string) {
	// url
	url := toc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_tpsl_order", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if direction != "" {
		content += fmt.Sprintf(",\"direction\": \"%s\"", direction)
	}
	if volume != "" {
		content += fmt.Sprintf(",\"volume\": \"%s\"", volume)
	}
	if tpTriggerPrice != "" {
		content += fmt.Sprintf(",\"tp_trigger_price\": \"%s\"", tpTriggerPrice)
	}
	if tpOrderPrice != "" {
		content += fmt.Sprintf(",\"tp_order_price\": \"%s\"", tpOrderPrice)
	}
	if tpOrderPriceType != "" {
		content += fmt.Sprintf(",\"tp_order_price_type\": \"%s\"", tpOrderPriceType)
	}
	if slTriggerPrice != "" {
		content += fmt.Sprintf(",\"sl_trigger_price\": \"%s\"", slTriggerPrice)
	}
	if slOrderPrice != "" {
		content += fmt.Sprintf(",\"sl_order_price\": \"%s\"", slOrderPrice)
	}
	if slOrderPriceType != "" {
		content += fmt.Sprintf(",\"sl_order_price_type\": \"%s\"", slOrderPriceType)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.SwapTpslOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapTpslOrderResponse error: %s", getErr)
	}
	data <- result
}

// SwapTpslCancelAsync (Cancel a Take-profit and Stop-loss Order)
func (toc *TriggerOrderClient) SwapTpslCancelAsync(data chan responsetriggerorder.SwapTpslCancelResponse,
	contractCode string, orderId string) {
	// url
	url := toc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_tpsl_cancel", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
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
	result := responsetriggerorder.SwapTpslCancelResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapTpslCancelResponse error: %s", getErr)
	}
	data <- result
}

// SwapTpslCancelAllAsync (Cancel all Take-profit and Stop-loss Orders)
func (toc *TriggerOrderClient) SwapTpslCancelAllAsync(data chan responsetriggerorder.SwapTpslCancelAllResponse,
	contractCode string, direction string) {
	// url
	url := toc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_tpsl_cancelall", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
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
	result := responsetriggerorder.SwapTpslCancelAllResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapTpslCancelAllResponse error: %s", getErr)
	}
	data <- result
}

// SwapTpslOpenOrdersAsync (Query Open Take-profit and Stop-loss Orders)
func (toc *TriggerOrderClient) SwapTpslOpenOrdersAsync(data chan responsetriggerorder.SwapTpslOpenOrdersResponse,
	contractCode string, pageIndex string, pageSize string, tradeType string) {
	// url
	url := toc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_tpsl_openorders", nil)

	// content
	content := ""
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
	result := responsetriggerorder.SwapTpslOpenOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapTpslOpenOrdersResponse error: %s", getErr)
	}
	data <- result
}

// SwapTpslHisOrdersAsync (Query Take-profit and Stop-loss History Orders)
func (toc *TriggerOrderClient) SwapTpslHisOrdersAsync(data chan responsetriggerorder.SwapTpslHisOrdersResponse,
	contractCode string, status string, createDate string, pageIndex string, pageSize string, sortBy string) {
	// url
	url := toc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_tpsl_hisorders", nil)

	// content
	content := ""
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
	result := responsetriggerorder.SwapTpslHisOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapTpslHisOrdersResponse error: %s", getErr)
	}
	data <- result
}

// SwapRelationTpslOrderAsync (Query Info Of Take-profit and Stop-loss Order That Related To Position Opening Order)
func (toc *TriggerOrderClient) SwapRelationTpslOrderAsync(data chan responsetriggerorder.SwapRelationTpslOrderResponse,
	contractCode string, orderId string) {
	// url
	url := toc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_relation_tpsl_order", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
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
	result := responsetriggerorder.SwapRelationTpslOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapRelationTpslOrderResponse error: %s", getErr)
	}
	data <- result
}

// SwapTrackOrderAsync (Place a Trailing Order)
func (toc *TriggerOrderClient) SwapTrackOrderAsync(data chan responsetriggerorder.SwapTrackOrderResponse,
	contractCode string, direction string, offset string, leverRate string, volume string, callbackRate string,
	activePrice string, orderPriceType string) {
	// url
	url := toc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_track_order", nil)

	// content
	content := ""
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
	result := responsetriggerorder.SwapTrackOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapTrackOrderResponse error: %s", getErr)
	}
	data <- result
}

// SwapTrackCancelAsync (Cancel a Trailing Order)
func (toc *TriggerOrderClient) SwapTrackCancelAsync(data chan responsetriggerorder.SwapTrackCancelResponse,
	contractCode string, orderId string) {
	// url
	url := toc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_track_cancel", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
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
	result := responsetriggerorder.SwapTrackCancelResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapTrackCancelResponse error: %s", getErr)
	}
	data <- result
}

// SwapTrackCancelAllAsync (Cancel All Trailing Orders)
func (toc *TriggerOrderClient) SwapTrackCancelAllAsync(data chan responsetriggerorder.SwapTrackCancelAllResponse,
	contractCode string, direction string, offset string) {
	// url
	url := toc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_track_cancelall", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
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
	result := responsetriggerorder.SwapTrackCancelAllResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapTrackCancelAllResponse error: %s", getErr)
	}
	data <- result
}

// SwapTrackOpenOrdersAsync (Current unfilled trailing order acquisition)
func (toc *TriggerOrderClient) SwapTrackOpenOrdersAsync(data chan responsetriggerorder.SwapTrackOpenOrdersResponse,
	contractCode string, tradeType string, pageIndex string, pageSize string) {
	// url
	url := toc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_track_openorders", nil)

	// content
	content := ""
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
	result := responsetriggerorder.SwapTrackOpenOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapTrackOpenOrdersResponse error: %s", getErr)
	}
	data <- result
}

// SwapTrackHisOrdersAsync (Get History Trailing Orders)
func (toc *TriggerOrderClient) SwapTrackHisOrdersAsync(data chan responsetriggerorder.SwapTrackHisOrdersResponse,
	contractCode string, status string, tradeType string, createDate string, pageIndex string, pageSize string, sortBy string) {
	// url
	url := toc.PUrlBuilder.Build(coinswap.POST_METHOD, "/swap-api/v1/swap_track_hisorders", nil)

	// content
	content := ""
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
	result := responsetriggerorder.SwapTrackHisOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapTrackHisOrdersResponse error: %s", getErr)
	}
	data <- result
}
