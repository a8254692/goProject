package controllers

import (
	"database/sql"
	"github.com/astaxie/beego"
)

type MysqlController struct {
	beego.Controller
}

func (c *MysqlController) ShowMysql() {
	conn,err := sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")

	if err != nil {
		c.Ctx.WriteString("1")
		beego.Info("conn errr",err)
		return
	}

	defer conn.Close()

	_,err = conn.Exec("CREATE TABLE userInfo(id int ,name varchar(11))")
	if err != nil {
		c.Ctx.WriteString("2")
		beego.Info("conn errr",err)
		return
	}

	c.Ctx.WriteString("OK")
}
1