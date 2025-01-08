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
	//var articles []models.Article
	var article models.Article
	qs := o.QueryTable("article")
	//_, err := qs.Filter("title__iexact", "kubernetes in action").Filter("is_deleted__gte", 0).All(&articles)
	//_, err := qs.Exclude("title", "庆余年").All(&articles)
	//_, err := qs.Exclude("title", "庆余年").Limit(3).Offset(1).All(&articles) // offset limit

	//_, err := qs.GroupBy("Classify").All(&articles)
	err := qs.OrderBy("read_count").One(&article)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(article)
	c2.TplName = "complexQuery2.tpl"
}
