package routers

import (
	"classOne/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.IndexController{},"get:ShowGet;post:Post")

	beego.Router("/mysql", &controllers.MysqlController{},"get:ShowMysql")
}
