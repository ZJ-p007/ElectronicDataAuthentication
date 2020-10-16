package routers

import (
	"DataCertPlatform/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//router:路由
    beego.Router("/", &controllers.MainController{})

    //用户注册接口
    beego.Router("/register",&controllers.RegisterController{})

   //用户登录接口
   beego.Router("/login",&controllers.LoginControllre{})

   //请求直接登录的页面
   beego.Router("/login.html",&controllers.LoginControllre{})

    //登录界面点击注册账号的跳转
    beego.Router("/register.html",&controllers.SkipControllers{})

    beego.Router("/upload",&controllers.UploadFileController{})




}
