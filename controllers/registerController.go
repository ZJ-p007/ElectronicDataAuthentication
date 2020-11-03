package controllers

import (
	"github.com/astaxie/beego"
	"DataCertPlatform/models"
)

type RegisterController struct {
	beego.Controller
}

/**
 * 该方法用于处理用户注册的逻辑
 */
func (r *RegisterController) Post(){
	//1、解析用户端提交的请求数据
	var user models.User
	err := r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("抱歉，数据解析失败,请重试!")
		return
	}

	//2、将解析到的数据保存到数据库中
	_, err = user.AddUser()
	if err != nil {
		r.Ctx.WriteString("抱歉，用户注册失败，请重试!")
		return
	}
	 r.TplName =  "login.html"

}
