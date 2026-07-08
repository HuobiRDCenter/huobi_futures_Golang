package account

type InviteeRebateBatcherRebateDetailResponse struct {
	Status string                           `json:"status,omitempty"`
	Data   InviteeRebateBatcherRebateDetail `json:"data,omitempty"`
	Ts     int64                            `json:"ts,omitempty"`
}

type InviteeRebateBatcherRebateDetail struct {
	InviteeUID                       int    `json:"invitee_uid,omitempty"`
	InviteeType                      string `json:"invitee_type,omitempty"`
	InviteeRebateRateSpotM2          string `json:"invitee_rebate_rate_spot_m2,omitempty"`
	InviteeRebateRateContractM2      string `json:"invitee_rebate_rate_contract_m2,omitempty"`
	InviteeRebateRatePartnerSpot     string `json:"invitee_rebate_rate_partner_spot,omitempty"`
	InviteeRebateRatePartnerContract string `json:"invitee_rebate_rate_partner_contract,omitempty"`
	JoinTimeM2                       string `json:"join_time_m2,omitempty"`
	JoinTimePartner                  string `json:"join_time_partner,omitempty"`
	InviteeTotalCommissionUsdt       string `json:"Invitee_total_commission_usdt,omitempty"`
	InviteeTotalCommissionTrx        string `json:"Invitee_total_commission_trx,omitempty"`
	InviteeTotalCommissionHtx        string `json:"Invitee_total_commission_htx,omitempty"`
	PartnerTotalCommissionUsdt       string `json:"partner_total_commission_usdt,omitempty"`
	PartnerTotalCommissionTrx        string `json:"partner_total_commission_trx,omitempty"`
	PartnerTotalCommissionHtx        string `json:"partner_total_commission_htx,omitempty"`
}
