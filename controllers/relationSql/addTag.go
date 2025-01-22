package relationSql

import (
	"beeDemo/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type AddTag struct {
	beego.Controller
}

func (a *AddTag) Get() {
	labelName := a.GetString("label_name")
	fmt.Println(labelName)
	a.TplName = "sqlExe/add_tag.html"
}

func (a *AddTag) Post() {
	var label models.Labels
	labelName := a.GetString("label_name")
	fmt.Println(labelName)
	label.Name = labelName
	o := orm.NewOrm()
	o.QueryTable("labels").Filter("name", label.Name).One(&label)
	if label.Id > 0 {
		fmt.Printf("标签重复, %d\n", label.Id)
	} else {
		_, err := o.Insert(&label)
		if err != nil {
			return
		}
	}
	a.TplName = "ok.html"
}

type ReportAddTagController struct {
	beego.Controller
}

func (r *ReportAddTagController) Get() {
	id := r.GetString("id")
	var report models.ReportMany
	o := orm.NewOrm()
	err := o.QueryTable("report_many").Filter("id", id).One(&report)
	if err != nil {
		fmt.Println(err)
	}
	var labels []models.Labels
	o.QueryTable("labels").All(&labels)
	fmt.Println(report)
	fmt.Println(labels)
	r.Data["report"] = report
	r.Data["labels"] = labels
	r.TplName = "sqlExe/m2m_add_label.html"
}
