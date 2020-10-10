package dbmysql

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
)

var Db *sql.DB

//连接mysql数据库
func Connect() {
	//项目配置
	config := beego.AppConfig
	driver := config.String("db_driver")
	dbUser := config.String("db_user ")
	dbPassword := config.String("db_password")
	dbIP := config.String("db_ip")
	dbName := config.String("db_name")
	fmt.Println(dbName, dbPassword, dbUser, dbIP, driver)

	//连接数据库
	connUrl := dbUser +":"+dbPassword+"@tcp("+dbIP+")/"+dbName+"?charset = utf8"
	//db,err :=sql.Open(driver,dbUser+":"+dbPassword+"@tcp("+dbIP+")/"+dbName+"?charset=utf8")
	db, err := sql.Open(driver, connUrl)
	if err != nil { //err不为nil，表示连接数据库时出现错误
		panic("数据库打开失败，请重试") //程序进入恐慌状态，程序会终止执行
	}
	Db = db
	fmt.Println(db)
}
