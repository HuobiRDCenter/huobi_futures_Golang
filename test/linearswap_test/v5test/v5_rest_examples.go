package v5test

import (
	"fmt"

	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/response/account"
	"github.com/HuobiRDCenter/huobi_futures_Golang/sdk/linearswap/restful/response/market"
)

// V5 实测统一用 SPX500-USDT：指数合约单张面值低，远价限价单（1000，市价≈7600）
// 挂上后立刻撤，0 成交 0 损失。口径与 C++ test/v5_restful_*_test.hpp 一致。
const (
	v5Contract = "SPX500-USDT"
	v5FarPrice = "1000"
)

// 每个用例独立建 client，与仓库既有测试写法一致
func acctClient() *restful.AccountClient {
	return new(restful.AccountClient).Init(accessKey(), secretKey(), host())
}

func mktClient() *restful.MarketClient {
	return new(restful.MarketClient).Init(host())
}

// -------------------------------------------------------------
// V5 account（私有，7 个）

func restAccountBalance() {
	name := "GET  /v5/account/balance"
	ch := make(chan account.GetAccountBalanceResponse, 1)
	go acctClient().GetAccountBalanceAsync(ch)
	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
	if r.Code == 200 {
		if r.Data.Equity == "" {
			bad(name, "code=200 但 data.equity 为空，应含账户权益")
			return
		}
		if len(r.Data.Details) == 0 {
			bad(name, "data.details 为空数组")
			return
		}
		// 时间戳字段线上是数字（与推送帧里的字符串形态不同），类型写错时会静默留 0
		if r.Data.Details[0].CreatedTime == 0 {
			bad(name, "data.details[0].created_time=0，该字段可能又被声明成了字符串（线上是数字）")
		}
	}
}

// 幂等：先查当前资产模式，再设回同一个值，避免改变账户真实配置
func restSetAssetMode() {
	name := "POST /v5/account/asset_mode"
	getCh := make(chan account.AssetModeResponse, 1)
	go acctClient().GetAssetModeAsync(getCh)
	cur := <-getCh
	if cur.Code != 200 {
		skip(name, fmt.Sprintf("前置查询资产模式失败 code=%d，不做幂等写", cur.Code))
		return
	}

	ch := make(chan account.AssetModeResponse, 1)
	go acctClient().SetAssetModeAsync(ch, cur.Data.AssetMode)
	r := <-ch
	// 200=设置成功；6010=已是该模式；1421=有持仓/挂单时不允许改。三者都说明接口可达且签名正确
	if r.Code == 200 || r.Code == 6010 || r.Code == 1421 {
		ok(name, int64(r.Code), r.Message)
		return
	}
	check(name, int64(r.Code), r.Ts, r.Message)
}

func restGetAssetMode() {
	name := "GET  /v5/account/asset_mode"
	ch := make(chan account.AssetModeResponse, 1)
	go acctClient().GetAssetModeAsync(ch)
	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
}

func restSetFeeDeductionCurrency() {
	name := "POST /v5/account/fee_deduction_currency"
	ch := make(chan account.FeeDeductionCurrencyResponse, 1)
	go acctClient().FeeDeductionCurrency(ch, 1, "htx")
	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
	if r.Code == 200 {
		if r.Data.FeeOption != 1 || r.Data.DeductionCurrency != "htx" {
			bad(name, "回显不符：fee_option=%d deduction_currency=%s，期望 1/htx",
				r.Data.FeeOption, r.Data.DeductionCurrency)
		}
	}
}

func restGetFeeDeductionCurrency() {
	name := "GET  /v5/account/fee_deduction_currency"
	ch := make(chan account.FeeDeductionCurrencyResponse, 1)
	go acctClient().GetFeeDeductionCurrencyAsync(ch)
	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
}

// bills 是带业务参数的 GET：参数必须进签名，否则 401（本仓库 BuildWithQuery 的存在理由）
func restAccountBills() {
	name := "GET  /v5/account/bills"
	ch := make(chan account.AccountBillsResponse, 1)
	go acctClient().GetAccountBillsAsync(ch, v5Contract, "isolated", "", "", "", 0, 0, "")

	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
}

// multi_assets_margin 不在 V5 官方清单（43 个 REST）内：C++ const_val.hpp 自己标注为
// "6 + 1 multi_assets_margin"，线上实测返回 404 Resources do not exist。
// 方法保留以对齐 Java/C++ 代码，实测按 SKIP 处理，不计入失败。
func restMultiAssetsMargin() {
	skip("POST /v5/account/multi_assets_margin", "不在 V5 清单内（C++ 标注 6+1），线上 404，接口未开放")
}

// -------------------------------------------------------------
// V5 market（公开，12 个，无需签名）

func restMarketRiskLimit() {
	name := "GET  /v5/market/risk/limit"
	ch := make(chan market.GetMarketRiskLimitResponse, 1)
	go mktClient().GetMarketRiskLimitAsync(ch, v5Contract, "", "", "")

	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
	if r.Code == 200 && len(r.Data) == 0 {
		bad(name, "code=200 但 data 为空数组")
	}
}

