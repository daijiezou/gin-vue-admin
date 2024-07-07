import service from '@/utils/request'

// @Tags Workplace
// @Summary 创建工地表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Workplace true "创建工地表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /workplace/createWorkplace [post]
export const createWorkplace = (data) => {
  return service({
    url: '/workplace/createWorkplace',
    method: 'post',
    data
  })
}

// @Tags Workplace
// @Summary 删除工地表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Workplace true "删除工地表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /workplace/deleteWorkplace [delete]
export const deleteWorkplace = (params) => {
  return service({
    url: '/workplace/deleteWorkplace',
    method: 'delete',
    params
  })
}

// @Tags Workplace
// @Summary 批量删除工地表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除工地表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /workplace/deleteWorkplace [delete]
export const deleteWorkplaceByIds = (params) => {
  return service({
    url: '/workplace/deleteWorkplaceByIds',
    method: 'delete',
    params
  })
}

// @Tags Workplace
// @Summary 更新工地表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Workplace true "更新工地表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /workplace/updateWorkplace [put]
export const updateWorkplace = (data) => {
  return service({
    url: '/workplace/updateWorkplace',
    method: 'put',
    data
  })
}

// @Tags Workplace
// @Summary 用id查询工地表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Workplace true "用id查询工地表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /workplace/findWorkplace [get]
export const findWorkplace = (params) => {
  return service({
    url: '/workplace/findWorkplace',
    method: 'get',
    params
  })
}

// @Tags Workplace
// @Summary 分页获取工地表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取工地表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /workplace/getWorkplaceList [get]
export const getWorkplaceList = (params) => {
  return service({
    url: '/workplace/getWorkplaceList',
    method: 'get',
    params
  })
}
