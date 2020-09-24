package routers

import (
	"BeeGoProjet_v1.0.0/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/register",&controllers.MainController{})
}
