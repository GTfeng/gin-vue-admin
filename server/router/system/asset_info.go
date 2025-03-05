package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AssetRouter struct{}

func (s *AssetRouter) InitAssetRouter(Router *gin.RouterGroup) {
	assetRouter := Router.Group("asset").Use(middleware.OperationRecord())
	assetRouterWithoutRecord := Router.Group("asset")
	{
		//assetRouter.GET("getTaskList", baseApi.GetUserList)
		assetRouter.GET("getTaskList", assetApi.GetAssetList)
	}
	assetRouterWithoutRecord.GET("getAssetList", baseApi.GetUserList) // 分页获取用户列表
}
