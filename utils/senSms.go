package utils

import (
	"encoding/json"
	"github.com/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/astaxie/beego"
)

type SmsCode struct {
	Code string `json:"code"`
}

type SmsResult struct {
	BiaId string
	Code string
	Message string
	RequestId string
}

const SMS_TLP_REGISTER  = ""//注册业务模板


//该函数用于发送信息
func SendSms(phone string, code string, templateType string) (*SmsResult,error) {
	config := beego.AppConfig
	//获取配置文件的sms_access_key
	accessKey := config.String("sms_access_key")
	accessKeySecret := config.String("sms_access_secret")
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", accessKey, accessKeySecret)
	if err != nil {

		return nil,err
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.PhoneNumbers = phone        //指定发送验证码的手机号
	request.SignName = "线上餐厅"           //签名信息
	request.TemplateCode = templateType //指定模板
	smsCode := SmsCode{Code: code}
	smsBytes, _ := json.Marshal(smsCode)
	request.TemplateParam = string(smsBytes)

	response, err := client.SendSms(request)
	if err != nil {

		return nil,err
	}
	//BizId；business 商业，业务
	smsResult := SmsResult{
		BiaId:    response.BizId,
		Code:      response.Code,
		Message:   response.Message,
		RequestId: response.RequestId,
	}
	return &smsResult,nil
}
