// 自动生成模板DxWorker
package worker

import (
	
	"time"
	
)

// dxWorker表 结构体  DxWorker
type DxWorker struct {

      Address  string `json:"address" form:"address" gorm:"column:address;comment:地址;size:256;"`  //地址 
      Birthday  *time.Time `json:"birthday" form:"birthday" gorm:"column:birthday;comment:生日;"`  //生日 
      CreateTime  *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:;"`  //createTime字段 
      DeleteTime  *int `json:"deleteTime" form:"deleteTime" gorm:"column:delete_time;comment:;size:19;"`  //deleteTime字段 
      Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:10;"`  //id 
      IdNumber  string `json:"idNumber" form:"idNumber" gorm:"column:id_number;comment:身份证号;"`  //身份证号 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:姓名;size:128;"`  //姓名 
      NickName  string `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:昵称;size:128;"`  //昵称 
      Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;size:191;"`  //手机号 
      Remarks  string `json:"remarks" form:"remarks" gorm:"column:remarks;comment:备注;size:1024;"`  //备注 
      Sex  string `json:"sex" form:"sex" gorm:"column:sex;comment:性别，1：男 2：女;size:191;"`  //性别 
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:1:空闲,2:忙碌,3:休息,4:其他;size:10;"`  //状态 
      TotalPeople  *int `json:"totalPeople" form:"totalPeople" gorm:"column:total_people;comment:总人数;size:10;"`  //总人数 
      UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:;"`  //updateTime字段 
}


// TableName dxWorker表 DxWorker自定义表名 dx_worker
func (DxWorker) TableName() string {
  return "dx_worker"
}

