package queryUtils

import (
	"beeDemo/models"
	"fmt"
	"github.com/astaxie/beego/orm"
)

func QueryArticlesByTitle(fieldName, fieldValue string) []models.Article {
	orm.Debug = true
	o := orm.NewOrm()
	var articles []models.Article
	r := o.Raw(fmt.Sprintf("SELECT * FROM  article WHERE %s LIKE ?", fieldName))
	r.SetArgs("%" + fieldValue + "%").QueryRows(&articles)
	return articles
}
