package routers

import (
	"beeDemo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/greet", &controllers.GreetController{})
}
