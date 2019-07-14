package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type OrmController struct {
	beego.Controller
}

type userInfo struct {
	id int
	Name string
}

func (c *OrmController) ShowMysql() {
	orm.RegisterDataBase("default","mysql","root:123@tcp(127.0.0.1:3306)/test")

	orm.RegisterModel(new(userInfo))

	orm.RunSyncdb("default",false,true)
}
