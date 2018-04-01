package routers

import (
	"Isoft_sso/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/user/regist",&controllers.UserController{},"get,post:Regist")
	beego.Router("/user/login",&controllers.UserController{},"get,post:Login")
	beego.Router("/user/logout",&controllers.UserController{},"get,post:Logout")
	beego.Router("/user/checkLogin",&controllers.UserController{},"get,post:CheckLogin")
	beego.Router("/user/deleteToken",&controllers.UserController{},"get,post:DeleteToken")

}

