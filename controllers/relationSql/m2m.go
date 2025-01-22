package relationSql

import (
	"beeDemo/models"
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
		fmt.Println(err.Error())
	} else {
		fmt.Println(labels)
	}
	o.QueryTable("report_many").Filter("id", 0)
	fmt.Println(report)
	fmt.Println(labels)
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
	fmt.Println(label)
	o := orm.NewOrm()
	o.Insert(&label)
	u.TplName = "sqlExe/add_tag.html"
}
