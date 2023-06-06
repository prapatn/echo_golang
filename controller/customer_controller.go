package controller

import (
	"echo_golang/model"
	"echo_golang/usecase"
	"net/http"

	"github.com/labstack/echo"
)

func GetCustomer(c echo.Context) error {

	id := c.QueryParam("id")
	customer, err := usecase.GetCustomerById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	customer.FisrtName = "test"
	customer.LastName = "test"
	customer.SetFullName()

	return c.JSON(http.StatusOK, customer)
}

func GetCustomers(c echo.Context) error {
	customers, err := usecase.GetCustomers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, customers)
}

func SaveCustomer(c echo.Context) error {
	customer := new(model.Customer)
	// customer := model.Customer{}
	// var customer model.Customer
	if err := c.Bind(&customer); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err := usecase.Insert(customer)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusCreated, "Success")
}

func UpdateCustomer(c echo.Context) error {
	customer := model.Customer{}
	if err := c.Bind(&customer); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err := usecase.Update(&customer)
	if err == 0 {
		return c.String(http.StatusBadRequest, "Not found Customer")
	}
	return c.String(http.StatusOK, "Success")
}
func DeleteCustomer(c echo.Context) error {
	err := usecase.Delete(c.Param("id"))
	if err == 0 {
		return c.String(http.StatusBadRequest, "Not found Customer")
	}
	return c.String(http.StatusOK, "Success")
}
