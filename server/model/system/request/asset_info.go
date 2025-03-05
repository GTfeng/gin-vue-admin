package request

import common "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

type GetAssetList struct {
	common.PageInfo
	Name       string `json:"name" form:"name"`
	SystemCode string `json:"systemCode" form:"systemCode"`
}
