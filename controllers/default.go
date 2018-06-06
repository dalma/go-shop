package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "go shop site"
	c.Data["Email"] = "szrednick@gmail.com"
	c.TplName = "index.tpl"
}
