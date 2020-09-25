package main

import (
	_ "BeeGoProjet_v1.0.0/db_mysql"
	_ "BeeGoProjet_v1.0.0/routers"
	"github.com/astaxie/beego"
)

func main() {
	//开启web port
	beego.Run()
}

