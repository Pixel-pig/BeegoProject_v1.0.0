package controllers

import (
	"BeeGoProjet_v1.0.0/db_mysql"
	"BeeGoProjet_v1.0.0/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type RegisterControllers struct {
	beego.Controller
}


/**
 *练习post类型的请求解析json数据
 */
func (r *RegisterControllers) Post() {
	//定义一个person结构体
	//var person models.Person
	//定义一个user结构体
	var user models.User

	//接收请求的请求体，得到一个[]byte数据
	dataByte, err := ioutil.ReadAll(r.Ctx.Request.Body)
	if err != nil {
		r.Ctx.WriteString("接收数据错误")
		return
	}

	//解析[]byte格式的json数据
	err = json.Unmarshal(dataByte,&user)
	if err != nil {
		r.Ctx.WriteString("解析数据错误")
		return
	}
	result := models.ResponseResult{
		Code:    0,
		Message: "请求成功",
		Data:    nil,
	}
	r.Data["json"] = &result;
	r.ServeJSON()
	//fmt.Println("名字", person.Name)
	//fmt.Println("地址", person.Address)
	//fmt.Println("昵称", person.Nick)
	//fmt.Println("生日", person.Birthday)
	fmt.Println(user.Id,user.Password,user.UserName)
	id, err := db_mysql.InserUser(user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(id)
}