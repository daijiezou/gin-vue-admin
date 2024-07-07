package worker

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/worker"
    workerReq "github.com/flipped-aurora/gin-vue-admin/server/model/worker/request"
)

type DxWorkerService struct {
}

// CreateDxWorker 创建dxWorker表记录
// Author [piexlmax](https://github.com/piexlmax)
func (dxWorkerService *DxWorkerService) CreateDxWorker(dxWorker *worker.DxWorker) (err error) {
	err = global.GVA_DB.Create(dxWorker).Error
	return err
}

// DeleteDxWorker 删除dxWorker表记录
// Author [piexlmax](https://github.com/piexlmax)
func (dxWorkerService *DxWorkerService)DeleteDxWorker(id string) (err error) {
	err = global.GVA_DB.Delete(&worker.DxWorker{},"id = ?",id).Error
	return err
}

// DeleteDxWorkerByIds 批量删除dxWorker表记录
// Author [piexlmax](https://github.com/piexlmax)
func (dxWorkerService *DxWorkerService)DeleteDxWorkerByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]worker.DxWorker{},"id in ?",ids).Error
	return err
}

// UpdateDxWorker 更新dxWorker表记录
// Author [piexlmax](https://github.com/piexlmax)
func (dxWorkerService *DxWorkerService)UpdateDxWorker(dxWorker worker.DxWorker) (err error) {
	err = global.GVA_DB.Model(&worker.DxWorker{}).Where("id = ?",dxWorker.Id).Updates(&dxWorker).Error
	return err
}

// GetDxWorker 根据id获取dxWorker表记录
// Author [piexlmax](https://github.com/piexlmax)
func (dxWorkerService *DxWorkerService)GetDxWorker(id string) (dxWorker worker.DxWorker, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&dxWorker).Error
	return
}

// GetDxWorkerInfoList 分页获取dxWorker表记录
// Author [piexlmax](https://github.com/piexlmax)
func (dxWorkerService *DxWorkerService)GetDxWorkerInfoList(info workerReq.DxWorkerSearch) (list []worker.DxWorker, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&worker.DxWorker{})
    var dxWorkers []worker.DxWorker
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Address != "" {
        db = db.Where("address LIKE ?","%"+ info.Address+"%")
    }
    if info.Birthday != nil {
        db = db.Where("birthday < ?",info.Birthday)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["birthday"] = true
         	orderMap["sex"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&dxWorkers).Error
	return  dxWorkers, total, err
}