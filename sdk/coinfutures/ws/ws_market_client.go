package ws

import (
	"encoding/json"
	"fmt"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinfutures"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinfutures/ws/response/market"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/wsbase"
	"reflect"
)

type WSMarketClient struct {
	WebSocketOp
}

func (wsMk *WSMarketClient) Init(host string, sign string) *WSMarketClient {
	if host == "" {
		host = coinfutures.COIN_FUTURES_DEFAULT_HOST
	}
	wsMk.open("/ws", host, "", "", true, sign)
	return wsMk
}

// -------------------------------------------------------------
// kline start

type OnSubKLineResponse func(*market.SubKLineResponse)
type OnReqKLineResponse func(*market.ReqKLineResponse)

func (wsMk *WSMarketClient) SubKLine(symbol string, period string, callbackFun OnSubKLineResponse, id string) {
	if id == "" {
		id = coinfutures.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.kline.%s", symbol, period)
	subData := wsbase.WSSubData{Sub: ch, Id: id}
	jdata, _ := json.Marshal(subData)
	wsMk.sub(jdata, ch, callbackFun, reflect.TypeOf(market.SubKLineResponse{}))
}

func (wsMk *WSMarketClient) ReqKLine(symbol string, period string, callbackFun OnReqKLineResponse, from int64, to int64, id string) {
	if id == "" {
		id = coinfutures.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.kline.%s", symbol, period)
	subData := wsbase.WSReqData{Req: ch, From: from, To: to, Id: id}
	jdata, _ := json.Marshal(subData)
	wsMk.req(jdata, ch, callbackFun, reflect.TypeOf(market.ReqKLineResponse{}))
}

// kline end
//-------------------------------------------------------------

//-------------------------------------------------------------
// depth start

type OnSubDepthResponse func(*market.SubDepthResponse)

func (wsMk *WSMarketClient) SubDepth(symbol string, fcType string, callbackFun OnSubDepthResponse, id string) {
	if id == "" {
		id = coinfutures.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.depth.%s", symbol, fcType)
	subData := wsbase.WSSubData{Sub: ch, Id: id}
	jdata, _ := json.Marshal(subData)
	wsMk.sub(jdata, ch, callbackFun, reflect.TypeOf(market.SubDepthResponse{}))
}

// depth end
//-------------------------------------------------------------

//-------------------------------------------------------------
// incrementa depth start

func (wsMk *WSMarketClient) SubIncrementalDepth(symbol string, size string, callbackFun OnSubDepthResponse, id string) {
	if id == "" {
		id = coinfutures.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.depth.size_%s.high_freq", symbol, size)
	subData := wsbase.WSSubData{Sub: ch, Id: id}
	jdata, _ := json.Marshal(subData)
	wsMk.sub(jdata, ch, callbackFun, reflect.TypeOf(market.SubDepthResponse{}))
}

// incrementa depth end
//-------------------------------------------------------------

// -------------------------------------------------------------
// detail start
type OnSubDetailResponse func(*market.SubKLineResponse)

func (wsMk *WSMarketClient) SubDetail(symbol string, callbackFun OnSubDetailResponse, id string) {
	if id == "" {
		id = coinfutures.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.detail", symbol)
	subData := wsbase.WSSubData{Sub: ch, Id: id}
	jdata, _ := json.Marshal(subData)
	wsMk.sub(jdata, ch, callbackFun, reflect.TypeOf(market.SubKLineResponse{}))
}

// detail end
//-------------------------------------------------------------

//-------------------------------------------------------------
// bbo start

type OnSubBBOResponse func(*market.SubBBOResponse)

func (wsMk *WSMarketClient) SubBBO(symbol string, callbackFun OnSubBBOResponse, id string) {
	if id == "" {
		id = coinfutures.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.bbo", symbol)
	subData := wsbase.WSSubData{Sub: ch, Id: id}
	jdata, _ := json.Marshal(subData)
	wsMk.sub(jdata, ch, callbackFun, reflect.TypeOf(market.SubBBOResponse{}))
}

// bbo end
//-------------------------------------------------------------

//-------------------------------------------------------------
// trade detail start

type OnSubTradeDetailResponse func(*market.SubTradeDetailResponse)
type OnReqTradeDetailResponse func(*market.ReqTradeDetailResponse)

func (wsMk *WSMarketClient) SubTradeDetail(symbol string, callbackFun OnSubTradeDetailResponse, id string) {
	if id == "" {
		id = coinfutures.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.trade.detail", symbol)
	subData := wsbase.WSSubData{Sub: ch, Id: id}
	jdata, _ := json.Marshal(subData)
	wsMk.sub(jdata, ch, callbackFun, reflect.TypeOf(market.SubTradeDetailResponse{}))
}

func (wsMk *WSMarketClient) ReqTradeDetail(symbol string, callbackFun OnReqTradeDetailResponse, id string) {
	if id == "" {
		id = coinfutures.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.trade.detail", symbol)
	subData := wsbase.WSReqData{Req: ch, Id: id}
	jdata, _ := json.Marshal(subData)
	wsMk.req(jdata, ch, callbackFun, reflect.TypeOf(market.ReqTradeDetailResponse{}))
}
