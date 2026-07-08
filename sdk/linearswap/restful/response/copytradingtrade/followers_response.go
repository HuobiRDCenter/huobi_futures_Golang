package copytradingtrade

type FollowersResponse struct {
    Id                    string `json:"id"`
    InstType              string `json:"instType"`
    FollowerAvatarLink    string `json:"followerAvatarLink"`
    FollowerName          string `json:"followerName"`
    FollowerUid           string `json:"followerUid"`
    FollowTime            string `json:"followTime"`
    FollowerAssetAmt      string `json:"followerAssetAmt"`
    FollowerProfitSharingAmt string `json:"followerProfitSharingAmt"`
    FollowerTradeAmt      string `json:"followerTradeAmt"`
    TotalProfitSharingAmt string `json:"totalProfitSharingAmt"`
}