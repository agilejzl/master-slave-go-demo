package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["master-slave-go-demo/controllers/api:OrdersController"] = append(beego.GlobalControllerRouter["master-slave-go-demo/controllers/api:OrdersController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["master-slave-go-demo/controllers/api:OrdersController"] = append(beego.GlobalControllerRouter["master-slave-go-demo/controllers/api:OrdersController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["master-slave-go-demo/controllers/api:OrdersController"] = append(beego.GlobalControllerRouter["master-slave-go-demo/controllers/api:OrdersController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["master-slave-go-demo/controllers/api:ProductsController"] = append(beego.GlobalControllerRouter["master-slave-go-demo/controllers/api:ProductsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["master-slave-go-demo/controllers/api:ProductsController"] = append(beego.GlobalControllerRouter["master-slave-go-demo/controllers/api:ProductsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["master-slave-go-demo/controllers/api:ProductsController"] = append(beego.GlobalControllerRouter["master-slave-go-demo/controllers/api:ProductsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["master-slave-go-demo/controllers/api:UsersController"] = append(beego.GlobalControllerRouter["master-slave-go-demo/controllers/api:UsersController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["master-slave-go-demo/controllers/api:UsersController"] = append(beego.GlobalControllerRouter["master-slave-go-demo/controllers/api:UsersController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
