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
    )
    beego.AddNamespace(ns)
}
