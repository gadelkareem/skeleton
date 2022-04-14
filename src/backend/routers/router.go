package routers

import (
	"backend/controllers"
	"backend/di"
	"github.com/astaxie/beego"
)

func InitRouters(c *di.Container) {
	api := controllers.NewApiController(c)
	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/auth",
			beego.NSInclude(
				&controllers.AuthController{ApiController: api},
			),
		),
		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UserController{ApiController: api},
			),
		),
		beego.NSNamespace("/audit-logs",
			beego.NSInclude(
				&controllers.AuditLogController{ApiController: api},
			),
		),
		beego.NSNamespace("/common",
			beego.NSInclude(
				&controllers.CommonController{ApiController: api},
			),
		),
		beego.NSNamespace("/products",
			beego.NSInclude(
				&controllers.ProductController{ApiController: api},
			),
		),
		beego.NSNamespace("/subscriptions",
			beego.NSInclude(
				&controllers.SubscriptionController{ApiController: api},
			),
		),
		beego.NSNamespace("/customers",
			beego.NSInclude(
				&controllers.CustomerController{ApiController: api},
			),
		),
		beego.NSNamespace("/invoices",
			beego.NSInclude(
				&controllers.InvoiceController{ApiController: api},
			),
		),
		beego.NSNamespace("/payments",
			beego.NSInclude(
				&controllers.PaymentController{ApiController: api},
			),
		),
		beego.NSNamespace("/payment-methods",
			beego.NSInclude(
				&controllers.PaymentMethodController{ApiController: api},
			),
		),
	)
	beego.AddNamespace(ns)
}
