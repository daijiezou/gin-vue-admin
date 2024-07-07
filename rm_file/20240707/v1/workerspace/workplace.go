package workerspace

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/workerspace"
    workerspaceReq "github.com/flipped-aurora/gin-vue-admin/server/model/workerspace/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type WorkplaceApi struct {
}

var workplaceService = service.ServiceGroupApp.WorkerspaceServiceGroup.WorkplaceService


// CreateWorkplace 创建工地表
// @Tags Workplace
// @Summary 创建工地表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body workerspace.Workplace true "创建工地表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /workplace/createWorkplace [post]
func (workplaceApi *WorkplaceApi) CreateWorkplace(c *gin.Context) {
	var workplace workerspace.Workplace
	err := c.ShouldBindJSON(&workplace)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := workplaceService.CreateWorkplace(&workplace); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWorkplace 删除工地表
// @Tags Workplace
// @Summary 删除工地表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body workerspace.Workplace true "删除工地表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /workplace/deleteWorkplace [delete]
func (workplaceApi *WorkplaceApi) DeleteWorkplace(c *gin.Context) {
	id := c.Query("id")
	if err := workplaceService.DeleteWorkplace(id); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWorkplaceByIds 批量删除工地表
// @Tags Workplace
// @Summary 批量删除工地表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /workplace/deleteWorkplaceByIds [delete]
func (workplaceApi *WorkplaceApi) DeleteWorkplaceByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := workplaceService.DeleteWorkplaceByIds(ids); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWorkplace 更新工地表
// @Tags Workplace
// @Summary 更新工地表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body workerspace.Workplace true "更新工地表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /workplace/updateWorkplace [put]
func (workplaceApi *WorkplaceApi) UpdateWorkplace(c *gin.Context) {
	var workplace workerspace.Workplace
	err := c.ShouldBindJSON(&workplace)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := workplaceService.UpdateWorkplace(workplace); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWorkplace 用id查询工地表
// @Tags Workplace
// @Summary 用id查询工地表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query workerspace.Workplace true "用id查询工地表"
// @Success 200 {object} response.Response{data=object{reworkplace=workerspace.Workplace},msg=string} "查询成功"
// @Router /workplace/findWorkplace [get]
func (workplaceApi *WorkplaceApi) FindWorkplace(c *gin.Context) {
	id := c.Query("id")
	if reworkplace, err := workplaceService.GetWorkplace(id); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(reworkplace, c)
	}
}

// GetWorkplaceList 分页获取工地表列表
// @Tags Workplace
// @Summary 分页获取工地表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query workerspaceReq.WorkplaceSearch true "分页获取工地表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /workplace/getWorkplaceList [get]
func (workplaceApi *WorkplaceApi) GetWorkplaceList(c *gin.Context) {
	var pageInfo workerspaceReq.WorkplaceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := workplaceService.GetWorkplaceInfoList(pageInfo); err != nil {
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

// GetWorkplacePublic 不需要鉴权的工地表接口
// @Tags Workplace
// @Summary 不需要鉴权的工地表接口
// @accept application/json
// @Produce application/json
// @Param data query workerspaceReq.WorkplaceSearch true "分页获取工地表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /workplace/getWorkplacePublic [get]
func (workplaceApi *WorkplaceApi) GetWorkplacePublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的工地表接口信息",
    }, "获取成功", c)
}
