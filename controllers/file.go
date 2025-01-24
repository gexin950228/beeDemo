package controllers

import (
	"beeDemo/utils"
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	"mime/multipart"
)

type File1Controller struct {
	beego.Controller
}

func (f *File1Controller) Get() {
	//fmt.Println("================url反转==================")
	//fmt.Println(beego.URLFor("File1Controller.Get", "name", "gexin", "age", 28))
	f.Data["xsrfData"] = template.HTML(f.XSRFFormHTML())
	f.TplName = "file/file1.html"
}

func (f *File1Controller) Post() {
	file, h, err := f.GetFile("file")
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	fileName := h.Filename
	err = f.SaveToFile("file", "upload/"+utils.UniqueName(fileName))
	if err != nil {
		utils.LogToFile("ERROR", fmt.Sprintf("报错文件%s出错，错误信息为: %s", fileName, err.Error()))
	}
	utils.LogToFile("Info", fmt.Sprintf("%s文件上传成功", h.Filename))
	f.TplName = "ok.html"
}

type FileAjaxController struct {
	beego.Controller
}

func Prepare(f2 FileAjaxController) {
	f2.EnableXSRF = false
}

func (f2 *FileAjaxController) Get() {
	f2.TplName = "file/file2.html"
}

func (f2 FileAjaxController) Post() {
	file, h, err := f2.GetFile("file")
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			utils.LogToFile("关闭问价你描述符出错从，Error", err.Error())
		}
	}(file)
	fileName := h.Filename
	f2.SetSession("beegoSession", "zhiliao")
	err = f2.SaveToFile("file", "upload/"+utils.UniqueName(fileName))
	if err != nil {
		utils.LogToFile("ERROR", fmt.Sprintf("报错文件出错，错误信息为: %s", err.Error()))
	}
	f2.Data["xsrfData"] = template.HTML(f2.XSRFFormHTML())
	f2.Data["json"] = map[string]string{"code": "200", "msg": "上传成功", "filename": fileName}
	f2.ServeJSON()
}
