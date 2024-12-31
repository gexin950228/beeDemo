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
	article := models.Article{Id: 1}
	err := orm.Read(&article)
	if err != nil {
		utils.LogToFile("Error", fmt.Sprintf("read article err: %s", err.Error()))
		a.TplName = "index/error.html"
	} else {
		a.Data["article"] = article
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
