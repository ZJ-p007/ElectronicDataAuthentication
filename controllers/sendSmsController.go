package controllers

import (
	"DataCertPlatform/models"
	"github.com/astaxie/beego"
)

type SendSmsController struct {
	beego.Controller
}

func (s *SendSmsController) Post() {
	var smsLogin models.SmsLogin
	err := s.ParseForm()


}
