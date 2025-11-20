package ws

import (
	"encoding/json"
	"fmt"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinfutures"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/coinfutures/ws/response/notify"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/wsbase"
	"reflect"
	"strings"
)

type WSNotifyClient struct {
	WebSocketOp
}

func (wsNf *WSNotifyClient) Init(accessKey string, secretKey string, host string, sign string) *WSNotifyClient {
	if host == "" {
		host = coinfutures.COIN_FUTURES_DEFAULT_HOST
	}
	wsNf.open("/notification", host, accessKey, secretKey, true, sign)
	return wsNf
}

// -------------------------------------------------------------
// orders start

type OnSubOrdersResponse func(*notify.SubOrdersResponse)

func (wsNf *WSNotifyClient) IsolatedSubOrders(symbol string, callbackFun OnSubOrdersResponse, cid string) {
	if cid == "" {
		cid = coinfutures.DEFAULT_CID
	}
	ch := fmt.Sprintf("orders.%s", symbol)
	opData := wsbase.WSOpData{Op: "sub", Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(notify.SubOrdersResponse{}))
}

func (wsNf *WSNotifyClient) IsolatedUnsubOrders(symbol string, cid string) {
	if cid == "" {
		cid = coinfutures.DEFAULT_CID
	}

	ch := fmt.Sprintf("orders.%s", symbol)
	opData := wsbase.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

// orders end
//-------------------------------------------------------------

// -------------------------------------------------------------
// accounts start

type OnSubAccountsResponse func(*notify.SubAccountsResponse)

func (wsNf *WSNotifyClient) IsolatedSubAcounts(symbol string, callbackFun OnSubAccountsResponse, cid string) {
	if cid == "" {
		cid = coinfutures.DEFAULT_CID
	}

	ch := fmt.Sprintf("accounts.%s", symbol)
	opData := wsbase.WSOpData{Op: "sub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(notify.SubAccountsResponse{}))
}

func (wsNf *WSNotifyClient) IsolatedUnsubAccounts(symbol string, cid string) {
	if cid == "" {
		cid = coinfutures.DEFAULT_CID
	}

	ch := fmt.Sprintf("accounts.%s", symbol)
	opData := wsbase.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

// accounts end
//-------------------------------------------------------------

// -------------------------------------------------------------
// positions start

type OnSubPositionsResponse func(*notify.SubPositionsResponse)

func (wsNf *WSNotifyClient) IsolatedSubPositions(symbol string, callbackFun OnSubPositionsResponse, cid string) {
	if cid == "" {
		cid = coinfutures.DEFAULT_CID
	}

	ch := fmt.Sprintf("positions.%s", symbol)
	opData := wsbase.WSOpData{Op: "sub", Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(notify.SubPositionsResponse{}))
}

func (wsNf *WSNotifyClient) IsolatdUnsubPositions(symbol string, cid string) {
	if cid == "" {
		cid = coinfutures.DEFAULT_CID
	}

	ch := fmt.Sprintf("positions.%s", symbol)
	opData := wsbase.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

// positions end
//-------------------------------------------------------------

// -------------------------------------------------------------
// match orders start

type OnSubMatchOrdersResponse func(*notify.SubOrdersResponse)

func (wsNf *WSNotifyClient) IsolatedSubMatchOrders(symbol string, callbackFun OnSubMatchOrdersResponse, cid string) {
	if cid == "" {
		cid = coinfutures.DEFAULT_CID
	}

	symbol = strings.ToLower(symbol)
	ch := fmt.Sprintf("matchOrders.%s", symbol)
	opData := wsbase.WSOpData{Op: "sub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(notify.SubOrdersResponse{}))
}

func (wsNf *WSNotifyClient) IsolatedUnsubMathOrders(symbol string, cid string) {
	if cid == "" {
		cid = coinfutures.DEFAULT_CID
	}

	symbol = strings.ToLower(symbol)
	ch := fmt.Sprintf("matchOrders.%s", symbol)
	opData := wsbase.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

// match orders end
//-------------------------------------------------------------

// -------------------------------------------------------------
// liquidation orders start

type OnSubLiquidationOrdersResponse func(*notify.SubLiquidationOrdersResponse)

func (wsNf *WSNotifyClient) SubLiquidationOrders(symbol string, callbackFun OnSubLiquidationOrdersResponse, cid string) {
	if cid == "" {
		cid = coinfutures.DEFAULT_CID
	}

	ch := fmt.Sprintf("public.%s.liquidation_orders", symbol)
	opData := wsbase.WSOpData{Op: "sub", Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(notify.SubLiquidationOrdersResponse{}))
}

func (wsNf *WSNotifyClient) UnsubLiquidationOrders(symbol string, cid string) {
	if cid == "" {
		cid = coinfutures.DEFAULT_CID
	}

	ch := fmt.Sprintf("public.%s.liquidation_orders", symbol)
	opData := wsbase.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

// liquidation orders end
//-------------------------------------------------------------

// -------------------------------------------------------------
// contract info start

type OnSubContractInfoResponse func(*notify.SubContractInfoResponse)

func (wsNf *WSNotifyClient) SubContractInfo(symbol string, callbackFun OnSubContractInfoResponse, cid string) {
	if cid == "" {
		cid = coinfutures.DEFAULT_CID
	}

	ch := fmt.Sprintf("public.%s.contract_info", symbol)
	opData := wsbase.WSOpData{Op: "sub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(notify.SubContractInfoResponse{}))
}

func (wsNf *WSNotifyClient) UnsubContractInfo(symbol string, cid string) {
	if cid == "" {
		cid = coinfutures.DEFAULT_CID
	}

	ch := fmt.Sprintf("public.%s.contract_info", symbol)
	opData := wsbase.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

// contract info end
//-------------------------------------------------------------

// -------------------------------------------------------------
// trigger order start

type OnSubTriggerOrderResponse func(*notify.SubTriggerOrderResponse)

func (wsNf *WSNotifyClient) IsolatedSubTriggerOrder(symbol string, callbackFun OnSubTriggerOrderResponse, cid string) {
	if cid == "" {
		cid = coinfutures.DEFAULT_CID
	}

	ch := fmt.Sprintf("trigger_order.%s", symbol)
	opData := wsbase.WSOpData{Op: "sub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(notify.SubTriggerOrderResponse{}))
}

func (wsNf *WSNotifyClient) IsolatedUnsubTriggerOrder(symbol string, cid string) {
	if cid == "" {
		cid = coinfutures.DEFAULT_CID
	}

	ch := fmt.Sprintf("trigger_order.%s", symbol)
	opData := wsbase.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

// trigger info end
//-------------------------------------------------------------

type OnSubContractElementsResponse func(*notify.SubContractElementsResponse)

func (wsNf *WSNotifyClient) SubContractElements(symbol string, callbackFun OnSubContractElementsResponse, cid string) {
	if cid == "" {
		cid = coinfutures.DEFAULT_CID
	}

	ch := fmt.Sprintf("public.%s.contract_elements", symbol)
	opData := wsbase.WSOpData{Op: "sub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(notify.SubContractElementsResponse{}))
}

func (wsNf *WSNotifyClient) UnSubContractElements(symbol string, cid string) {
	if cid == "" {
		cid = coinfutures.DEFAULT_CID
	}

	ch := fmt.Sprintf("public.%s.contract_elements", symbol)
	opData := wsbase.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}
