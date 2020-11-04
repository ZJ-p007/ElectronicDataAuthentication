package controllers

import "github.com/astaxie/beego"
//验证码登录
type LoginSmsController struct {
	beego.Controller
}
//发起请求
func (m *LoginSmsController) Get() {
	m.TplName = "login.sms.html"
}
//验证码登录
func (m *LoginSmsController) Post() {

	m.TplName = "login.html"
}
