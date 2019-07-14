package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["test"] = "qukuailian"
	//c.TplName = "index.tpl"
	c.TplName = "test.html"
}

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Post() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["test"] = "朱盼盼"
	//c.TplName = "index.tpl"
	c.TplName = "test.html"
}

func (c *IndexController) ShowGet() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["test"] = "朱盼盼GET"
	//c.TplName = "index.tpl"
	c.TplName = "test.html"
}