package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PayloadInfoRouter struct {}

// InitPayloadInfoRouter 初始化 payloadInfo表 路由信息
func (s *PayloadInfoRouter) InitPayloadInfoRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	payloadInfoRouter := Router.Group("payloadInfo").Use(middleware.OperationRecord())
	payloadInfoRouterWithoutRecord := Router.Group("payloadInfo")
	payloadInfoRouterWithoutAuth := PublicRouter.Group("payloadInfo")
	{
		payloadInfoRouter.POST("createPayloadInfo", payloadInfoApi.CreatePayloadInfo)   // 新建payloadInfo表
		payloadInfoRouter.DELETE("deletePayloadInfo", payloadInfoApi.DeletePayloadInfo) // 删除payloadInfo表
		payloadInfoRouter.DELETE("deletePayloadInfoByIds", payloadInfoApi.DeletePayloadInfoByIds) // 批量删除payloadInfo表
		payloadInfoRouter.PUT("updatePayloadInfo", payloadInfoApi.UpdatePayloadInfo)    // 更新payloadInfo表
	}
	{
		payloadInfoRouterWithoutRecord.GET("findPayloadInfo", payloadInfoApi.FindPayloadInfo)        // 根据ID获取payloadInfo表
		payloadInfoRouterWithoutRecord.GET("getPayloadInfoList", payloadInfoApi.GetPayloadInfoList)  // 获取payloadInfo表列表
	}
	{
	    payloadInfoRouterWithoutAuth.GET("getPayloadInfoPublic", payloadInfoApi.GetPayloadInfoPublic)  // payloadInfo表开放接口
	}
}
