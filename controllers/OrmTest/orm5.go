package OrmTest

import (
	"beeDemo/models"
	"beeDemo/queryUtils"
	"fmt"
	"github.com/astaxie/beego"
)

type OrmInterfaceController struct {
	beego.Controller
}

func (c *OrmInterfaceController) Get() {
	//o := orm.NewOrm()
	//orm.Debug = true

	// QueryRaws SetArgs接口
	//article := models.Article{}
	//r := o.Raw("SELECT * FROM article WHERE title=?", "红楼梦")
	//err := r.SetArgs("水浒传").QueryRows(&article)
	//if err != nil {
	//	utils.LogToFile("Error", fmt.Sprintf("查询数据错误，错误信息为: %s", err.Error()))
	//}
	//fmt.Println(article)

	var articles []models.Article
	articles = queryUtils.QueryArticlesByTitle("title", "ubernetes")
	fmt.Println(articles)
	c.TplName = "orm5.html"
}
