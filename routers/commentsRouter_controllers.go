package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["master-slave-go-demo/controllers:OrdersController"] = append(beego.GlobalControllerRouter["master-slave-go-demo/controllers:OrdersController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["master-slave-go-demo/controllers:OrdersController"] = append(beego.GlobalControllerRouter["master-slave-go-demo/controllers:OrdersController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["master-slave-go-demo/controllers:OrdersController"] = append(beego.GlobalControllerRouter["master-slave-go-demo/controllers:OrdersController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["master-slave-go-demo/controllers:OrdersController"] = append(beego.GlobalControllerRouter["master-slave-go-demo/controllers:OrdersController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["master-slave-go-demo/controllers:OrdersController"] = append(beego.GlobalControllerRouter["master-slave-go-demo/controllers:OrdersController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["master-slave-go-demo/controllers:ProductsController"] = append(beego.GlobalControllerRouter["master-slave-go-demo/controllers:ProductsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["master-slave-go-demo/controllers:ProductsController"] = append(beego.GlobalControllerRouter["master-slave-go-demo/controllers:ProductsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["master-slave-go-demo/controllers:ProductsController"] = append(beego.GlobalControllerRouter["master-slave-go-demo/controllers:ProductsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["master-slave-go-demo/controllers:ProductsController"] = append(beego.GlobalControllerRouter["master-slave-go-demo/controllers:ProductsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["master-slave-go-demo/controllers:ProductsController"] = append(beego.GlobalControllerRouter["master-slave-go-demo/controllers:ProductsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["master-slave-go-demo/controllers:UsersController"] = append(beego.GlobalControllerRouter["master-slave-go-demo/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["master-slave-go-demo/controllers:UsersController"] = append(beego.GlobalControllerRouter["master-slave-go-demo/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["master-slave-go-demo/controllers:UsersController"] = append(beego.GlobalControllerRouter["master-slave-go-demo/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["master-slave-go-demo/controllers:UsersController"] = append(beego.GlobalControllerRouter["master-slave-go-demo/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["master-slave-go-demo/controllers:UsersController"] = append(beego.GlobalControllerRouter["master-slave-go-demo/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
