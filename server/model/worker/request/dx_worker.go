package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type DxWorkerSearch struct{
    
                      Address  string `json:"address" form:"address" `
                      Birthday  *time.Time `json:"birthday" form:"birthday" `
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
