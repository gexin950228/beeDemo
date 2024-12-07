package routers

import (
	"beeDemo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/greet", &controllers.GreetController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/statics", &controllers.StaticsController{})
	beego.Router("/parameter1", &controllers.ParamaterController{})
	// beego.Router("/parameter1/?:name", &controllers.ParamaterController{})
	// beego.Router("/parameter1/:name", &controllers.ParamaterController{})
	beego.Router("/xml", &controllers.XmlController{})
	beego.Router("/flash", &controllers.FlashController{})
}
