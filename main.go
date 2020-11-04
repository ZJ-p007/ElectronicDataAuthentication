package main

import (
	"DataCertPlatform/db_mysql"
	"DataCertProject/blockchain"
	_ "DataCertProject/routers"
	"github.com/astaxie/beego"
)
//https://github.com/ZJ-p007/ElectronicDataAuthentication.git
func main() {

	//先准备一条区块链
	//blockchain.NewBlockChain()
   blockchain.NewBlockChain()
	//连接数据库
	db_mysql.Connect()

	//设置静态资源文件映射
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")

	beego.Run() //阻塞g

}
