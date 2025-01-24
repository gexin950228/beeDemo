package relationSql

import (
	"beeDemo/models"
	"beeDemo/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
)

type TagListReportController struct {
	beego.Controller
}

func (c *TagListReportController) Get() {
	o := orm.NewOrm()
	var report []models.ReportMany
	var labels []models.Labels
	_, err := o.QueryTable("labels").All(&labels)
	if err != nil {
		utils.LogToFile("ERROR", fmt.Sprintf("查询labels表出错，错误信息为: %v", err))
	}
	o.QueryTable("report_many").Filter("id", 0)
	c.Data["report"] = report
	c.Data["labels"] = labels
	c.TplName = "sqlExe/relation-list.html"
}

type UpdateTagController struct {
	beego.Controller
}

func (u *UpdateTagController) Get() {
	var label models.Labels
	label.Name = strings.TrimSpace(u.GetString("label_name"))
	o := orm.NewOrm()
	_, err := o.Insert(&label)
	if err != nil {
		utils.LogToFile("ERROR", fmt.Sprintf("获取标签名出错: %s", err.Error()))
	}
	u.TplName = "sqlExe/add_tag.html"
}
