package restful

import (
    "encoding/json"
    "fmt"
    "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap"
    "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/request/algo"
    responsealgo "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/response/algo"
    "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
    "github.com/HuobiRDCenter/huobi_futures_Golang/sdk/reqbuilder"
    "strings"
)

type AlgoClient struct {
    PUrlBuilder *reqbuilder.PrivateUrlBuilder
}

func (ac *AlgoClient) Init(accessKey string, secretKey string, host string) *AlgoClient {
    if host == "" {
        host = linearswap.LINEAR_SWAP_DEFAULT_HOST
    }
    ac.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
    return ac
}

// AlgoOrderAsync 策略订单委托
func (ac *AlgoClient) AlgoOrderAsync(data chan responsealgo.AlgoOrderResponse, request algo.AlgoOrderRequest) {
    url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/v5/algo/order", nil)

    // 构建请求内容
    content := ""

    // 必填字段
    if request.ContractCode != "" {
        content += fmt.Sprintf(",\"contract_code\":\"%s\"", request.ContractCode)
    }
    if request.Type != "" {
        content += fmt.Sprintf(",\"type\":\"%s\"", request.Type)
    }
    if request.PositionSide != "" {
        content += fmt.Sprintf(",\"position_side\":\"%s\"", request.PositionSide)
    }
    if request.Side != "" {
        content += fmt.Sprintf(",\"side\":\"%s\"", request.Side)
    }
    if request.MarginMode != "" {
        content += fmt.Sprintf(",\"margin_mode\":\"%s\"", request.MarginMode)
    }

    // 可选字段
    if request.AlgoClientOrderId != "" {
        content += fmt.Sprintf(",\"algo_client_order_id\":\"%s\"", request.AlgoClientOrderId)
    }
    if request.Volume != "" {
        content += fmt.Sprintf(",\"volume\":\"%s\"", request.Volume)
    }

    // 止盈相关参数
    if request.TpTriggerPrice != "" {
        content += fmt.Sprintf(",\"tp_trigger_price\":\"%s\"", request.TpTriggerPrice)
    }
    if request.TpOrderPrice != "" {
        content += fmt.Sprintf(",\"tp_order_price\":\"%s\"", request.TpOrderPrice)
    }
    if request.TpType != "" {
        content += fmt.Sprintf(",\"tp_type\":\"%s\"", request.TpType)
    }
    if request.TpTriggerPriceType != "" {
        content += fmt.Sprintf(",\"tp_trigger_price_type\":\"%s\"", request.TpTriggerPriceType)
    }

    // 止损相关参数
    if request.SlTriggerPrice != "" {
        content += fmt.Sprintf(",\"sl_trigger_price\":\"%s\"", request.SlTriggerPrice)
    }
    if request.SlOrderPrice != "" {
        content += fmt.Sprintf(",\"sl_order_price\":\"%s\"", request.SlOrderPrice)
    }
    if request.SlType != "" {
        content += fmt.Sprintf(",\"sl_type\":\"%s\"", request.SlType)
    }
    if request.SlTriggerPriceType != "" {
        content += fmt.Sprintf(",\"sl_trigger_price_type\":\"%s\"", request.SlTriggerPriceType)
    }

    // 计划委托相关参数
    if request.Price != "" {
        content += fmt.Sprintf(",\"price\":\"%s\"", request.Price)
    }
    if request.PriceMatch != "" {
        content += fmt.Sprintf(",\"price_match\":\"%s\"", request.PriceMatch)
    }
    if request.TriggerPrice != "" {
        content += fmt.Sprintf(",\"trigger_price\":\"%s\"", request.TriggerPrice)
    }
    if request.TriggerPriceType != "" {
        content += fmt.Sprintf(",\"trigger_price_type\":\"%s\"", request.TriggerPriceType)
    }

    // 跟踪委托相关参数
    if request.ActivePrice != "" {
        content += fmt.Sprintf(",\"active_price\":\"%s\"", request.ActivePrice)
    }
    if request.OrderPriceType != "" {
        content += fmt.Sprintf(",\"order_price_type\":\"%s\"", request.OrderPriceType)
    }
    if request.CallbackRate != "" {
        content += fmt.Sprintf(",\"callback_rate\":\"%s\"", request.CallbackRate)
    }
    if request.ReduceOnly != nil {
        content += fmt.Sprintf(",\"reduce_only\":%v", *request.ReduceOnly)
    }

    if content != "" {
        content = fmt.Sprintf("{%s}", content[1:])
    }

    getResp, getErr := reqbuilder.HttpPost(url, content)
    if getErr != nil {
        log.Error("http post error: %s", getErr)
        return
    }

    var result responsealgo.AlgoOrderResponse
    jsonErr := json.Unmarshal([]byte(getResp), &result)
    if jsonErr != nil {
        log.Error("convert json to AlgoOrderResponse error: %s", jsonErr)
        return
    }

    data <- result
}

