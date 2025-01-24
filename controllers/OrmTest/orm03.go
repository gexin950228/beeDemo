package OrmTest

import (
	"beeDemo/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Complex1QueryController struct {
	beego.Controller
}

func (c1 *Complex1QueryController) Get() {
	o := orm.NewOrm()
	//orm.QueryTable("article")
	var article models.Article
	var articles []models.Article
	qs := o.QueryTable(article)
	//qs.Filter("id__exact", "2").One(&article) // 等于,大小写敏感
	//qs.Filter("title__iexact", "kubernetes in action").All(&articles) // 大小写不敏感
	//fmt.Println(articles)

	//// contains
	//qs.Filter("title__icontains", "kubernetes").All(&articles)
	//fmt.Printf("articles: %v\n", articles)

	// gt gte le lte
	//qs.Filter("read_count__gt", 50000).All(&articles)
	//_, err := qs.Filter("read_count__lte", 5000).Filter("read_count__gt", 1000).All(&articles)
	//if err != nil {
	//	return
	//}

	// startswith endswith istartswith iendswith
	//_, err := qs.Filter("title__startswith", "kubernetes").All(&articles)
	//_, err := qs.Filter("title__istartswith", "kubernetes").All(&articles)
	//_, err := qs.Filter("title__iendswith", "action").All(&articles)
	//_, err := qs.Filter("read_count__isnull", false).All(&articles)
	_, err := qs.Filter("read_count__isnull", true).All(&articles)
	if err != nil {
		return
	}
	c1.Data["articles"] = articles
	c1.TplName = "complexQuery1.html"
}
