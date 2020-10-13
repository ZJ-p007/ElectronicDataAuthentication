package controllers

import (
	"DataCertPlatform/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}


//该方法用于处理用户注册的逻辑
func (r *RegisterController) Post() {
	//1.解析用户提交的请求数据
	var user models.User
	err := r.ParseForm(&user)
	if err != nil{
		r.Ctx.WriteString("抱歉数据解析失败，请重试!")
		return
	}

	//2.将解析到的数据保存到数据库中
	   _,err =user.AddUser()
       //_,err = dbmysql.AddUser(user)
       if err != nil{
       	r.Ctx.WriteString("数据保存失败")
       	fmt.Println(err.Error())
		   return
	   }
	/*_,err =user.AddUser()
	if err != nil{
		r.Ctx.WriteString("数据保存失败")
		fmt.Println(err.Error())
		return
	}

	*/
	   //fmt.Println(row)
       //.Println(user.Phone)
       //fmt.Println(user.Password)

	//3.将处理结果返回给客户端浏览器
	//3.1:如果成功，跳转登录页面 template模板
	r.TplName = "login.html"
	//3.2.:如果失败，则提示错误信息

}

/*该方法用于处理用户登录的逻辑
  用户名，密码正确，用户登录成功，则跳转到电子数据认证系统查询界面 index.html
 */


