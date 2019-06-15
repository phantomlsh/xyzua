package main

import (
	_ "xyzua/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))
	beego.BConfig.WebConfig.StaticDir["/"] = "web"
	beego.Run()
}
