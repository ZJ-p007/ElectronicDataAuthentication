package controllers

import "github.com/astaxie/beego"

type SkipControllers struct {

	beego.Controller
}

//处理登录界面中的注册账号跳转
func (m *SkipControllers) Get() {
	m.TplName = "register.html"
}
