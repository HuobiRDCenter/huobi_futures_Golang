package copytrading

type CopytradingTraderQueryFollowersResponse struct {
	Tid  string                          `json:"tid"`
	Data CopytradingTraderQueryFollowers `json:"data"`
	Code int64                           `json:"code"`
}

type CopytradingTraderQueryFollowers struct {
	Follower []Follower `json:"follower"`
	QueryID  string     `json:"query_id"`
}

type Follower struct {
	FollowerHeadPic     string `json:"follower_head_pic"`
	FollowerName        string `json:"follower_name"`
	FollowerUID         string `json:"follower_uid"`
	FollowerTime        int64  `json:"follower_time"`
	FollowerProfitAmont string `json:"follower_profit_amont"`
	FollowerTradeAmount string `json:"follower_trade_amount"`
}
