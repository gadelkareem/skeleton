package main

import (
    "backend/controllers"
    "backend/di"
    "backend/kernel"
    "backend/routers"
)

func main() {
    kernel.Bootstrap()
    routers.InitRouters(di.InitContainer())

    kernel.App.SetupServer(&controllers.ErrorController{})

    kernel.App.RunCommand()
    kernel.App.Run()
}
