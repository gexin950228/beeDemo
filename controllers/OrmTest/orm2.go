package OrmTest

import (
	"beeDemo/controllers"
	"beeDemo/models"
	"beeDemo/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"strconv"
	"strings"
)

type ArticleController struct {
	beego.Controller
}

func (a *ArticleController) Get() {
	orm := orm.NewOrm()
	var articles []models.Article
	qs := orm.QueryTable(new(models.Article)).Filter("IsDeleted", 0)
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
	description := a.GetString("description")
	article.Description = description
	data := map[string]string{"title": title, "author": author, "description": description}
	a.Data["json"] = data
	a.ServeJSON()
}

type AddArticleController struct {
	beego.Controller
}

func (a *AddArticleController) Get() {
	a.TplName = "addArticle.html"
}

func (a *AddArticleController) Post() {
	var article models.Article
	var MessagTpls = map[string]string{
		"Required": "不能为空",
		"MinSize":  "最短长度为 %d",
		"Length":   "长度必须为 %d",
		"Numeric":  "必须是有效的数字",
		"Email":    "无效的电子邮件地址",
		"Mobile":   "无效的手机号码",
		"Min":      "年龄最小不能低于0",
		"Max":      "年龄不能超过150岁",
	}
	validation.SetDefaultMessage(MessagTpls)
	valid := validation.Validation{}
	a.ParseForm(&article)
	valid.Valid(article)
	var validateError controllers.ValidateResult
	if valid.HasErrors() {
		validateError.Code = 0
		validateError.ErrorMsg = []string{}
		for _, err := range valid.Errors {
			//fmt.Println(fmt.Sprintf("%s校验出错，validation error: %s", err.Key, err.Error()))
			validateError.ErrorMsg = append(validateError.ErrorMsg, err.Key+err.Message)
			//fmt.Println(validateError.ErrorMsg)
			utils.LogToFile("Error", fmt.Sprintf("%s校验错误，错误信息: %v", err.Name, err.Error()))
			a.Data["json"] = validateError
			a.ServeJSON()

		}
		return
	}
	//fmt.Println(article)
	orm := orm.NewOrm()
	_, err := orm.Insert(&article)
	if err != nil {
		return
	}

	a.Redirect("/article", 302)
}

type UpdateArticleController struct {
	beego.Controller
}

func (u *UpdateArticleController) Get() {
	var article models.Article
	ids := u.GetString("id")
	id, _ := strconv.Atoi(ids)
	article.Id = id
	o := orm.NewOrm()
	err := o.Read(&article, "id")
	if err != nil {
		utils.LogToFile("Error", fmt.Sprintf("read article err: %s", err.Error()))
	}
	u.Data["article"] = article
	u.TplName = "updateArticle.html"
}

func (u *UpdateArticleController) Post() {
	var article models.Article
	id := u.GetString("id")
	article.Id, _ = strconv.Atoi(id)
	readCount := strings.TrimSpace(u.GetString("read_count"))
	count, _ := strconv.ParseInt(readCount, 10, 64)
	article.ReadCount = count
	article.Description = u.GetString("desc")
	article.Author = u.GetString("auth")
	article.Title = u.GetString("title")
	article.Author = u.GetString("author")
	article.Classify = u.GetString("classify")
	//err := u.ParseForm(&article)

	orm := orm.NewOrm()
	_, err := orm.Update(&article)
	if err != nil {
		utils.LogToFile("Error", fmt.Sprintf("update article err: %s", err.Error()))
		errResponse := map[string]string{"code": "error2", "error": fmt.Sprintf("提交数据错误，错误信息为：%s", err)}
		u.Data["json"] = errResponse
		u.ServeJSON()
	}
	err = orm.Commit()
	u.Redirect("/article", 302)
}

func (u *UpdateArticleController) Delete() {
	var article models.Article
	id := u.GetString("id")
	article.Id, _ = strconv.Atoi(id)
	//fmt.Printf("删除的文章的id是： %v\n", article.Id)
	orm := orm.NewOrm()
	err := u.ParseForm(&article)
	errResponse := map[string]string{"code": "400", "error": fmt.Sprintf("提交数据错误，错误信息为：%s", err)}
	if err != nil {
		utils.LogToFile("Error", fmt.Sprintf("parse article err: %s", err.Error()))
		//fmt.Printf("删除文章出错,错误信息为: %s\n", err.Error())
		u.Data["json"] = errResponse
		u.ServeJSON()
	}
	article.IsDeleted = 1
	_, err = orm.Update(&article)
	if err != nil {
		utils.LogToFile("Error", fmt.Sprintf("update article err: %s", err.Error()))
		u.Data["json"] = errResponse
		u.ServeJSON()
	}
	u.Redirect("/article", 302)
}

type DeleteArticleController struct {
	beego.Controller
}

func (d *DeleteArticleController) Get() {
	id := d.GetString("id")
	idDelete, err := strconv.Atoi(id)
	if err != nil {
		utils.LogToFile("Error", fmt.Sprintf("parse id err: %s", err.Error()))
	}
	var article models.Article
	article.Id = idDelete
	o := orm.NewOrm()
	err = o.Read(&article, "id")
	if err != nil {
		utils.LogToFile("Error", fmt.Sprintf("read article err: %s", err.Error()))
	}
	d.Data["article"] = article
	d.TplName = "deleteArticle.html"
}

func (d *DeleteArticleController) Post() {
	var article models.Article
	var o = orm.NewOrm()
	err := d.ParseForm(&article)
	if err != nil {
		utils.LogToFile("Error", fmt.Sprintf("parse article err: %s", err.Error()))
		//fmt.Printf("删除文章出错,错误信息为: %s\n", err.Error())
	}
	article.IsDeleted = 1
	_, err = o.Update(&article)
	if err != nil {
		utils.LogToFile("Error", fmt.Sprintf("update article err: %s", err.Error()))
	}
	d.Redirect("/article", 302)
}
