package main

import (
	"DataCertPlatform/blockchain"
	"DataCertPlatform/dbmysql"
	_ "DataCertPlatform/routers"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	block0 := blockchain.CreateGenesisBlock() //创建创世区块
	fmt.Println(block0)

	block1 := blockchain.NewBlock(
		block0.Height+1,
		block0.Hash,
		[]byte{})

	fmt.Printf("block0的哈希:%x\n", block0.Hash)
	fmt.Printf("bloc1的哈希:%x\n", block1.Hash)
	fmt.Printf("block1的PrevHash:%x\n", block1.PrevHash)

	//序列化
	blockJson, _ := json.Marshal(block0)
	fmt.Println("通过json序列化以后的block", string(blockJson))

	block0Bytes := block0.Serialize()
	fmt.Println("创世区块通过gob序列化后:", block0Bytes)
	deBlock0, err := blockchain.DSerialize(block0Bytes)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("反序列化后区块的的高度是:", deBlock0.Height)
	fmt.Printf("反序列化后的区块的哈希:%x\n", deBlock0.Hash)

	//return

	//1.连接数据库
	dbmysql.Connect()

	//设置静态资源文件映射
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")

	beego.Run()

}
