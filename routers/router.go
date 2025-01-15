package routers

import (
	"beeDemo/controllers"
	"beeDemo/controllers/OrmTest"
	"beeDemo/controllers/User"
	"beeDemo/controllers/relationSql"

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
	beego.Router("/login/?:redirectUri", &User.LoginController{})
	beego.Router("/verifyCode", &User.SendVerifyCodeController{})
	beego.Router("/validator", &controllers.ValicatorController{})
	beego.Router("/abort", &controllers.AbortController{})
	beego.Router("/tpl1", &controllers.TemplateController{})
	beego.Router("/tpl2", &controllers.Template2Controller{})
	beego.Router("/orm1", &OrmTest.TestOrmController{})
	beego.Router("/article", &OrmTest.ArticleController{})
	beego.Router("/addArticle", &OrmTest.AddArticleController{})
	beego.Router("/update/?:id", &OrmTest.UpdateArticleController{})
	beego.Router("/delete/?:id", &OrmTest.DeleteArticleController{})
	beego.Router("/orm3/", &OrmTest.Complex1QueryController{})
	beego.Router("/orm4/", &OrmTest.Complex2QueryController{})
	beego.Router("/native_sql/", &OrmTest.NativeSqlController{})
	beego.Router("/query_article/", &OrmTest.QueryArticleController{})
	beego.Router("/register", &User.RegisterController{})
	beego.Router("/orm5", &OrmTest.OrmInterfaceController{})
	beego.Router("/one_to_one", &relationSql.OneToOneController{})
	beego.Router("/one_to_many", &relationSql.OneToMany{})
}
