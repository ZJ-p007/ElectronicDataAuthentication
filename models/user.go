package models

import (
	"DataCertPlatform/dbmysql"
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	Id int `form:"id"`
	Phone string `form:"phone"`
	Password string `form:"password"`
}

//将用户的信息保存到数据库中
func (u User) AddUser()(int64,error){
	//将密码进行Hash计算，得到密码hash,脱敏
	HashMd5 := md5.New()
	HashMd5.Write([]byte(u.Password))
	pwdBytes := HashMd5.Sum(nil)

	//把脱敏的密码的md5值重新赋值为密码
	u.Password = hex.EncodeToString(pwdBytes)

	rs,err:=dbmysql.Db.Exec("insert into user(phone,password) values(?,?)",
		u.Phone,u.Password)
	if err !=nil{//保存数据遇到错误
		return -1,err
	}
	id,err := rs.RowsAffected()
	if err != nil{
		return -1,err
	}
	//id代表的是此此数据操作影响的行数,ids是一个整数int64类型
	return id,nil

}

//查询用户信息
func (u User) QuerUser() (*User,error){
	HashMd5 := md5.New()
	HashMd5.Write([]byte(u.Password))
	pwdBytes := HashMd5.Sum(nil)

	//把脱敏的密码的md5值重新赋值为密码
	u.Password = hex.EncodeToString(pwdBytes)

	row :=dbmysql.Db.QueryRow("select phone from user where phone = ? and password = ?",
		u.Phone,u.Password)
	//读取数据
	err := row.Scan(&u.Phone)
  if err != nil{
	  return nil,err
  }
   return &u,nil
}