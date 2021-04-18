package main

import (
	"github.com/labstack/echo/v4"
	"golang-echo-rest-api-example/controller"
)

func main() {

	e := echo.New()
	basePath := "/api"
	e.GET(basePath+"/users", controller.GetAllUser)
	e.POST(basePath+"/users", controller.SaveUser)
	e.GET(basePath+"/users/:id", controller.GetUser)
	e.PUT(basePath+"/users/:id", controller.UpdateUser)
	e.DELETE(basePath+"/users/:id", controller.DeleteUser)

	// echo server 9000 de başlatıldı.
	e.Logger.Fatal(e.Start(":9000"))
}
