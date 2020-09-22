package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.baidu.com" //"beego.me"
	c.Data["Email"] = "2282725505@qq.com" //"astaxie@gmail.com"
	c.TplName = "index.tpl"
}
