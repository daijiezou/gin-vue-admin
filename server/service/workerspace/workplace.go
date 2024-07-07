package workerspace

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/workerspace"
	workerspaceReq "github.com/flipped-aurora/gin-vue-admin/server/model/workerspace/request"
)

type WorkplaceService struct {
}

// CreateWorkplace 创建工地表记录
// Author [piexlmax](https://github.com/piexlmax)
func (workplaceService *WorkplaceService) CreateWorkplace(workplace *workerspace.Workplace) (err error) {
	err = global.GVA_DB.Create(workplace).Error
	return err
}

// DeleteWorkplace 删除工地表记录
// Author [piexlmax](https://github.com/piexlmax)
func (workplaceService *WorkplaceService) DeleteWorkplace(id string) (err error) {
	err = global.GVA_DB.Delete(&workerspace.Workplace{}, "id = ?", id).Error
	return err
}

// DeleteWorkplaceByIds 批量删除工地表记录
// Author [piexlmax](https://github.com/piexlmax)
func (workplaceService *WorkplaceService) DeleteWorkplaceByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]workerspace.Workplace{}, "id in ?", ids).Error
	return err
}

// UpdateWorkplace 更新工地表记录
// Author [piexlmax](https://github.com/piexlmax)
func (workplaceService *WorkplaceService) UpdateWorkplace(workplace workerspace.Workplace) (err error) {
	err = global.GVA_DB.Model(&workerspace.Workplace{}).Where("id = ?", workplace.Id).Updates(&workplace).Error
	return err
}

// GetWorkplace 根据id获取工地表记录
// Author [piexlmax](https://github.com/piexlmax)
func (workplaceService *WorkplaceService) GetWorkplace(id string) (workplace workerspace.Workplace, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&workplace).Error
	return
}

// GetWorkplaceInfoList 分页获取工地表记录
// Author [piexlmax](https://github.com/piexlmax)
func (workplaceService *WorkplaceService) GetWorkplaceInfoList(info workerspaceReq.WorkplaceSearch) (list []workerspace.Workplace, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&workerspace.Workplace{})
	var workplaces []workerspace.Workplace
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.WorkOrderTime != nil {
		db = db.Where("work_order_time = ?", info.WorkOrderTime)
	}
	if info.Address != "" {
		db = db.Where("address LIKE ?", "%"+info.Address+"%")
	}
	if info.OwnerName != "" {
		db = db.Where("owner_name LIKE ?", "%"+info.OwnerName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["work_order_time"] = true
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

	err = db.Find(&workplaces).Error
	return workplaces, total, err
}
