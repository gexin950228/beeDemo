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
		utils.LogToFile("ERROR", fmt.Sprintf("关联查询report标签信息报错，错误信息为: %s", err.Error()))
	}
	var addLabels []models.Labels
	if len(report.Labels) == 0 {
		_, err := o.QueryTable("Labels").All(&addLabels)
		if err != nil {
			utils.LogToFile("ERROR", fmt.Sprintf("report标签为空的时候获取可添加标签失败，错误信息: %s", err.Error()))
		}
	} else {
		_, err = o.QueryTable("labels").Exclude("id__in", report.Labels).All(&addLabels)
		if err != nil {
			utils.LogToFile("ERROR", err.Error())
		}
	}
	r.Data["existLabels"] = report.Labels
	r.Data["report"] = report
	r.Data["addLabels"] = addLabels
	r.TplName = "sqlExe/m2m_add_label.html"
}

func (r *ReportAddTagController) Post() {
	title := r.GetString("title")
	description := r.GetString("description")
	id := r.GetString("id")
	author := r.GetString("author")
	report := models.ReportMany{}
	orm.Debug = true
	o := orm.NewOrm()

	err := o.QueryTable("report_many").Filter("id", id).One(&report)
	if err != nil {
		utils.LogToFile("ERROR", fmt.Sprintf("查询id为%d的report报错，错误信息为: %s", id, err.Error()))
	}
	report.Title = title
	report.Description = description
	report.Author = author
	addLabels := r.GetStrings("addLabels")
	existLabels := r.GetStrings("existLabels")
	var modifiedLabels []models.Labels
	var label = models.Labels{}
	// 获取新增标签
	for i := 0; i < len(addLabels); i++ {
		labelId, _ := strconv.ParseInt(addLabels[i], 10, 64)
		o.QueryTable("labels").Filter("id", labelId).One(&label)
		modifiedLabels = append(modifiedLabels, label)
	}
	// 获取原有标签的修改
	for i := 0; i < len(existLabels); i++ {
		labelId, _ := strconv.ParseInt(existLabels[i], 10, 64)
		o.QueryTable("labels").Filter("id", labelId).One(&label)
		modifiedLabels = append(modifiedLabels, label)
	}
	var pModifyLabels []*models.Labels
	for label := range modifiedLabels {
		pModifyLabels = append(pModifyLabels, &modifiedLabels[label])
	}
	report.Labels = pModifyLabels
	m2m := o.QueryM2M(&report, "Labels")
	// 先删除再添加
	_, err = m2m.Clear()
	if err != nil {
		utils.LogToFile("ERROR", fmt.Sprintf("id为%v的report多对多删除标签报错，错误信息为: %s", id, err.Error()))
	}
	for i := 0; i < len(modifiedLabels); i++ {
		_, err := m2m.Add(modifiedLabels[i])
		if err != nil {
			utils.LogToFile("ERROR", fmt.Sprintf("id为%v的report多对多添加标签%v是报错，错误信息为: %s", id, label, err.Error()))
		}
	}
	r.TplName = "ok.html"
}
