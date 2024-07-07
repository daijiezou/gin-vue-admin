import service from '@/utils/request'

// @Tags DxWorker
// @Summary 创建dxWorker表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DxWorker true "创建dxWorker表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dxWorker/createDxWorker [post]
export const createDxWorker = (data) => {
  return service({
    url: '/dxWorker/createDxWorker',
    method: 'post',
    data
  })
}

// @Tags DxWorker
// @Summary 删除dxWorker表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DxWorker true "删除dxWorker表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dxWorker/deleteDxWorker [delete]
export const deleteDxWorker = (params) => {
  return service({
    url: '/dxWorker/deleteDxWorker',
    method: 'delete',
    params
  })
}

// @Tags DxWorker
// @Summary 批量删除dxWorker表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除dxWorker表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dxWorker/deleteDxWorker [delete]
export const deleteDxWorkerByIds = (params) => {
  return service({
    url: '/dxWorker/deleteDxWorkerByIds',
    method: 'delete',
    params
  })
}

// @Tags DxWorker
// @Summary 更新dxWorker表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DxWorker true "更新dxWorker表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dxWorker/updateDxWorker [put]
export const updateDxWorker = (data) => {
  return service({
    url: '/dxWorker/updateDxWorker',
    method: 'put',
    data
  })
}

// @Tags DxWorker
// @Summary 用id查询dxWorker表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DxWorker true "用id查询dxWorker表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dxWorker/findDxWorker [get]
export const findDxWorker = (params) => {
  return service({
    url: '/dxWorker/findDxWorker',
    method: 'get',
    params
  })
}

// @Tags DxWorker
// @Summary 分页获取dxWorker表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取dxWorker表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dxWorker/getDxWorkerList [get]
export const getDxWorkerList = (params) => {
  return service({
    url: '/dxWorker/getDxWorkerList',
    method: 'get',
    params
  })
}
