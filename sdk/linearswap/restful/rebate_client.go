package restful

import (
	"encoding/json"
	"fmt"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/reqbuilder"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/response/rebate"
	"strconv"
)

type RebateClient struct {
	PUrlBuilder *reqbuilder.PrivateUrlBuilder
}

func (rc *RebateClient) Init(accessKey string, secretKey string, host string) *RebateClient {
	if host == "" {
		host = linearswap.LINEAR_SWAP_DEFAULT_HOST
	}
	rc.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return rc
}

// GetReferralsAsync 查询邀请用户列表（异步）
func (rc *RebateClient) GetReferralsAsync(data chan rebate.InviteeRebateReferralsResponse,
	userId int64,
	inviteeUidList string,
	referralCode string,
	startTime string,
	endTime string,
	direct string,
	fromId string,
	limit int64) {
	
	// 构建URL和参数
	request := new(reqbuilder.GetRequest).Init()
	request.AddParam("userId", strconv.FormatInt(userId, 10))
	
	if inviteeUidList != "" {
		request.AddParam("inviteeUidList", inviteeUidList)
	}
	if referralCode != "" {
		request.AddParam("referralCode", referralCode)
	}
	if startTime != "" {
		request.AddParam("startTime", startTime)
	}
	if endTime != "" {
		request.AddParam("endTime", endTime)
	}
	if direct != "" {
		request.AddParam("direct", direct)
	}
	if fromId != "" {
		request.AddParam("fromId", fromId)
	}
	if limit != 0 {
		request.AddParam("limit", strconv.FormatInt(limit, 10))
	}
	
	// url
	url := rc.PUrlBuilder.Build(linearswap.GET_METHOD, "/v2/invitee/rebate/referrals", request)
	
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	
	result := rebate.InviteeRebateReferralsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json error: %s", jsonErr)
	}
	data <- result
}