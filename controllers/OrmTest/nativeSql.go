package OrmTest

import (
	"beeDemo/models"
	"beeDemo/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type NativeSqlController struct {
	beego.Controller
}

func (n *NativeSqlController) Get() {
	o := orm.NewOrm()

	//article := models.Article{
	//	Title:       "了不起的盖茨比",
	//	Author:      "弗·司各特·菲茨杰拉德",
	//	Classify:    "现实",
	//	Description: "揭露了爵士时代的衰落，奠定了作者在现代美国文学史上的地位。",
	//	ReadCount:   9875,
	//	IsDeleted:   0,
	//}
	//sql := fmt.Sprintf("INSERT INTO `article` (`title`, `author`, `description`, `is_deleted`, `read_count`, `classify`) VALUES ('%s', '%s', '%s', %d, %d, '%s');",
	//	article.Title, article.Author, article.Description, article.IsDeleted, article.ReadCount, article.Classify)
	//fmt.Println(sql)
	//
	//insertId, err := raw.Exec()

	sql := fmt.Sprintf("SELECT * FROM article WHERE id = %d;", 16)
	r := o.Raw(sql)
	//var id, read_count, is_deleted int64
	//var title, author, description, classify string
	//err := r.QueryRow(&id, &title, &author, &description, &is_deleted, &read_count, &classify)
	var article models.Article
	err := r.QueryRow(&article)
	if err != nil {
		utils.LogToFile("Error", fmt.Sprintf("执行原生sql失败，错误信息：%s", err.Error()))
		return
	}
	fmt.Println(article)
	n.TplName = "nativeSqlController.tpl"
}

func (n *NativeSqlController) Post() {
	n.TplName = "nativeSqlController.tpl"
}

type QueryArticleController struct {
	beego.Controller
}

func (q *QueryArticleController) Get() {
	var articles []models.Article
	q.Data["article"] = articles
	q.TplName = "queryArticle.html"
}

func (q *QueryArticleController) Post() {
	o := orm.NewOrm()
	qs := o.QueryTable("article")
	var articles []models.Article
	title := q.GetString("title")
	author := q.GetString("author")
	id := q.GetString("id")
	articleInfo := make(map[string]string)
	if title != "" {
		articleInfo["title"] = title
	}
	if author != "" {
		articleInfo["author"] = author
	}
	if id != "" {
		articleInfo["id"] = id
	}
	if len(articleInfo) == 0 {
		q.Data["article"] = articles
		q.TplName = "queryArticle.html"
	} else {
		_, ok := articleInfo["id"]
		if ok {
			articleId, _ := strconv.ParseInt(articleInfo["id"], 10, 64)
			_, err := qs.Filter("id__exact", articleId).All(&articles)
			if err != nil {
				fmt.Println(err)
			}
			q.Data["article"] = articles
			q.TplName = "queryArticle.html"
		} else {
			var articles []models.Article
			articleTitleExist := articleInfo["title"] != ""
			articleAuthorExist := articleInfo["author"] != ""
			if articleTitleExist && articleAuthorExist {
				_, err := qs.Filter("title__icontains", articleInfo["title"]).Filter("author__iexact", articleInfo["author"]).All(&articles)
				if err != nil {
					return
				} else {
					q.Data["articles"] = articles
				}
			} else if articleTitleExist && !articleAuthorExist {
				_, err := qs.Filter("title__icontains", articleInfo["title"]).All(&articles)
				if err != nil {
					return
				} else {
					q.Data["articles"] = articles
				}
			} else if !articleTitleExist && articleAuthorExist {
				_, err := qs.Filter("author__exact", articleInfo["author"]).All(&articles)
				if err != nil {
					utils.LogToFile("Error", err.Error())
				} else {
					q.Data["articles"] = articles
				}
			}
		}
	}
	q.TplName = "queryArticle.html"
}
