package system

type AssetInfo struct {
	Id         int    `json:"id"`
	AssetId    string `json:"assetId" gorm:"column:assetId"`
	Name       string `json:"name"`
	SystemCode string `json:"systemCode" gorm:"column:systemCode"'`
	AssetType  string `json:"assetType" gorm:"column:assetType"`
	Subtype    string `json:"subtype" gorm:"column:subType"`
	Ip         string `json:"ip" gorm:"column:ip"`
}

func (AssetInfo) TableName() string {
	return "asset_info"
}