func restMarketAssetsDeductionCurrency() {
	name := "GET  /v5/market/assets_deduction_currency"
	ch := make(chan market.GetMarketAssetsDeductionCurrencyResponse, 1)
	go mktClient().GetMarketAssetsDeductionCurrencyAsync(ch)

	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
}

func restMarketMultiAssetsMarginList() {
	name := "GET  /v5/market/multi_assets_margin"
	ch := make(chan market.GetMarketMultiAssetsMarginListResponse, 1)
	go mktClient().GetMarketMultiAssetsMarginListAsync(ch)

	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
	if r.Code == 200 && len(r.Data.MultiAssets) == 0 {
		bad(name, "code=200 但 data.multi_assets 为空数组")
	}
}

func restMarketFundingRate() {
	name := "GET  /v5/market/funding_rate"
	ch := make(chan market.MarketFundingRateResponse, 1)
	go mktClient().GetMarketFundingRateAsync(ch, "BTC-USDT")
	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
	if r.Code == 200 {
		if len(r.Data) == 0 {
			bad(name, "code=200 但 data 为空数组")
		} else if r.Data[0].FundingRate == "" {
			bad(name, "data[0].funding_rate 为空")
		}
	}
}

func restMarketFundingRateHistory() {
	name := "GET  /v5/market/funding_rate_history"
	ch := make(chan market.MarketFundingRateHistoryResponse, 1)
	go mktClient().GetFundingRateHistoryAsync(ch, v5Contract, nil, nil, nil, nil, "")

	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
}

func restMarketOpenInterest() {
	name := "GET  /v5/market/open_interest"
	ch := make(chan market.MarketOpenInterestResponse, 1)
	go mktClient().GetMarketOpenInterestAsync(ch, "BTC-USDT")

	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
	if r.Code == 200 && r.Data.ContractCode == "" {
		bad(name, "data.contract_code 为空")
	}
}

func restMarketPriceLimit() {
	name := "GET  /v5/market/price_limit"
	ch := make(chan market.MarketPriceLimitResponse, 1)
	go mktClient().GetMarketPriceLimitAsync(ch, "BTC-USDT")
	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
	if r.Code == 200 {
		if len(r.Data) == 0 {
			bad(name, "code=200 但 data 为空数组")
		} else if r.Data[0].HighLimit == "" {
			bad(name, "data[0].high_limit 为空")
		}
	}
}

// 强平订单平时为空数组，只验接口可达
func restMarketLiquidationOrders() {
	name := "GET  /v5/market/liquidation_orders"
	ch := make(chan market.MarketLiquidationOrdersResponse, 1)
	go mktClient().GetMarketLiquidationOrdersAsync(ch, "BTC-USDT", "", nil, nil, "", nil, nil)

	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
}

func restMarketSettlementHistory() {
	name := "GET  /v5/market/settlement_history"
	ch := make(chan market.MarketSettlementHistoryResponse, 1)
	go mktClient().GetSettlementHistoryAsync(ch, v5Contract, nil, nil, "", "", nil)

	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
}

func restMarketEliteAccountRatio() {
	name := "GET  /v5/market/elite_account_ratio"
	ch := make(chan market.MarketEliteAccountRatioResponse, 1)
	go mktClient().GetMarketEliteAccountRatioAsync(ch, "BTC-USDT", "5min")

	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
}

func restMarketElitePositionRatio() {
	name := "GET  /v5/market/elite_position_ratio"
	ch := make(chan market.MarketElitePositionRatioResponse, 1)
	go mktClient().GetMarketElitePositionRatioAsync(ch, "BTC-USDT", "5min")

	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
}

func restMarketEstimatedSettlementPrice() {
	name := "GET  /v5/market/estimated_settlement_price"
	ch := make(chan market.MarketEstimatedSettlementPriceResponse, 1)
	go mktClient().GetMarketEstimatedSettlementPriceAsync(ch, "BTC-USDT")

	r := <-ch
	check(name, int64(r.Code), r.Ts, r.Message)
}

// RunRestAccountMarket 跑 V5 account（7）与 market（12）
func RunRestAccountMarket() {
	Section("V5 REST · account（私有，7）")
	restAccountBalance()
	restSetAssetMode()
	restGetAssetMode()
	restSetFeeDeductionCurrency()
	restGetFeeDeductionCurrency()
	restAccountBills()
	restMultiAssetsMargin()

	Section("V5 REST · market（公开，12）")
	restMarketRiskLimit()
	restMarketAssetsDeductionCurrency()
	restMarketMultiAssetsMarginList()
	restMarketFundingRate()
	restMarketFundingRateHistory()
	restMarketOpenInterest()
	restMarketPriceLimit()
	restMarketLiquidationOrders()
	restMarketSettlementHistory()
	restMarketEliteAccountRatio()
	restMarketElitePositionRatio()
	restMarketEstimatedSettlementPrice()
}
