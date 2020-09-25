package db_mysql

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB //定义sql的全局变量

func init() {
	fmt.Println("连接MySQL数据库")
	//读取配置文件
	config := beego.AppConfig
	fmt.Println("应用名称",config.String("appname"))
	port, err := config.Int("httpport")
	if err != nil {
		panic("项目配置文件解析错误，请重试")
	}
	fmt.Println(port)

	dbDriver := config.String("dirverName")
	dbUser := config.String("db_user")
	dbPassWord := config.String("db_password")
	dbIP := config.String("db_ip")
	dbPort := config.String("db_port")
	dbName := config.String("db_name")

	//database,err :=sql.Open("mysql","root:zjt3606810515@tcp(127.0.0.1:3306)/movie?charset=utf8")//获取MySQL对象
	connUrl := dbUser + ":" +dbPassWord + "@tcp("+dbIP+":"+dbPort+")/"+dbName+"?charset=utf8"
	db, err := sql.Open(dbDriver,connUrl)
	if err != nil {
		panic("数据库连接错误，请检查配置文件")
	}
	fmt.Println("数据库连接成功")
	DB = db
}
