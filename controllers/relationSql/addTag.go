package relationSql

import (
	"beeDemo/models"
	"beeDemo/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type ReportAddTagController struct {
	beego.Controller
}

type ExistLabels struct {
	Id           int64 `orm:"column(id)"`
	LabelsId     int64 `orm:"column(labelsId)"`
	ReportManyId int64 `orm:"column(report_manyId)"`
}

func removeLabel(existLabels []models.Labels, allLabels []models.Labels) []models.Labels {
	var retireLabels []models.Labels
	var flag bool
	for i := 0; i < len(allLabels); i++ {
		flag = true
		for j := 0; j < len(existLabels); j++ {
			if allLabels[i].Name == existLabels[j].Name {
				flag = false
			}
		}
		if flag {
			retireLabels = append(retireLabels, allLabels[i])
		}
	}
	return retireLabels
}

func (r *ReportAddTagController) Get() {
	o := orm.NewOrm()
	idStr := r.GetString("id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	var report models.ReportMany
	err := o.QueryTable("report_many").Filter("id", id).One(&report)
	if err != nil {
		utils.LogToFile("ERROR", fmt.Sprintf("查询report报错，错误信息为: %s", err.Error()))
	}
	_, err = o.LoadRelated(&report, "Labels")
	if err != nil {
		utils.LogToFile("ERROR", fmt.Sprintf("查询report标签信息报错，错误信息为: %s", err.Error()))
	}
	var allLabels []models.Labels
	var existLabels []models.Labels
	for i := range report.Labels {
		var label models.Labels
		o.QueryTable("labels").Filter("id", report.Labels[i].Id).One(&label)
		existLabels = append(existLabels, label)
	}
	_, err = o.QueryTable("labels").All(&allLabels)
	if err != nil {
		utils.LogToFile("ERROR", fmt.Sprintf("查询所有标签报错，错误信息为: %s", err.Error()))
	}

	retireLabels := removeLabel(existLabels, allLabels)
	r.Data["existLabels"] = existLabels
	r.Data["labels"] = retireLabels
	r.Data["report"] = report
	r.TplName = "sqlExe/m2m_add_label.html"
}

func (r *ReportAddTagController) Post() {
	title := r.GetString("title")
	description := r.GetString("description")
	id := r.GetString("id")
	author := r.GetString("author")
	report := models.ReportMany{}
	parseInt, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		utils.LogToFile("ERROR", err.Error())
	}
	labelsId := r.GetStrings("labels")
	report.Id = int(parseInt)
	o := orm.NewOrm()
	qs := o.QueryTable("report_many").RelatedSel().Filter("id", id)
	err = qs.One(&report)
	if err != nil {
		utils.LogToFile("ERROR", err.Error())
	}
	_, err = o.LoadRelated(&report, "Labels")
	if err != nil {
		utils.LogToFile("ERROR", fmt.Sprintf("查询report出错，错误信息: %v", err.Error()))
	}
	report.Title = title
	report.Description = description
	report.Author = author
	var label []models.Labels
	var updateLabels []models.Labels
	for i := range len(labelsId) - 1 {
		idInt, _ := strconv.ParseInt(labelsId[i], 10, 0)
		if err != nil {
			utils.LogToFile("ERROR", err.Error())
		} else {
			o.QueryTable("Labels").Filter("id", idInt).One(&label)
			updateLabels = append(updateLabels, label[i])
		}
	}
	_, err = o.Update(report)
	if err != nil {
		utils.LogToFile("ERROR", fmt.Sprintf("report更新后结果插入数据库报错，错误信息为: %v", err.Error()))
	}
	r.TplName = "ok.html"
}
