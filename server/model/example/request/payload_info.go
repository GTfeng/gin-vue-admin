
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type PayloadInfoSearch struct{
    Name  *string `json:"name" form:"name" `
    VulType  *string `json:"vulType" form:"vulType" `
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
