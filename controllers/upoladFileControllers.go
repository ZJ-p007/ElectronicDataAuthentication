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


//处理文件上传
type UploadFileController struct {
	beego.Controller
}

func (u *UploadFileController) Get() {
	u.TplName = "home.html"
}

//该post方法用于处理用户在客户端提交的文件

func (u *UploadFileController) Post() {

	phone := u.Ctx.Request.PostFormValue("phone")//
	fmt.Println(phone)
	title := u.Ctx.Request.PostFormValue("upload_title")//用户输入的标题
	fmt.Println("电子数据认证标题",title)
	//用户上传的文件
	file, header, err := u.GetFile("file")

	if err != nil {//解析客户端提交的文件出现错误
		fmt.Println(err.Error())
		u.Ctx.WriteString("抱歉，文件解析失败，请重试！")
		return
	}

	defer file.Close()//延迟执行invalid(无效的) memorey(无效的) or nil pointer dereference:空指针错误
     
	//调用工具函数保存文件
	saveFilePath := "static/upload" + header.Filename
	_,err = utils.SavaFile(saveFilePath,file)
	if err != nil{
		u.Ctx.WriteString("抱歉，文件数据认证失败，请重试！")
		return
	}

	/**saveFile,err :=os.OpenFile(saveFilePath,os.O_CREATE | os.O_RDWR,777)
		if err!=nil{
			u.Ctx.WriteString("抱歉，电子数据认证失败，请重试！")
			return
		}
		//saveFile.Write()
		_,err =io.Copy(saveFile,file)
	    if err != nil{
	    	u.Ctx.WriteString("抱歉，电子数据认证失败，请重新尝试！")
			return
		}

	*/

	//2.计算文件的SHA256值
	fileHash,err :=utils.SHA256HashReader(file)
	//fmt.Println(err.Error())
	fmt.Println(fileHash)
	/**hash256:=sha256.New()
	fileBytes,_ :=ioutil.ReadAll(file)
	hash256.Write(fileBytes)
	hashBytes:=hash256.Sum(nil)
	fmt.Println(hex.EncodeToString(hashBytes))
	*/

	//先查询用户id
	user1,err := models.User{Phone:phone}.QueryUserIdByPhone()
	if err != nil{
		u.Ctx.WriteString("抱歉，电子数据认证失败，请稍后再试！")
		return
	}

	//把上传的文件作为记录保存到数据库中
	//计算md5值
	/**md5Hash := md5.New()
	fileMd5Bytes,err :=ioutil.ReadAll(file)
	md5Hash.Write(fileMd5Bytes)
	*/
	saveFile,err :=os.Open(saveFilePath)
	md5HashString,err :=utils.MD5HashReader(saveFile)
	if err !=nil{
		u.Ctx.WriteString("抱歉，电子数据认证失败")
		return
	}
	//bytes := md5Hash.Sum(nil)

	record := models.UploadRecord{
		UserId:    user1.Id,
		FileName:  header.Filename,
		FileSize:  header.Size,
		FileCert:  md5HashString,
		FileTitle: title,
		CertTime:  time.Now().Unix(),
	}
	//保存到数据库中
	_ ,err =record.SaveRecord()
	if err != nil{
		u.Ctx.WriteString("抱歉，电子数据认证保存失败，请重试！")
		fmt.Println(err.Error())
		return
	}

	//将用户上传的文件的md5值和sha256值保存到区块链上，即上链
    blockchain.CHAIN.AddData([]byte(fileHash))



	//上传文件数据保存到数据库
	records, err := models.QueryRecordsByUserId(user1.Id)
	if err != nil{
		fmt.Println("获取数据列表",err.Error())
		u.Ctx.WriteString("抱歉, 获取电子数据列表失败, 请重新尝试!")
		fmt.Println(err.Error())
		return
	}
	fmt.Println(records)
	u.Data["Records"] = records
	u.Data["Phone"] = phone
	u.TplName = "list_record.html"
	//u.Ctx.WriteString("hello")

}


//处理用户提交的文件
/*func (u *UploadFileController) Post1(){
	//1、解析用户上传的数据及文件内容
	//用户上传的自定义的标题
	title := u.Ctx.Request.PostFormValue("upload_title")//用户输入的标题

	//用户上传的文件
	file, header, err := u.GetFile("yuhongwei")

	if err != nil {//解析客户端提交的文件出现错误
		u.Ctx.WriteString("抱歉，文件解析失败，请重试！")
		return
	}
	defer file.Close()

	fmt.Println("自定义的标题：",title)
	//获得到了上传的文件
	fmt.Println("上传的文件名称:",header.Filename)
	//eg：支持jpg,png类型，不支持jpeg，gif类型
	//文件名： 文件名 + "." + 扩展名

	fileNameSlice := strings.Split(header.Filename,".")
	fileType := fileNameSlice[1]
	fmt.Println(fileNameSlice)
	fmt.Println(":",strings.TrimSpace(fileType))

	isJpg := strings.HasSuffix(header.Filename,".jpg")
	isPng := strings.HasSuffix(header.Filename,".png")
	if !isJpg && !isPng{
		//文件类型不支持
		u.Ctx.WriteString("抱歉，文件类型不符合, 请上传符合格式的文件")
		return
	}

 */

	//if fileType != " jpg" || fileType != "png" {
	//	//文件类型不支持
	//	u.Ctx.WriteString("抱歉，文件类型不符合, 请上传符合格式的文件")
	//	return
	//}

	//文件的大小 200kb
/*	config := beego.AppConfig
	fileSize,err := config.Int64("file_size")

	if header.Size / 1024 > fileSize {
		u.Ctx.WriteString("抱歉，文件大小超出范围，请上传符合要求的文件")
		return
	}

	fmt.Println("上传的文件的大小:",header.Size)//字节大小

	//perm:permission:权限
	/*权限的组成:a+b+c
	①：a文件所有者对文件的操作权限，读4 写2 执行1
	②：b文件所有者所在组的用户操作权限，读4 写2 执行1
	③:其他用户的操作权限，读4 写2 执行1
	*/
	/*savaDir := "static/upload"
	//打开文件①
	_,err =os.Open(savaDir)
	if err !=nil{
		err =os.Mkdir(savaDir,777)
		if err != nil{//file exists
			fmt.Println(err.Error())
			u.Ctx.WriteString("抱歉文件认证遇到错误，请重试")
			return
		}
	}
	//fmt.Println("打开文件夹",f.Name())

	//文件名：文件路径 + 文件名 + "." + "文件扩展名"Filename
	savaName := savaDir +"/" + header.Filename
	fmt.Println("要保存的文件名",savaName)
	//fromFile:文件
	//toFile:要保存的文件路径
	err =u.SaveToFile("yuhongwei",savaName)
	if err != nil{
		fmt.Println(err.Error())
		u.Ctx.WriteString("抱歉，文件认证失败，请重试")
		return
	}
	fmt.Println("上传的文件:",file)
	u.Ctx.WriteString("已获取到上传文件。")
}

	 */
