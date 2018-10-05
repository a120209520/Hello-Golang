package controllers

import "github.com/astaxie/beego"

type TestController struct {
	beego.Controller
}

type User struct {
	Username string
	Password string
}

func testInput(c *TestController) {
	//id := c.GetString("id")
	id := c.Input().Get("id")
	c.Ctx.WriteString(id)
}

func testInputStruct(c *TestController) {
	user := User{}
	if err := c.ParseForm(&user); err == nil {
		c.Ctx.WriteString("username="+user.Username+"password="+user.Password)
	}
}

func (c *TestController) Get() {
	//testInput(c)
}

func (c *TestController) Post() {
	testInputStruct(c)
}
