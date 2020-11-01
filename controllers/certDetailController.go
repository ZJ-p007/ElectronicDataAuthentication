package controllers

import (
	"DataCertPlatform/blockchain"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type CertDetailController struct {
	beego.Controller
}

//该get方法用于处理浏览器的get请求，查看证书
func (c *CertDetailController) Get() {
	//1.解析和接收前端页面传递的数据
	cert_id := c.GetString("cert_id")

	//2.到区块链上查询区块数据
    block,err := blockchain.CHAIN.QueryBlockByCertUd(cert_id)
    if err != nil{
    	c.Ctx.WriteString("抱歉，查询链上数据，请重试")
	}
	if block == nil{//遍历整条区块链，但未查到数据
		c.Ctx.WriteString("抱歉，未查询到链上数据，请重试")
		return
	}
	fmt.Println("查询到的区块高度",block.Height)
    //certId := string(block.Data)
    c.Data["CertId"] =strings.ToUpper(string(block.Data))


	//3.跳转证书页面
	c.TplName = "cert_detail.html"
}