// CancelAlgoOrderAsync 撤销策略订单
func (ac *AlgoClient) CancelAlgoOrderAsync(data chan responsealgo.CancelAlgoOrdersResponse, request algo.CancelAlgoOrdersRequest) {
    url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/v5/algo/cancel_orders", nil)

    content := ""

    if request.AlgoId != "" {
        content += fmt.Sprintf(",\"algo_id\":\"%s\"", request.AlgoId)
    }
    if request.AlgoClientOrderId != "" {
        content += fmt.Sprintf(",\"algo_client_order_id\":\"%s\"", request.AlgoClientOrderId)
    }
    if request.ContractCode != "" {
        content += fmt.Sprintf(",\"contract_code\":\"%s\"", request.ContractCode)
    }

    if content != "" {
        // V5 策略撤单的 body 是**顶层数组** [{...}]，对象形式服务端直接 400
        // （Cannot deserialize ArrayList<CancelAlgoOrdersReq> from Object value）。
        // 口径对齐 C++ AlgoClient.hpp「单对象包装进顶层数组」。
        content = fmt.Sprintf("[{%s}]", content[1:])
    }

    getResp, getErr := reqbuilder.HttpPost(url, content)
    if getErr != nil {
        log.Error("http post error: %s", getErr)
        return
    }

    var result responsealgo.CancelAlgoOrdersResponse
    jsonErr := json.Unmarshal([]byte(getResp), &result)
    if jsonErr != nil {
        log.Error("convert json to CancelAlgoOrdersResponse error: %s", jsonErr)
        return
    }

    data <- result
}

// QueryAlgoOrderAsync 查询策略订单
func (ac *AlgoClient) QueryAlgoOrderAsync(data chan responsealgo.QueryAlgoOrdersResponse, request algo.QueryAlgoOrdersRequest) {
    // 参数校验
    if request.Type == "" {
        log.Error("type参数为必填项")
        return
    }

    // option
    option := ""
    if request.Type != "" {
        option += fmt.Sprintf("?type=%s", request.Type)
    }
    if request.AlgoId != "" {
        if option == "" {
            option += fmt.Sprintf("?algo_id=%s", request.AlgoId)
        } else {
            option += fmt.Sprintf("&algo_id=%s", request.AlgoId)
        }
    }
    if request.AlgoClientOrderId != "" {
        if option == "" {
            option += fmt.Sprintf("?algo_client_order_id=%s", request.AlgoClientOrderId)
        } else {
            option += fmt.Sprintf("&algo_client_order_id=%s", request.AlgoClientOrderId)
        }
    }
    if request.ContractCode != "" {
        if option == "" {
            option += fmt.Sprintf("?contract_code=%s", request.ContractCode)
        } else {
            option += fmt.Sprintf("&contract_code=%s", request.ContractCode)
        }
    }

    url := ac.PUrlBuilder.BuildWithQuery(linearswap.GET_METHOD, "/v5/algo/order", nil, option)

    getResp, getErr := reqbuilder.HttpGet(url)
    if getErr != nil {
        log.Error("http get error: %s", getErr)
        return
    }

    var result responsealgo.QueryAlgoOrdersResponse
    jsonErr := json.Unmarshal([]byte(getResp), &result)
    if jsonErr != nil {
        log.Error("convert json to QueryAlgoOrdersResponse error: %s", jsonErr)
        return
    }

    data <- result
}

