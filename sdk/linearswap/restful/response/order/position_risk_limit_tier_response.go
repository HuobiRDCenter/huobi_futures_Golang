package order

type PositionRiskLimitTierResponse struct {
    Code    int64                    `json:"code"`
    Data    []PositionRiskLimitTierData `json:"data"`
    Message string                   `json:"message"`
    Ts      int64                    `json:"ts"`
}

type PositionRiskLimitTierData struct {
    ContractCode           string `json:"contract_code"`
    MarginMode             string `json:"margin_mode"`
    Tier                   string `json:"tier"`
    MaxLever               string `json:"max_lever"`
    MaintenanceMarginRate  string `json:"maintenance_margin_rate"`
    MaxVolume              string `json:"max_volume"`
    MinVolume              string `json:"min_volume"`
    VolumeUnit             string `json:"volume_unit"`
}