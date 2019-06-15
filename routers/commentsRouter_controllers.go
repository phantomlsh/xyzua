package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["xyzua/controllers:AppController"] = append(beego.GlobalControllerRouter["xyzua/controllers:AppController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:AppController"] = append(beego.GlobalControllerRouter["xyzua/controllers:AppController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:AppController"] = append(beego.GlobalControllerRouter["xyzua/controllers:AppController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:AppController"] = append(beego.GlobalControllerRouter["xyzua/controllers:AppController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:AppController"] = append(beego.GlobalControllerRouter["xyzua/controllers:AppController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:CodeController"] = append(beego.GlobalControllerRouter["xyzua/controllers:CodeController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:CodeController"] = append(beego.GlobalControllerRouter["xyzua/controllers:CodeController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:CodeController"] = append(beego.GlobalControllerRouter["xyzua/controllers:CodeController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:CodeController"] = append(beego.GlobalControllerRouter["xyzua/controllers:CodeController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:CodeController"] = append(beego.GlobalControllerRouter["xyzua/controllers:CodeController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:IdentityController"] = append(beego.GlobalControllerRouter["xyzua/controllers:IdentityController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:IdentityController"] = append(beego.GlobalControllerRouter["xyzua/controllers:IdentityController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:IdentityController"] = append(beego.GlobalControllerRouter["xyzua/controllers:IdentityController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:token`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:IdentityController"] = append(beego.GlobalControllerRouter["xyzua/controllers:IdentityController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:token`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:IdentityController"] = append(beego.GlobalControllerRouter["xyzua/controllers:IdentityController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:username`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:InfoController"] = append(beego.GlobalControllerRouter["xyzua/controllers:InfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:InfoController"] = append(beego.GlobalControllerRouter["xyzua/controllers:InfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:InfoController"] = append(beego.GlobalControllerRouter["xyzua/controllers:InfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:InfoController"] = append(beego.GlobalControllerRouter["xyzua/controllers:InfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:InfoController"] = append(beego.GlobalControllerRouter["xyzua/controllers:InfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:token`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:MarkController"] = append(beego.GlobalControllerRouter["xyzua/controllers:MarkController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:MarkController"] = append(beego.GlobalControllerRouter["xyzua/controllers:MarkController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:MarkController"] = append(beego.GlobalControllerRouter["xyzua/controllers:MarkController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:MarkController"] = append(beego.GlobalControllerRouter["xyzua/controllers:MarkController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:MarkController"] = append(beego.GlobalControllerRouter["xyzua/controllers:MarkController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:TicketController"] = append(beego.GlobalControllerRouter["xyzua/controllers:TicketController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:TicketController"] = append(beego.GlobalControllerRouter["xyzua/controllers:TicketController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:TicketController"] = append(beego.GlobalControllerRouter["xyzua/controllers:TicketController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:TicketController"] = append(beego.GlobalControllerRouter["xyzua/controllers:TicketController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:TicketController"] = append(beego.GlobalControllerRouter["xyzua/controllers:TicketController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:UserController"] = append(beego.GlobalControllerRouter["xyzua/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:UserController"] = append(beego.GlobalControllerRouter["xyzua/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:UserController"] = append(beego.GlobalControllerRouter["xyzua/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:UserController"] = append(beego.GlobalControllerRouter["xyzua/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:UserController"] = append(beego.GlobalControllerRouter["xyzua/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:UtilController"] = append(beego.GlobalControllerRouter["xyzua/controllers:UtilController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:UtilController"] = append(beego.GlobalControllerRouter["xyzua/controllers:UtilController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:UtilController"] = append(beego.GlobalControllerRouter["xyzua/controllers:UtilController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:UtilController"] = append(beego.GlobalControllerRouter["xyzua/controllers:UtilController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xyzua/controllers:UtilController"] = append(beego.GlobalControllerRouter["xyzua/controllers:UtilController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:t`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
