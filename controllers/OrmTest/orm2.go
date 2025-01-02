package OrmTest

import (
	"beeDemo/models"
	"beeDemo/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ArticleController struct {
	beego.Controller
}

func (a *ArticleController) Get() {
	orm := orm.NewOrm()
	var articles []models.Article
	qs := orm.QueryTable(new(models.Article))
	_, err := qs.All(&articles)
	if err != nil {
		utils.LogToFile("Error", fmt.Sprintf("read article err: %s", err.Error()))
		a.TplName = "index/error.html"
	} else {
		//fmt.Println(all)
		//fmt.Printf("%v, %T", articles, articles)
		//data := map[string]interface{}{
		//	"Code": "200",
		//	"Data": articles,
		//}
		//a.Data["json"] = data
		//a.ServeJSON()
		a.Data["articles"] = articles
		a.TplName = "article.html"
	}
}

func (a *ArticleController) Post() {
	var article models.Article
	title := a.GetString("title")
	article.Title = title
	author := a.GetString("author")
	article.Author = author
	desc := a.GetString("desc")
	article.Desc = desc
	fmt.Println(article)
	data := map[string]string{"title": title, "author": author, "desc": desc}
	a.Data["json"] = data
	a.ServeJSON()
}

type AddArticle struct {
	beego.Controller
}

func (a *AddArticle) Get() {
	a.TplName = "addArticle.html"
}

func (a *AddArticle) Post() {
	var article models.Article
	article.Title = a.GetString("title")
	article.Desc = a.GetString("desc")
	article.Author = a.GetString("author")
	//fmt.Println(article)
	orm := orm.NewOrm()
	insert, err := orm.Insert(&article)
	if err != nil {
		return
	} else {
		fmt.Println("Insert id:", insert)
	}

	a.Redirect("/article", 302)
}
