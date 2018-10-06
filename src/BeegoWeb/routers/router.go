package routers

import (
	"BeegoWeb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/test", &controllers.TestController{})   //默认get-->Get()  post-->Post()
    beego.Router("/test/login", &controllers.TestController{}, "*:Login")
    beego.Router("/test/model", &controllers.TestController{}, "*:Model")
}
