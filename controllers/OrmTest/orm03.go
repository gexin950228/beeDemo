package OrmTest

import (
	"beeDemo/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ComplexQueryController struct {
	beego.Controller
}

func (c *ComplexQueryController) Get() {
	orm := orm.NewOrm()
	//orm.QueryTable("article")
	var article models.Article
	var articles []models.Article
	qs := orm.QueryTable(article)
	//qs.Filter("id__exact", "2").One(&article) // 等于,大小写敏感
	qs.Filter("title__iexact", "kubernetes in action").All(&articles) // 大小写不敏感
	fmt.Println(articles)
	c.TplName = "complexQuery.html"
}
