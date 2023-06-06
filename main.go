package main

import (
	"echo_golang/controller"
	"echo_golang/database"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	database.DB = database.NewDB()

	e.POST("/customers", controller.SaveCustomer)
	e.GET("/customer", controller.GetCustomer)
	e.GET("/customers", controller.GetCustomers)
	e.PUT("/customers", controller.UpdateCustomer)
	e.DELETE("/customers/:id", controller.DeleteCustomer)

	e.Logger.Fatal(e.Start(":8080"))
}
