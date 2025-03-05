package penetration

import "time"

type TaskInfo struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	AssetId    string     `json:"assetId"`
	SystemCode string     `json:"systemCode"`
	AssetType  string     `json:"assetType"`
	SubType    string     `json:"subType"`
	Ip         string     `json:"ip"`
	UpdateTime *time.Time `json:"updateTime"`
}
