package OrmTest

import (
	"github.com/astaxie/beego"
)

type Complex2QueryController struct {
	beego.Controller
}

func (c2 Complex2QueryController) Get() {
	c2.TplName = "complexQuery2.html"
}
