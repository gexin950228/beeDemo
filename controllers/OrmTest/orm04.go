package OrmTest

import (
	"beeDemo/utils"
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
	//var article models.Article
	qs := o.QueryTable("article")
	//_, err := qs.Filter("title__iexact", "kubernetes in action").Filter("is_deleted__gte", 0).All(&articles)
	//_, err := qs.Exclude("title", "庆余年").All(&articles)
	//_, err := qs.Exclude("title", "庆余年").Limit(3).Offset(1).All(&articles) // offset limit

	//_, err := qs.GroupBy("Classify").All(&articles)
	//err := qs.OrderBy("read_count").One(&article)

	// Distinct
	//qs.Distinct().All(&articles, "author")
	//for _, article := range articles {
	//	fmt.Println(article.Author)
	//}
	//number, err := qs.Filter("author__exact", "七牛云团队").Count()
	//fmt.Println(number, err)

	// Exist
	//e := qs.Filter("title__exact", "Kubernetes In Action").Exist()
	//fmt.Println(e)

	// Update
	//updateId, err := qs.Filter("id__exact", 3).Update(orm.Params{"title": "Kubernetes实战"})
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(updateId)
	//}

	// PrepareInsert
	//articles := []models.Article{
	//	{Title: "变形记", Author: "弗兰兹·卡夫卡", Desc: "奥地利作家弗兰兹·卡夫卡的经典作品，通过主人公格里高尔·萨姆沙变成甲虫的故事，揭示了资本主义社会中人的异化和孤独。", Classify: "现实", ReadCount: 7815},
	//	{Title: "傲慢与偏见", Author: "简·奥斯汀", Desc: "描绘了19世纪英国乡村的生活，通过伊丽莎白和达西的爱情故事，探讨了社会阶层和婚姻观念。", Classify: "现实", ReadCount: 2132},
	//}
	//insert, _ := qs.PrepareInsert()
	//for _, article := range articles {
	//	_, err := insert.Insert(&article)
	//	if err != nil {
	//		utils.LogToFile("Error", err.Error())
	//	}
	//}
	//err := insert.Close()
	//if err != nil {
	//	utils.LogToFile("Error", err.Error())
	//}
	//fmt.Println(qs.Filter("id__exact", 3))

	// Values, ValuesMap, ValueFlat
	//var maps []orm.Params
	//num, err := qs.Values(&maps)
	//if err != nil {
	//	fmt.Println(err)
	//	utils.LogToFile("Error", fmt.Sprintf("查询出错，错误信息为: %+v", err))
	//}
	//fmt.Println(num)
	//var valueList []orm.ParamsList
	//list, err := qs.ValuesList(&valueList)
	//if err != nil {
	//	return
	//}
	//fmt.Println(list)
	//fmt.Println(valueList)

	var valueList orm.ParamsList
	num, err := qs.ValuesFlat(&valueList, "title")
	if err != nil {
		utils.LogToFile("Error", err.Error())
	}
	fmt.Println(num, valueList)
	c2.TplName = "complexQuery2.tpl"
}
