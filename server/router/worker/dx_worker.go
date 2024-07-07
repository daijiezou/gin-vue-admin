package worker

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DxWorkerRouter struct {
}

// InitDxWorkerRouter 初始化 dxWorker表 路由信息
func (s *DxWorkerRouter) InitDxWorkerRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	dxWorkerRouter := Router.Group("dxWorker").Use(middleware.OperationRecord())
	dxWorkerRouterWithoutRecord := Router.Group("dxWorker")
	dxWorkerRouterWithoutAuth := PublicRouter.Group("dxWorker")

	var dxWorkerApi = v1.ApiGroupApp.WorkerApiGroup.DxWorkerApi
	{
		dxWorkerRouter.POST("createDxWorker", dxWorkerApi.CreateDxWorker)   // 新建dxWorker表
		dxWorkerRouter.DELETE("deleteDxWorker", dxWorkerApi.DeleteDxWorker) // 删除dxWorker表
		dxWorkerRouter.DELETE("deleteDxWorkerByIds", dxWorkerApi.DeleteDxWorkerByIds) // 批量删除dxWorker表
		dxWorkerRouter.PUT("updateDxWorker", dxWorkerApi.UpdateDxWorker)    // 更新dxWorker表
	}
	{
		dxWorkerRouterWithoutRecord.GET("findDxWorker", dxWorkerApi.FindDxWorker)        // 根据ID获取dxWorker表
		dxWorkerRouterWithoutRecord.GET("getDxWorkerList", dxWorkerApi.GetDxWorkerList)  // 获取dxWorker表列表
	}
	{
	    dxWorkerRouterWithoutAuth.GET("getDxWorkerPublic", dxWorkerApi.GetDxWorkerPublic)  // 获取dxWorker表列表
	}
}
