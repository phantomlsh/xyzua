// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"xyzua/controllers"
	"xyzua/services"

	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func checkAdmin(ctx *context.Context) bool {
	var input map[string]interface{}
	if err := json.Unmarshal(ctx.Input.RequestBody, &input); err == nil {
		if input["admintoken"] == nil {
			return false
		}
		admintoken := input["admintoken"].(string)
		u := services.QueryIdentity(admintoken, false)
		if u != nil && u.Role == "ADMIN" {
			return true
		}
	}
	return false
}

func init() {

	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/app",
			beego.NSCond(checkAdmin),
			beego.NSInclude(
				&controllers.AppController{},
			),
		),

		beego.NSNamespace("/code",
			beego.NSCond(checkAdmin),
			beego.NSInclude(
				&controllers.CodeController{},
			),
		),

		beego.NSNamespace("/mark",
			beego.NSCond(checkAdmin),
			beego.NSInclude(
				&controllers.MarkController{},
			),
		),

		beego.NSNamespace("/ticket",
			beego.NSCond(checkAdmin),
			beego.NSInclude(
				&controllers.TicketController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSCond(checkAdmin),
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSNamespace("/identity",
			beego.NSInclude(
				&controllers.IdentityController{},
			),
		),

		beego.NSNamespace("/info",
			beego.NSInclude(
				&controllers.InfoController{},
			),
		),

		beego.NSNamespace("/util",
			beego.NSInclude(
				&controllers.UtilController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
