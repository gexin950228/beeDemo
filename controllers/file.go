package controllers

import (
	"beeDemo/utils"
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
)

type File1Controller struct {
	beego.Controller
}

func (f *File1Controller) Get() {
	fmt.Println("================url反转==================")
	fmt.Println(beego.URLFor("File1Controller.Get", "name", "gexin", "age", 28))
	f.Data["xsrfData"] = template.HTML(f.XSRFFormHTML())
	f.TplName = "file/file1.html"
}

func (f *File1Controller) Post() {
	file, h, err := f.GetFile("file")
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	fileName := h.Filename
	f.SaveToFile("file", "upload/"+utils.UniqueName(fileName))
	utils.LogToFile("Info", fmt.Sprintf("%s上传成功", fileName))
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
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	fileName := h.Filename
	fmt.Println(fileName)
	f2.SetSession("beegoSession", "zhiliao")
	f2.SaveToFile("file", "upload/"+utils.UniqueName(fileName))
	f2.Data["xsrfData"] = template.HTML(f2.XSRFFormHTML())
	f2.Data["json"] = map[string]string{"code": "200", "msg": "上传成功", "filename": fileName}
	f2.ServeJSON()
}
