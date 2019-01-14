package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "Janking Website"
	c.Data["Email"] = "jankingwon@foxmail.com"
	c.TplName = "index.tpl"
}
func (main *MainController) HelloSitepoint() {
    main.Data["Website"] = "Janking Website"
    main.Data["Email"] = "jankingwon@foxmail.com"
    main.Data["EmailName"] = "Janking"
    main.Data["Id"] = main.Ctx.Input.Param(":id")
    main.TplName = "default/hello-sitepoint.tpl"
}
func (main *MainController) LoginUser() {
	main.Data["Name"] = main.Ctx.Input.Param(":name")
    main.Data["Website"] = main.Ctx.Input.Param(":name") + " Website"
    main.Data["Email"] = main.Ctx.Input.Param(":name") +  "@beego.com"
    main.Data["EmailName"] = main.Ctx.Input.Param(":name")
    main.Data["Id"] = main.Ctx.Input.Param(":id")
    main.TplName = "default/hello-sitepoint.tpl"
}