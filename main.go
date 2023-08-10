package main

import (
	"echo_golang/controller"
	"echo_golang/database"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	database.DB = database.NewDB()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "AccessTime => ${time_rfc3339_nano}\n" +
			"Host => ${host}, RemoteIP => ${remote_ip},\n" +
			"Method => ${method},\n" +
			"URI => ${uri}, Status => ${status},\n" +
			"Error => ${error},\n" +
			"UserAgent => ${user_agent}\n" +
			"--------------\n",
		Output: e.Logger.Output(),
	}))

	var controller controller.Controller
	validate := validator.New()
	controller.Validate = validate

	e.POST("/customers", controller.SaveCustomer)
	e.GET("/customer", controller.GetCustomer)
	e.GET("/customers", controller.GetCustomers)
	e.PUT("/customers", controller.UpdateCustomer)
	e.DELETE("/customers/:id", controller.DeleteCustomer)

	e.Logger.Fatal(e.Start(":8080"))
}
