package controllers

import (
	"BeegoWeb/models"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

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
	testInput(c)
}

func (c *TestController) Post() {
	testInputStruct(c)
}

func (c *TestController) Login() {
	cookieUser := c.Ctx.GetCookie("username")
	cookiePass := c.Ctx.GetCookie("password")
	fmt.Println("cookieUser=", cookieUser)
	fmt.Println("cookiePass=", cookiePass)

	user := User{}
	if err := c.ParseForm(&user); err == nil {
		c.Ctx.SetCookie("username", user.Username)
		c.Ctx.SetCookie("password", user.Password)
		c.Ctx.WriteString("username="+user.Username+"password="+user.Password)
	}
}

func (c *TestController) Model() {
	c.Ctx.WriteString("last insert id="+strconv.Itoa(models.TestModel()))
}
