package routers

import (
	"myweb/controllers"
	"github.com/astaxie/beego"
)
func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/hello-world/:id([0-9]+)", &controllers.MainController{}, "get:HelloSitepoint")
}
