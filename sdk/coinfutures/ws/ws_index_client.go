package ws

import (
	"encoding/json"
	"fmt"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinfutures"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinfutures/ws/response/index"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/wsbase"
	"reflect"
)

type WSIndexClient struct {
	WebSocketOp
}

func (wsIx *WSIndexClient) Init(host string, sign string) *WSIndexClient {
	if host == "" {
		host = coinfutures.COIN_FUTURES_DEFAULT_HOST
	}
	wsIx.open("/ws_index", host, "", "", true, sign)
	return wsIx
}

// -------------------------------------------------------------
// index kline start
type OnSubIndexKLineResponse func(*index.SubIndexKLineResponse)
type OnReqIndexKLineResponse func(*index.ReqIndexKLineResponse)

func (wsIx *WSIndexClient) SubIndexKLine(symbol string, period string, callbackFun OnSubIndexKLineResponse, id string) {
	if id == "" {
		id = coinfutures.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.index.%s", symbol, period)
	subData := wsbase.WSSubData{Sub: ch, Id: id}
	jdata, _ := json.Marshal(subData)

	wsIx.sub(jdata, ch, callbackFun, reflect.TypeOf(index.SubIndexKLineResponse{}))
}

func (wsIx *WSIndexClient) ReqIndexKLine(symbol string, period string, callbackFun OnReqIndexKLineResponse,
	from int64, to int64, id string) {
	if id == "" {
		id = coinfutures.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.index.%s", symbol, period)
	reqData := wsbase.WSReqData{Req: ch, Id: id, From: from, To: to}
	jdata, _ := json.Marshal(reqData)

	wsIx.req(jdata, ch, callbackFun, reflect.TypeOf(index.ReqIndexKLineResponse{}))
}

// index kline end
//-------------------------------------------------------------

// -------------------------------------------------------------
// mark price kline start
type OnSubMarkPriceKLineResponse func(*index.SubIndexKLineResponse)
type OnReqMarkPriceKLineResponse func(*index.ReqIndexKLineResponse)

func (wsIx *WSIndexClient) SubMarkPriceKLine(symbol string, period string, callbackFun OnSubMarkPriceKLineResponse, id string) {
	if id == "" {
		id = coinfutures.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.mark_price.%s", symbol, period)
	subData := wsbase.WSSubData{Sub: ch, Id: id}
	jdata, _ := json.Marshal(subData)

	wsIx.sub(jdata, ch, callbackFun, reflect.TypeOf(index.SubIndexKLineResponse{}))
}

func (wsIx *WSIndexClient) ReqMarkPriceKLine(symbol string, period string, callbackFun OnReqMarkPriceKLineResponse,
	from int64, to int64, id string) {
	if id == "" {
		id = coinfutures.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.mark_price.%s", symbol, period)
	reqData := wsbase.WSReqData{Req: ch, Id: id, From: from, To: to}
	jdata, _ := json.Marshal(reqData)

	wsIx.req(jdata, ch, callbackFun, reflect.TypeOf(index.ReqIndexKLineResponse{}))
}

// mark price kline end
//-------------------------------------------------------------

// -------------------------------------------------------------
// basis start
type OnSubBasisResponse func(*index.SubBasiesResponse)
type OnReqBasisResponse func(*index.ReqBasisResponse)

func (wsIx *WSIndexClient) SubBasis(symbol string, period string, callbackFun OnSubBasisResponse, basisPriceType string, id string) {
	if basisPriceType == "" {
		basisPriceType = "open"
	}
	if id == "" {
		id = coinfutures.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.basis.%s.%s", symbol, period, basisPriceType)
	subData := wsbase.WSSubData{Sub: ch, Id: id}
	jdata, _ := json.Marshal(subData)

	wsIx.sub(jdata, ch, callbackFun, reflect.TypeOf(index.SubBasiesResponse{}))
}

func (wsIx *WSIndexClient) ReqBasis(symbol string, period string, callbackFun OnReqBasisResponse, from int64, to int64,
	basisPriceType string, id string) {
	if basisPriceType == "" {
		basisPriceType = "open"
	}
	if id == "" {
		id = coinfutures.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.basis.%s.%s", symbol, period, basisPriceType)
	reqData := wsbase.WSReqData{Req: ch, Id: id, From: from, To: to}
	jdata, _ := json.Marshal(reqData)

	wsIx.req(jdata, ch, callbackFun, reflect.TypeOf(index.ReqBasisResponse{}))
}
