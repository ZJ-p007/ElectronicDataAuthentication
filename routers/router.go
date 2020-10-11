package routers

import (
	"DataCertPlatform/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//router:路由
    beego.Router("/index", &controllers.MainController{})

    //用户注册接口
    beego.Router("/register",&controllers.RegisterController{})
}
