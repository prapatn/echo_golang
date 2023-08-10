package controller

import (
	"echo_golang/model"
	"echo_golang/usecase"
	"echo_golang/validate"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type Controller struct {
	Validate *validator.Validate
}

func (controller *Controller) GetCustomer(c echo.Context) error {
	customer := new(model.Users)
	id := c.QueryParam("id")
	err := usecase.GetCustomerById(customer, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{
			Code:    http.StatusBadRequest,
			Message: "Fail",
			Errors:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.ResponseValue{
		Code:    http.StatusOK,
		Message: "Success",
		Value:   customer,
	})
}

func (controller *Controller) GetCustomers(c echo.Context) error {
	customers := new([]model.Users)
	err := usecase.GetCustomers(customers)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{
			Code:    http.StatusBadRequest,
			Message: "Fail",
			Errors:  err.Error(),
		})
	}
	return c.JSON(http.StatusOK, model.ResponseValue{
		Code:    http.StatusOK,
		Message: "Success",
		Value:   customers,
	})
}

func (controller *Controller) SaveCustomer(c echo.Context) error {
	customer := new(model.Users)
	err := c.Bind(customer)

	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{
			Code:    http.StatusBadRequest,
			Message: "Fail",
			Errors:  err.Error(),
		})
	}
	err = controller.Validate.Struct(customer)

	if err != nil {
		errors := validate.MapErrorValidate(err)
		return c.JSON(http.StatusBadRequest, model.ResponseError{
			Code:    http.StatusBadRequest,
			Message: "Fail",
			Errors:  errors,
		})
	}

	err = usecase.Insert(customer)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusCreated, "Success")
}

func (controller *Controller) UpdateCustomer(c echo.Context) error {
	customer := new(model.Users)
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
func (controller *Controller) DeleteCustomer(c echo.Context) error {
	err := usecase.Delete(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "Success")
}
