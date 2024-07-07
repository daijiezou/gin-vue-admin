package workerspace

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WorkplaceRouter struct {
}

// InitWorkplaceRouter 初始化 工地表 路由信息
func (s *WorkplaceRouter) InitWorkplaceRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	workplaceRouter := Router.Group("workplace").Use(middleware.OperationRecord())
	workplaceRouterWithoutRecord := Router.Group("workplace")
	workplaceRouterWithoutAuth := PublicRouter.Group("workplace")

	var workplaceApi = v1.ApiGroupApp.WorkerspaceApiGroup.WorkplaceApi
	{
		workplaceRouter.POST("createWorkplace", workplaceApi.CreateWorkplace)   // 新建工地表
		workplaceRouter.DELETE("deleteWorkplace", workplaceApi.DeleteWorkplace) // 删除工地表
		workplaceRouter.DELETE("deleteWorkplaceByIds", workplaceApi.DeleteWorkplaceByIds) // 批量删除工地表
		workplaceRouter.PUT("updateWorkplace", workplaceApi.UpdateWorkplace)    // 更新工地表
	}
	{
		workplaceRouterWithoutRecord.GET("findWorkplace", workplaceApi.FindWorkplace)        // 根据ID获取工地表
		workplaceRouterWithoutRecord.GET("getWorkplaceList", workplaceApi.GetWorkplaceList)  // 获取工地表列表
	}
	{
	    workplaceRouterWithoutAuth.GET("getWorkplacePublic", workplaceApi.GetWorkplacePublic)  // 获取工地表列表
	}
}
