package main

import (
	"DataCertPlatform/dbmysql"
	_ "DataCertPlatform/routers"
	"github.com/astaxie/beego"
	//mysql驱动
	_"github.com/go-sql-driver/mysql"
)

func main() {

	//1.连接数据库
	dbmysql.Connect()

	//设置静态资源文件映射
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")

	beego.Run()


}

