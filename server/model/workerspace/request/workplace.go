package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type WorkplaceSearch struct {
	WorkOrderTime *time.Time `json:"workOrderTime" form:"workOrderTime" `
	Address       string     `json:"address" form:"address" `
	OwnerName     string     `json:"ownerName" form:"ownerName" `
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