// QueryOpenAlgoOrdersAsync 查询当前未触发策略订单
func (ac *AlgoClient) QueryOpenAlgoOrdersAsync(data chan responsealgo.QueryOpenAlgoOrdersResponse, request algo.QueryOpenAlgoOrdersRequest) {

    // option
    option := ""
    if request.Type != "" {
        option += fmt.Sprintf("?type=%s", request.Type)
    }
    if request.ContractCode != "" {
        if option == "" {
            option += fmt.Sprintf("?contract_code=%s", request.ContractCode)
        } else {
            option += fmt.Sprintf("&contract_code=%s", request.ContractCode)
        }
    }
    if request.AlgoId != "" {
        if option == "" {
            option += fmt.Sprintf("?algo_id=%s", request.AlgoId)
        } else {
            option += fmt.Sprintf("&algo_id=%s", request.AlgoId)
        }
    }
    if request.AlgoClientOrderId != "" {
        if option == "" {
            option += fmt.Sprintf("?algo_client_order_id=%s", request.AlgoClientOrderId)
        } else {
            option += fmt.Sprintf("&algo_client_order_id=%s", request.AlgoClientOrderId)
        }
    }
    if request.From != nil && *request.From > 0 {
        if option == "" {
            option += fmt.Sprintf("?from=%d", *request.From)
        } else {
            option += fmt.Sprintf("&from=%d", *request.From)
        }
    }
    if request.Limit != nil && *request.Limit > 0 {
        if option == "" {
            option += fmt.Sprintf("?limit=%d", *request.Limit)
        } else {
            option += fmt.Sprintf("&limit=%d", *request.Limit)
        }
    }
    if request.Direct != "" {
        if option == "" {
            option += fmt.Sprintf("?direct=%s", request.Direct)
        } else {
            option += fmt.Sprintf("&direct=%s", request.Direct)
        }
    }

    url := ac.PUrlBuilder.BuildWithQuery(linearswap.GET_METHOD, "/v5/algo/order/opens", nil, option)

    getResp, getErr := reqbuilder.HttpGet(url)
    if getErr != nil {
        log.Error("http get error: %s", getErr)
        return
    }

    var result responsealgo.QueryOpenAlgoOrdersResponse
    jsonErr := json.Unmarshal([]byte(getResp), &result)
    if jsonErr != nil {
        log.Error("convert json to QueryOpenAlgoOrdersResponse error: %s", jsonErr)
        return
    }

    data <- result
}

// QueryAlgoOrderHistoryAsync 查询历史策略订单
func (ac *AlgoClient) QueryAlgoOrderHistoryAsync(data chan responsealgo.QueryAlgoOrderHistoryResponse, request algo.QueryAlgoOrderHistoryRequest) {
    // 参数校验
    if request.Type == "" {
        log.Error("type参数为必填项")
        return
    }

    // option
    option := ""
    if request.Type != "" {
        option += fmt.Sprintf("?type=%s", request.Type)
    }
    if request.ContractCode != "" {
        if option == "" {
            option += fmt.Sprintf("?contract_code=%s", request.ContractCode)
        } else {
            option += fmt.Sprintf("&contract_code=%s", request.ContractCode)
        }
    }
    if request.MarginMode != "" {
        if option == "" {
            option += fmt.Sprintf("?margin_mode=%s", request.MarginMode)
        } else {
            option += fmt.Sprintf("&margin_mode=%s", request.MarginMode)
        }
    }
    if len(request.States) > 0 {
        if option == "" {
            option += fmt.Sprintf("?states=%s", strings.Join(request.States, ","))
        } else {
            option += fmt.Sprintf("&states=%s", strings.Join(request.States, ","))
        }
    }
    if request.StartTime != nil {
        if option == "" {
            option += fmt.Sprintf("?start_time=%d", *request.StartTime)
        } else {
            option += fmt.Sprintf("&start_time=%d", *request.StartTime)
        }
    }
    if request.EndTime != nil {
        if option == "" {
            option += fmt.Sprintf("?end_time=%d", *request.EndTime)
        } else {
            option += fmt.Sprintf("&end_time=%d", *request.EndTime)
        }
    }
    if request.From != nil && *request.From > 0 {
        if option == "" {
            option += fmt.Sprintf("?from=%d", *request.From)
        } else {
            option += fmt.Sprintf("&from=%d", *request.From)
        }
    }
    if request.Limit != nil && *request.Limit > 0 {
        if option == "" {
            option += fmt.Sprintf("?limit=%d", *request.Limit)
        } else {
            option += fmt.Sprintf("&limit=%d", *request.Limit)
        }
    }
    if request.Direct != "" {
        if option == "" {
            option += fmt.Sprintf("?direct=%s", request.Direct)
        } else {
            option += fmt.Sprintf("&direct=%s", request.Direct)
        }
    }

    url := ac.PUrlBuilder.BuildWithQuery(linearswap.GET_METHOD, "/v5/algo/order/history", nil, option)

    getResp, getErr := reqbuilder.HttpGet(url)
    if getErr != nil {
        log.Error("http get error: %s", getErr)
        return
    }

    var result responsealgo.QueryAlgoOrderHistoryResponse
    jsonErr := json.Unmarshal([]byte(getResp), &result)
    if jsonErr != nil {
        log.Error("convert json to QueryAlgoOrderHistoryResponse error: %s", jsonErr)
        return
    }

    data <- result
}
