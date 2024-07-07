package worker

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/worker"
    workerReq "github.com/flipped-aurora/gin-vue-admin/server/model/worker/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type DxWorkerApi struct {
}

var dxWorkerService = service.ServiceGroupApp.WorkerServiceGroup.DxWorkerService


// CreateDxWorker 创建dxWorker表
// @Tags DxWorker
// @Summary 创建dxWorker表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body worker.DxWorker true "创建dxWorker表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /dxWorker/createDxWorker [post]
func (dxWorkerApi *DxWorkerApi) CreateDxWorker(c *gin.Context) {
	var dxWorker worker.DxWorker
	err := c.ShouldBindJSON(&dxWorker)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := dxWorkerService.CreateDxWorker(&dxWorker); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDxWorker 删除dxWorker表
// @Tags DxWorker
// @Summary 删除dxWorker表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body worker.DxWorker true "删除dxWorker表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /dxWorker/deleteDxWorker [delete]
func (dxWorkerApi *DxWorkerApi) DeleteDxWorker(c *gin.Context) {
	id := c.Query("id")
	if err := dxWorkerService.DeleteDxWorker(id); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDxWorkerByIds 批量删除dxWorker表
// @Tags DxWorker
// @Summary 批量删除dxWorker表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /dxWorker/deleteDxWorkerByIds [delete]
func (dxWorkerApi *DxWorkerApi) DeleteDxWorkerByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := dxWorkerService.DeleteDxWorkerByIds(ids); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDxWorker 更新dxWorker表
// @Tags DxWorker
// @Summary 更新dxWorker表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body worker.DxWorker true "更新dxWorker表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /dxWorker/updateDxWorker [put]
func (dxWorkerApi *DxWorkerApi) UpdateDxWorker(c *gin.Context) {
	var dxWorker worker.DxWorker
	err := c.ShouldBindJSON(&dxWorker)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := dxWorkerService.UpdateDxWorker(dxWorker); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDxWorker 用id查询dxWorker表
// @Tags DxWorker
// @Summary 用id查询dxWorker表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query worker.DxWorker true "用id查询dxWorker表"
// @Success 200 {object} response.Response{data=object{redxWorker=worker.DxWorker},msg=string} "查询成功"
// @Router /dxWorker/findDxWorker [get]
func (dxWorkerApi *DxWorkerApi) FindDxWorker(c *gin.Context) {
	id := c.Query("id")
	if redxWorker, err := dxWorkerService.GetDxWorker(id); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(redxWorker, c)
	}
}

// GetDxWorkerList 分页获取dxWorker表列表
// @Tags DxWorker
// @Summary 分页获取dxWorker表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query workerReq.DxWorkerSearch true "分页获取dxWorker表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /dxWorker/getDxWorkerList [get]
func (dxWorkerApi *DxWorkerApi) GetDxWorkerList(c *gin.Context) {
	var pageInfo workerReq.DxWorkerSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := dxWorkerService.GetDxWorkerInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

// GetDxWorkerPublic 不需要鉴权的dxWorker表接口
// @Tags DxWorker
// @Summary 不需要鉴权的dxWorker表接口
// @accept application/json
// @Produce application/json
// @Param data query workerReq.DxWorkerSearch true "分页获取dxWorker表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /dxWorker/getDxWorkerPublic [get]
func (dxWorkerApi *DxWorkerApi) GetDxWorkerPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的dxWorker表接口信息",
    }, "获取成功", c)
}
