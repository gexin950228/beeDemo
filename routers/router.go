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
	// beego.Router("/parameter1", &controllers.ParamaterController{})
	//  正则路由
	// beego.Router("/parameter1/:name:string", &controllers.ParamaterController{})
	// 自动路由
	// beego.AutoRouter(&controllers.ParameterController{})
	// beego.Router("/parameter1/:name", &controllers.ParamaterController{})
	//  自定义路由
	beego.Router("/parameter/?:name", &controllers.ParameterController{}, "get:Get")
	beego.Router("/parameter/", &controllers.ParameterController{}, "post:Post")
	beego.Router("/xml", &controllers.XmlController{})
	beego.Router("/flash", &controllers.FlashController{})
	beego.Router("/xsrf", &controllers.TestXsrfController{})
	beego.Router("/file1", &controllers.File1Controller{})
	beego.Router("/file2", &controllers.FileAjaxController{})
	beego.Router("/login/?:redirectUri", &controllers.LoginController{})
	beego.Router("/loginVerifyCode", &controllers.SendVerifyCodeController{})
}
