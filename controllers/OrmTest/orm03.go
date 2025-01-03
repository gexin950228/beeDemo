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
	article := models.Article{}
	qs := orm.QueryTable(article)
	qs.Filter("id__exact", "2").One(&article) // 等于

	fmt.Println(article)
	c.TplName = "complexQuery.html"
}
