package main

import (
	_ "xyzua/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))
  orm.RunSyncdb("default", false, false) // automatically create table
  beego.BConfig.WebConfig.StaticDir["/"] = "web"
	beego.Run()
}
