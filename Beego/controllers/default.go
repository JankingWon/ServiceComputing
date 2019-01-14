package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Id"] = c.Ctx.Input.Param(":id")
	c.TplName = "index.tpl"


}
func (main *MainController) HelloSitepoint() {
    main.Data["Website"] = "Janking Website"
    main.Data["Email"] = "jankingwon@foxmail.com"
    main.Data["EmailName"] = "Janking"
    main.TplName = "default/hello-sitepoint.tpl"
}
