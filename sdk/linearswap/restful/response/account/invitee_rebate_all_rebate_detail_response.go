package account

type InviteeRebateAllRebateDetailResponse struct {
	Status     string                       `json:"status,omitempty"`
	InviteeUID int                          `json:"invitee_uid,omitempty"`
	Data       InviteeRebateAllRebateDetail `json:"data,omitempty"`
	Ts         int64                        `json:"ts,omitempty"`
	NextID     string                       `json:"nextId,omitempty"`
}

type InviteeRebateAllRebateDetail struct {
	InviteeType                      string `json:"invitee_type,omitempty"`
	InviteeRebateRateSpotM2          string `json:"invitee_rebate_rate_spot_m2,omitempty"`
	InviteeRebateRateContractM2      string `json:"invitee_rebate_rate_contract_m2,omitempty"`
	InviteeRebateRatePartnerSpot     string `json:"invitee_rebate_rate_partner_spot,omitempty"`
	InviteeRebateRatePartnerContract string `json:"invitee_rebate_rate_partner_contract,omitempty"`
	JoinTimeM2                       string `json:"join_time_m2,omitempty"`
	JoinTimePartner                  string `json:"join_time_partner,omitempty"`
	InviteeTotalCommissionUSDT       string `json:"Invitee_total_commission_usdt,omitempty"`
	InviteeTotalCommissionTRX        string `json:"Invitee_total_commission_trx,omitempty"`
	InviteeTotalCommissionHTX        string `json:"Invitee_total_commission_htx,omitempty"`
	PartnerTotalCommissionUSDT       string `json:"partner_total_commission_usdt,omitempty"`
	PartnerTotalCommissionTRX        string `json:"partner_total_commission_trx,omitempty"`
	PartnerTotalCommissionHTX        string `json:"partner_total_commission_htx,omitempty"`
}
