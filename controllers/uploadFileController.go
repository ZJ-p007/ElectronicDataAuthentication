package controllers

import (
	"DataCertPlatform/blockchain"
	"DataCertPlatform/models"
	"DataCertPlatform/utils"
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"time"
)

/**
 * 该控制器结构体用于处理文件上传的功能
 */
type UploadFileController struct {
	beego.Controller
}

/**
 * 该post方法用于处理用户在客户端提交的文件
 */
func (u *UploadFileController) Post() {
	//1、解析客户端提交的数据和文件
	phone := u.Ctx.Request.PostFormValue("phone")        //获取用户的phone信息
	title := u.Ctx.Request.PostFormValue("upload_title") //用户输入的标题

	fmt.Println("电子数据标签：", title)
	file, header, err := u.GetFile("fils")
	if err != nil { //解析客户端提交的文件出现错误
		u.Ctx.WriteString("抱歉，文件解析失败，请重试！")
		return
	}

	defer file.Close() //延迟执行 空指针错误：invalid memory or nil pointer dereferenece

	//2、调用工具函数保存文件到本地
	saveFilePath := "static/upload/" + header.Filename
	_, err = utils.SaveFile(saveFilePath, file)
	if err != nil {
		u.Ctx.WriteString("抱歉，文件数据认证失败，请重试!")
		return
	}

	//3、计算文件的SHA256值
	fileHash, err := utils.SHA256HashReader(file)
	//fmt.Println(fileHash)

	//先查询用户id
	user1, err := models.User{Phone: phone}.QueryUserByPhone()
	if err != nil {
		fmt.Println("查询用户:", err.Error())
		u.Ctx.WriteString("抱歉，电子数据认证失败，请稍后再试!")
		return
	}

	//把上传的文件作为记录保存到数据库当中
	//① 计算md5值
	saveFile, err := os.Open(saveFilePath)
	md5String, err := utils.MD5HashReader(saveFile)
	if err != nil {
		u.Ctx.WriteString("抱歉, 电子数据认证失败。")
		return
	}
	record := models.UploadRecord{
		UserId:    user1.Id,
		FileName:  header.Filename,
		FileSize:  header.Size,
		FileCert:  md5String,
		FileTitle: title,
		CertTime:  time.Now().Unix(),
	}
	//② 保存认证数据到数据库中
	_, err = record.SaveRecord()
	if err != nil {
		fmt.Println("保存认证记录:", err.Error())
		u.Ctx.WriteString("抱歉，电子数据认证保存失败，请稍后再试!")
		return
	}

	user := &models.User{
		Phone: phone,
	}
	user, _ = user.QueryUserByPhone()
	fmt.Println("用户的信息：", user.Name, user.Phone, user.Card)
	//③ 将用户上传的文件的md5值和sha256值保存到区块链上，即数据上链
	certRecord := models.CertRecord{
		CertId:   []byte(md5String),
		CertHash: []byte(fileHash),
		CertName: user.Name,
		CertCard: user.Card,
		Phone:    user.Phone,
		FileName: header.Filename,
		FileSize: header.Size,
		CertTime: time.Now().Unix(),
	}
	//序列化
	certBytes, _ := certRecord.Serialize()
	_, err = blockchain.CHAIN.SaveData(certBytes)
	if err != nil {
		u.Ctx.WriteString("抱歉，数据上链错误：" + err.Error())
		return
	}
	//fmt.Println("恭喜，已经数据保存到区块链中，区块高度是:", block.Height)

	//上传文件保存到数据库中成功
	records, err := models.QueryRecordsByUserId(user1.Id)
	if err != nil {
		fmt.Println("获取数据列表:", err.Error())
		u.Ctx.WriteString("抱歉, 获取电子数据列表失败, 请重新尝试!")
		return
	}
	u.Data["Records"] = records
	u.TplName = "list_record.html"
}

