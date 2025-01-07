package OrmTest

import (
	"beeDemo/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Complex2QueryController struct {
	beego.Controller
}

func (c2 Complex2QueryController) Get() {
	o := orm.NewOrm()
	var articles []models.Article
	qs := o.QueryTable("article")
	_, err := qs.Filter("title__iexact", "kubernetes in action").Filter("is_deleted__gte", 0).All(&articles)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(articles)
	}
	c2.TplName = "complexQuery2.tpl"
}
