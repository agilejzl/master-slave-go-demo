// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"master-slave-go-demo/controllers"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/orders",
			beego.NSInclude(
				&controllers.OrdersController{},
			),
		),

		beego.NSNamespace("/products",
			beego.NSInclude(
				&controllers.ProductsController{},
			),
		),

		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UsersController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
