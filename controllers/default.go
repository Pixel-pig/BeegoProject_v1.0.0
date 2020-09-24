package controllers

import (
	"BeeGoProjet_v1.0.0/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}


//get类型的请求
func (c *MainController) Get() {
	c.Data["Website"] = "www.baidu.com" //"beego.me"
	c.Data["Email"] = "2282725505@qq.com" //"astaxie@gmail.com"
	c.TplName = "index.tpl"

	//获取用户输入的数据get接收url地址参数的的方法
	user := c.Ctx.Input.Query("User")//key 为前端传递的字段名
	pwd := c.Ctx.Input.Query("password")

	//将数据写在HTML上
	c.Ctx.WriteString(user);
	c.Ctx.WriteString(pwd);

	//数据校验(响应)
	if user != "pixel-pig" || pwd != "123456" {
		c.Ctx.WriteString("数据校验错误")
	} else {
		c.Ctx.ResponseWriter.Write([]byte("数据校验成功"))
	}
}



/**post类型的请求**/
//func (c *MainController) Post() {
//	//1.接收post请求的参数
//	name := c.Ctx.Request.FormValue("name");
//	age := c.Ctx.Request.FormValue("age");
//	//sex := c.Ctx.Request.FormValue("sex");
//	//fmt.Println(name,age)
//
//	//2.数据校验
//	if name != "admin" || age != "18" {
//		c.Ctx.WriteString("数据校验失败")
//		return
//	} else {
//		c.Ctx.WriteString("数据校验成功")
//	}
//}

//post类型的请求(解析json类型)
func (c *MainController) Post() {
	//定义一个接受json的结构体
	var person models.Person

	//c.Ctx.Request.Body请求体数据
	data, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("数据接收错误")
		return
	}
	err = json.Unmarshal(data,&person)
	if err != nil {
		c.Ctx.WriteString("数据解析错误")
		return
	}
	c.Ctx.WriteString("数据接收成功")
	fmt.Println("姓名",person.Name)
	fmt.Println("年龄",person.Age)
	fmt.Println("性别",person.Sex)
}


