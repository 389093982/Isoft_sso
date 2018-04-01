package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type LoginLog struct {
	Id       				int 		`pk json:"id"`
	UserName 				string 		`json:"user_name"`			// 登录用户名
	LoginIp					string		`json:"login_ip"`			// 登录用户IP
	Origin					string		`json:"origin"`
	Referer					string		`json:"referer"`
	LoginStatus				string		`json:"login_status"`		// 登录状态
	LoginResult				string		`json:"login_result"`		// 登录结果
	CreatedBy				string 		`json:"created_by"`			// 创建人
	CreatedTime				time.Time	`json:"created_time"`		// 创建时间
	LastUpdatedBy			string		`json:"last_updated_by"`	// 修改人
	LastUpdatedTime			time.Time	`json:"last_updated_time"`	// 修改时间
}

func AddLoginLog(log LoginLog) error {
	o := orm.NewOrm()
	_, err := o.Insert(&log)
	return err
}