package example

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/example"
    exampleReq "github.com/flipped-aurora/gin-vue-admin/server/model/example/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type PayloadInfoApi struct {}



// CreatePayloadInfo 创建payloadInfo表
// @Tags PayloadInfo
// @Summary 创建payloadInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body example.PayloadInfo true "创建payloadInfo表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /payloadInfo/createPayloadInfo [post]
func (payloadInfoApi *PayloadInfoApi) CreatePayloadInfo(c *gin.Context) {
	var payloadInfo example.PayloadInfo
	err := c.ShouldBindJSON(&payloadInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = payloadInfoService.CreatePayloadInfo(&payloadInfo)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeletePayloadInfo 删除payloadInfo表
// @Tags PayloadInfo
// @Summary 删除payloadInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body example.PayloadInfo true "删除payloadInfo表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /payloadInfo/deletePayloadInfo [delete]
func (payloadInfoApi *PayloadInfoApi) DeletePayloadInfo(c *gin.Context) {
	id := c.Query("id")
	err := payloadInfoService.DeletePayloadInfo(id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePayloadInfoByIds 批量删除payloadInfo表
// @Tags PayloadInfo
// @Summary 批量删除payloadInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /payloadInfo/deletePayloadInfoByIds [delete]
func (payloadInfoApi *PayloadInfoApi) DeletePayloadInfoByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := payloadInfoService.DeletePayloadInfoByIds(ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePayloadInfo 更新payloadInfo表
// @Tags PayloadInfo
// @Summary 更新payloadInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body example.PayloadInfo true "更新payloadInfo表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /payloadInfo/updatePayloadInfo [put]
func (payloadInfoApi *PayloadInfoApi) UpdatePayloadInfo(c *gin.Context) {
	var payloadInfo example.PayloadInfo
	err := c.ShouldBindJSON(&payloadInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = payloadInfoService.UpdatePayloadInfo(payloadInfo)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPayloadInfo 用id查询payloadInfo表
// @Tags PayloadInfo
// @Summary 用id查询payloadInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询payloadInfo表"
// @Success 200 {object} response.Response{data=example.PayloadInfo,msg=string} "查询成功"
// @Router /payloadInfo/findPayloadInfo [get]
func (payloadInfoApi *PayloadInfoApi) FindPayloadInfo(c *gin.Context) {
	id := c.Query("id")
	repayloadInfo, err := payloadInfoService.GetPayloadInfo(id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(repayloadInfo, c)
}
// GetPayloadInfoList 分页获取payloadInfo表列表
// @Tags PayloadInfo
// @Summary 分页获取payloadInfo表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query exampleReq.PayloadInfoSearch true "分页获取payloadInfo表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /payloadInfo/getPayloadInfoList [get]
func (payloadInfoApi *PayloadInfoApi) GetPayloadInfoList(c *gin.Context) {
	var pageInfo exampleReq.PayloadInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := payloadInfoService.GetPayloadInfoInfoList(pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetPayloadInfoPublic 不需要鉴权的payloadInfo表接口
// @Tags PayloadInfo
// @Summary 不需要鉴权的payloadInfo表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /payloadInfo/getPayloadInfoPublic [get]
func (payloadInfoApi *PayloadInfoApi) GetPayloadInfoPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    payloadInfoService.GetPayloadInfoPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的payloadInfo表接口信息",
    }, "获取成功", c)
}
