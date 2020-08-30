package main

import (
    "backend/controllers"
    "backend/di"
    "backend/kernel"
    "backend/routers"
)

func main() {
    kernel.Bootstrap()
    c := di.InitContainer()
    routers.InitRouters(c)

    kernel.App.SetupServer(&controllers.ErrorController{})

    kernel.App.RunCommand()
    kernel.App.Run(c.QueManager)
}
