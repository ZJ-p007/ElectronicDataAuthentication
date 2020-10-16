package main

import (
	"DataCertPlatform/dbmysql"
	_ "DataCertPlatform/routers"
	"github.com/astaxie/beego"
)

func main() {
	//1.连接数据库
	dbmysql.Connect()

	//设置静态资源文件映射
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")

	beego.Run()

}
