package controllers

import (
	"DataCertPlatform/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginControllre struct {
	beego.Controller
}

//直接跳转展示用户登录页面
func (l *LoginControllre) Get() {
	l.TplName = "login.html"
}

//处理用户登录请求
func (l *LoginControllre) Post() {
	//1.解析客户端用户提交的登录数据
	var user models.User
	err :=l.ParseForm(&user)
	if err != nil{
		fmt.Println(err.Error())
		l.Ctx.WriteString("抱歉用户登录信息解析解析失败，请重试!")
		return
	}
    //2.根据解析到的数据，执行数据库查询操作
    u,err:=user.QueryUser()

	//3.判断数据库查询结果
	if err !=nil{
		//sql: no rows in result set结果集当中无数据
		fmt.Println(err.Error())
		l.Ctx.WriteString("用户登录失败，请重试!")
		return
	}

/*	//判断用户是否实名认证
	if{
		l.TplName = "user_kyc.html"
	}
*/


	//4.根据查询结果返回客户端相应的信息或者页面跳转
	l.Data["Phone"] = u.Phone//动态数据设置
    l.TplName = "home.html"

}
