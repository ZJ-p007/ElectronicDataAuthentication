package models

import "DataCertPlatform/db_mysql"

type SmsRecord struct {
	BizId string
	Phone string
	Code string
	Status string
	massage string
	timeStamp int64
}

func (s  SmsRecord) SaveSmsRecord() {
	rs, err :=db_mysql.Db.Exec("insert into sms_record (biz_id,phone,code,status,message,timestamp ) value (?,?,?,?,?,?)" +
		")")


	
}