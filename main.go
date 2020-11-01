package main

import (
	"DataCertPlatform/blockchain"
	_ "DataCertPlatform/routers"
	"github.com/astaxie/beego"

)
//https://github.com/ZJ-p007/ElectronicDataAuthentication.git
func main() {
	/*block0 := blockchain.CreateGenesisBlock() //创建创世区块
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
*/
	//创世区块
	/*gensis :=blockchain.CreateGenesisBlock()
	db,_ :=bolt.Open("chain.db",0600,nil)
	bc := blockchain.BlockChain{
		LastHash:gensis.Hash,
		BoltDB:   db,
	}*/
/*	bc := blockchain.NewBlockChain()//封装
	fmt.Printf("最新区块的哈希:%x\n",bc.LastHash)*/
	//block1,err := bc.AddData([]byte("用户要保存到区块中的数据"))
   /* if err != nil{
   	   fmt.Println(err.Error())
		return
    }

    fmt.Printf("区块高度:%d\n",block1.Height)
    fmt.Printf("区块的哈希值:%d\n",block1.Hash)
	fmt.Printf("前一个区块的哈希值:%x\n",block1.PrevHash)*/
	/*block2,err := bc.QueryBlockRyHeigth(2)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("区块的高度:%d\n",block2)
	fmt.Println("区块中的数据是",string(block2.Data))

	return*/
	/*bc1 := blockchain.NewBlockChain()
	fmt.Println(bc1)
	bc1.AddData([]byte("x-x-x-x-x-x"))
	blocks,err :=bc.QueryAllBlocks()
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	//block是一个切片
	for _, block := range blocks{
		fmt.Printf("序号:%d,区块高度:%d,区块hash:%x",block.Height)
	}

	return
*/

	//先准备一条区块链
	blockchain.NewBlockChain()


	//1.连接数据库
	//dbmysql.Connect()

	//设置静态资源文件映射
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")

	beego.Run()

}
