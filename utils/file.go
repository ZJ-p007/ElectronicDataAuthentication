package utils

import (
	"io"
	"os"
)

//保存一个文件

func SavaFile(fileName string,file io.Reader) (int64,error) {
	saveFile,err :=os.OpenFile(fileName,os.O_CREATE | os.O_RDWR,777)
	if err!=nil{
		//u.Ctx.WriteString("抱歉，电子数据认证失败，请重试！")
		return -1,err
	}
	//saveFile.Write()
	length,err :=io.Copy(saveFile,file)
	if err != nil{
		//u.Ctx.WriteString("抱歉，电子数据认证失败，请重新尝试！")
		return -1,err
	}
	return length,nil
}