package restful

import (
	"encoding/json"
	"fmt"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/response/common"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/log"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/reqbuilder"
)

type CommonClient struct {
	PUrlBuilder *reqbuilder.PrivateUrlBuilder
}

func (ac *CommonClient) Init(accessKey string, secretKey string, host string, sign string) *CommonClient {
	if host == "" {
		host = linearswap.LINEAR_SWAP_DEFAULT_HOST
	}
	ac.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host, sign)
	return ac
}

// GetSwapUnifiedAccountTypeAsync Account type query
func (ac *CommonClient) GetSwapUnifiedAccountTypeAsync(data chan common.GetSwapUnifiedAccountTypeResponse) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.GET_METHOD, "/linear-swap-api/v3/swap_unified_account_type", nil)

	// content is nil
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.GetSwapUnifiedAccountTypeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSwapUnifiedAccountTypeResponse error: %s", jsonErr)
	}
	data <- result
}

// SwapSwitchAccountTypeAsync Account Type Change
func (ac *CommonClient) SwapSwitchAccountTypeAsync(data chan common.SwapSwitchAccountTypeResponse, accountType int) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v3/swap_switch_account_type", nil)

	// content
	content := ""
	if accountType != 0 {
		content += fmt.Sprintf(",\"account_type\": \"%d\"", accountType)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := common.SwapSwitchAccountTypeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwapSwitchAccountTypeResponse error: %s", jsonErr)
	}
	data <- result
}
