package account

type CrossGetSettlementRecordsResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		TotalPage         int `json:"total_page"`
		CurrentPage       int `json:"current_page"`
		TotalSize         int `json:"total_size"`
		SettlementRecords []struct {
			MarginMode string `json:"margin_mode"`

			MarginAccount string `json:"margin_account"`

			MarginBalanceInit float32 `json:"margin_balance_init"`

			MarginBalance float32 `json:"margin_balance"`

			SettlementProfitReal float32 `json:"settlement_profit_real"`

			SettlementTime int64 `json:"settlement_time"`

			Clawback float32 `json:"clawback"`

			FundingFee float32 `json:"funding_fee"`

			OffsetProfitloss float32 `json:"offset_profitloss"`

			Fee float32 `json:"fee"`

			FeeAsset string `json:"fee_asset"`

			ContractDetail []struct {
				Symbol string `json:"symbol"`

				ContractCode string `json:"contract_code"`

				OffsetProfitloss float32 `json:"offset_profitloss"`

				Fee float32 `json:"fee"`

				FeeAsset string `json:"fee_asset"`

				Pair string `json:"pair"`

				Positions []struct {
					Symbol string `json:"symbol"`

					ContractCode string `json:"contract_code"`

					Direction string `json:"direction"`

					Volume float32 `json:"volume"`

					CostOpen float32 `json:"cost_open"`

					CostHoldPre float32 `json:"cost_hold_pre"`

					CostHold float32 `json:"cost_hold"`

					SettlementProfitUnreal float32 `json:"settlement_profit_unreal"`

					SettlementPrice float32 `json:"settlement_price"`

					SettlementType string `json:"settlement_type"`

					Pair string `json:"pair"`
				} `json:"positions"`
			} `json:"contract_detail"`
		} `json:"settlement_records"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
