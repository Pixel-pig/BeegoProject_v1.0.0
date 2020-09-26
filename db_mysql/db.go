package db_mysql

import (
	"BeeGoProjet_v1.0.0/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB //定义sql的全局变量
//读取配置文件
var Config = beego.AppConfig //配置文件对象

func init() {
	fmt.Println("连接MySQL数据库")
	fmt.Println("应用名称",Config.String("appname"))
	port, err := Config.Int("httpport")
	if err != nil {
		panic("项目配置文件解析错误，请重试")
	}
	fmt.Println(port)

	dbDriver := Config.String("dirverName")
	dbUser := Config.String("db_user")
	dbPassWord := Config.String("db_password")
	dbIP := Config.String("db_ip")
	dbPort := Config.String("db_port")
	dbName := Config.String("db_name")

	//database,err :=sql.Open("mysql","root:zjt3606810515@tcp(127.0.0.1:3306)/movie?charset=utf8")//获取MySQL对象
	connUrl := dbUser + ":" +dbPassWord + "@tcp("+dbIP+":"+dbPort+")/"+dbName+"?charset=utf8"
	db, err := sql.Open(dbDriver,connUrl)
	if err != nil {
		panic("数据库连接错误，请检查配置文件")
	}
	fmt.Println("数据库连接成功")
	DB = db
}


//查询用户名
func QueryUser () {
	//sql语句
	//sqlStr := "select * from " + Config.String("db_name")
	//r := DB.QueryRow(sqlStr)
}

/**
 *将用户信息保存的数据库中
 */
func InserUser(user models.User) (int64, error){

	//将用户密码进行脱敏
	hashMd5:= md5.New()
	hashMd5.Write([]byte(user.Password))
	bytes := hashMd5.Sum(nil)
	user.Password = hex.EncodeToString(bytes)

	fmt.Println(user.UserName,user.Password)
	//sql语句
	sqlStr := "insert into user(id,username, password, address, sex) value(?,?,?,?,?)"
	result, err := DB.Exec(sqlStr,user.Id,user.UserName,user.Password,user.Address,user.Sex)
	if err != nil {
		return -1,err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return  -1,err
	}
	return id,nil
}