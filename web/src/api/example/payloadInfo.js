import service from '@/utils/request'
// @Tags PayloadInfo
// @Summary 创建payloadInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PayloadInfo true "创建payloadInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /payloadInfo/createPayloadInfo [post]
export const createPayloadInfo = (data) => {
  return service({
    url: '/payloadInfo/createPayloadInfo',
    method: 'post',
    data
  })
}

// @Tags PayloadInfo
// @Summary 删除payloadInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PayloadInfo true "删除payloadInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payloadInfo/deletePayloadInfo [delete]
export const deletePayloadInfo = (params) => {
  return service({
    url: '/payloadInfo/deletePayloadInfo',
    method: 'delete',
    params
  })
}

// @Tags PayloadInfo
// @Summary 批量删除payloadInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除payloadInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payloadInfo/deletePayloadInfo [delete]
export const deletePayloadInfoByIds = (params) => {
  return service({
    url: '/payloadInfo/deletePayloadInfoByIds',
    method: 'delete',
    params
  })
}

// @Tags PayloadInfo
// @Summary 更新payloadInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PayloadInfo true "更新payloadInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /payloadInfo/updatePayloadInfo [put]
export const updatePayloadInfo = (data) => {
  return service({
    url: '/payloadInfo/updatePayloadInfo',
    method: 'put',
    data
  })
}

// @Tags PayloadInfo
// @Summary 用id查询payloadInfo表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.PayloadInfo true "用id查询payloadInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /payloadInfo/findPayloadInfo [get]
export const findPayloadInfo = (params) => {
  return service({
    url: '/payloadInfo/findPayloadInfo',
    method: 'get',
    params
  })
}

// @Tags PayloadInfo
// @Summary 分页获取payloadInfo表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取payloadInfo表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payloadInfo/getPayloadInfoList [get]
export const getPayloadInfoList = (params) => {
  return service({
    url: '/payloadInfo/getPayloadInfoList',
    method: 'get',
    params
  })
}

// @Tags PayloadInfo
// @Summary 不需要鉴权的payloadInfo表接口
// @Accept application/json
// @Produce application/json
// @Param data query exampleReq.PayloadInfoSearch true "分页获取payloadInfo表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /payloadInfo/getPayloadInfoPublic [get]
export const getPayloadInfoPublic = () => {
  return service({
    url: '/payloadInfo/getPayloadInfoPublic',
    method: 'get',
  })
}
