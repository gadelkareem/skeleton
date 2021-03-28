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
    kernel.App.RunCommand()

    routers.InitRouters(c)
    kernel.App.SetupServer(&controllers.ErrorController{ApiController: controllers.NewApiController(c)})
    c.QueManager.StartWorkers()
    kernel.App.Run()
}
