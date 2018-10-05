package controllers

import "github.com/astaxie/beego"

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "a120209520@icloud.com"
	c.TplName = "index.tpl"
	c.Ctx.WriteString("Test Controller")
}
