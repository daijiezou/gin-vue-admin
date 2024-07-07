// 自动生成模板Workplace
package workerspace

import (
	
	"time"
	
)

// 工地表 结构体  Workplace
type Workplace struct {

      WorkOrderTime  *time.Time `json:"workOrderTime" form:"workOrderTime" gorm:"column:work_order_time;comment:用工日期;" binding:"required"`  //日期 
      Address  string `json:"address" form:"address" gorm:"column:address;comment:工地地址;size:128;"`  //地址 
      Car  *int `json:"car" form:"car" gorm:"column:car;comment:关联的车;size:10;"`  //关联的车 
      CreateTime  *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:;"`  //createTime字段 
      DeleteTime  *int `json:"deleteTime" form:"deleteTime" gorm:"column:delete_time;comment:;size:19;"`  //deleteTime字段 
      Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:10;"`  //id字段 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:工地名称;size:128;"`  //工地名称 
      NeedWorkers  *int `json:"needWorkers" form:"needWorkers" gorm:"column:need_workers;comment:需要工人人数;size:10;"`  //工人人数 
      OwnerIdNumber  string `json:"ownerIdNumber" form:"ownerIdNumber" gorm:"column:owner_id_number;comment:工地拥有者身份证号码;"`  //工地拥有者身份证号码 
      OwnerName  string `json:"ownerName" form:"ownerName" gorm:"column:owner_name;comment:拥有者姓名;size:128;"`  //拥有者姓名 
      OwnerPhone  string `json:"ownerPhone" form:"ownerPhone" gorm:"column:owner_phone;comment:联系人号码;"`  //联系人号码 
      RequireDescription  string `json:"requireDescription" form:"requireDescription" gorm:"column:require_description;comment:需求描述;"`  //需求描述 
      UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:;"`  //updateTime字段 
}


// TableName 工地表 Workplace自定义表名 dx_workplace
func (Workplace) TableName() string {
  return "dx_workplace"
}

