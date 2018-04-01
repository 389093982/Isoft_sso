package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type AppRegister struct {
	Id       				int 		`pk json:"id"`
	AppAddress 				string 		`json:"app_address"`		// 注册的应用地址
	CreatedBy				string 		`json:"created_by"`			// 创建人
	CreatedTime				time.Time	`json:"created_time"`		// 创建时间
	LastUpdatedBy			string		`json:"last_updated_by"`	// 修改人
	LastUpdatedTime			time.Time	`json:"last_updated_time"`	// 修改时间
}

// 检查是否经过应用注册
func CheckRegister(app_address string) bool {
	o := orm.NewOrm()
	count, err := o.QueryTable("app_register").Filter("app_address",app_address).Count()
	if err == nil && count > 0{
		return true
	}else{
		return false
	}
}