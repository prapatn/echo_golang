package controller

import (
	"echo_golang/model"
	"echo_golang/usecase"
	"net/http"

	"github.com/labstack/echo"
)

func GetCustomer(c echo.Context) error {
	customer := new(model.Customer)
	id := c.QueryParam("id")
	err := usecase.GetCustomerById(customer, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, customer)
}

func GetCustomers(c echo.Context) error {
	customers := new([]model.Customer)
	err := usecase.GetCustomers(customers)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, customers)
}

func SaveCustomer(c echo.Context) error {
	customer := new(model.Customer)
	err := c.Bind(customer)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = usecase.Insert(customer)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusCreated, "Success")
}

func UpdateCustomer(c echo.Context) error {
	customer := new(model.Customer)
	err := c.Bind(customer)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = usecase.Update(customer)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "Success")
}
func DeleteCustomer(c echo.Context) error {
	err := usecase.Delete(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "Success")
}
