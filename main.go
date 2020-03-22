package main

import (
	"mercari-clone-backend/controllers"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/items", controllers.ItemIndex)
	e.GET("/items/:id", controllers.ItemShow)
	e.POST("/items", controllers.ItemNew)
	e.PUT("/items/:id", controllers.ItemEdit)
	e.DELETE("/items/:id", controllers.ItemDelete)

	e.Logger.Fatal(e.Start(":8080"))
}
