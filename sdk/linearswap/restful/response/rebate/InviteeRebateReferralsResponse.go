package rebate

// 查询邀请用户列表响应
type InviteeRebateReferralsResponse struct {
	Code    int                       `json:"code"`
	Message string                    `json:"message"`
	Ts      int64                     `json:"ts"`
	Data    InviteeRebateReferralsData `json:"data"`
}

type InviteeRebateReferralsData struct {
	NextId          int64              `json:"nextId"`
	RebateDetailList []ReferralDetailVO `json:"rebateDetailList"`
}

type ReferralDetailVO struct {
	Id                              int    `json:"id"`
	InviteUid                       int    `json:"inviteUid"`
	Remark                          string `json:"remark"`
	ReferralCode                    string `json:"referralCode"`
	InviteeType                     int    `json:"inviteeType"`                           // 用户类型. 0: 直客(m2)，1: 我的合伙人(kn)，2：直客 + 我的合伙人
	InviteeRebateRateSpotM2         string `json:"inviteeRebateRateSpotM2"`              // m2的现货返佣比例
	InviteeRebateRateContractM2     string `json:"inviteeRebateRateContractM2"`          // m2的合约返佣比例
	InviteeRebateRatePartnerSpot    string `json:"inviteeRebateRatePartnerSpot"`         // 合伙人现货返佣比例
	InviteeRebateRatePartnerContract string `json:"inviteeRebateRatePartnerContract"`    // 合伙人合约返佣比例
	JoinTimeM2                      int64  `json:"joinTimeM2"`                           // 成为m2的时间
	JoinTimePartner                 int64  `json:"joinTimePartner"`                      // 成为多级时间
	SpotTradingVolume               string `json:"spotTradingVolume"`
	FuturesTradingVolume            string `json:"futuresTradingVolume"`
}