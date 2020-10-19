package main

import (
	"DataCertPlatform/blockchain"
	"DataCertPlatform/dbmysql"
	_ "DataCertPlatform/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {

	block0 :=blockchain.CreateGenesisBlock()
    block1 := blockchain.NewBlock(block0.Height+1,block0.Hash,[]byte("a"))
	fmt.Println(block0,block1)
    //return
	//1.连接数据库
	dbmysql.Connect()

	//设置静态资源文件映射
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")
    
	beego.Run()

}
