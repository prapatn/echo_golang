package controller

import (
	"echo_golang/model"
	"echo_golang/usecase"
	"echo_golang/validate"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func GetCustomer(c echo.Context) error {
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

func GetCustomers(c echo.Context) error {
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

func SaveCustomer(c echo.Context) error {
	customer := new(model.Users)
	err := c.Bind(customer)

	if err != nil {
		responeErr, ok := validate.MapErrorBind(err)
		if ok != nil {
			return c.JSON(http.StatusBadRequest, ok.Error())
		}

		return c.JSON(http.StatusBadRequest, model.ResponseError{
			Code:    http.StatusBadRequest,
			Message: "Fail",
			Errors:  responeErr,
		})
	}
	err = c.Validate(customer)

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

func UpdateCustomer(c echo.Context) error {
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
func DeleteCustomer(c echo.Context) error {
	customer := new(model.DeleteUser)
	err := c.Bind(customer)
	if err != nil {
		responeErr, ok := validate.MapErrorBind(err)
		if ok != nil {
			return c.JSON(http.StatusBadRequest, ok.Error())
		}

		return c.JSON(http.StatusBadRequest, model.ResponseError{
			Code:    http.StatusBadRequest,
			Message: "Fail",
			Errors:  responeErr,
		})
	}
	log.Println(customer.Active)
	err = c.Validate(customer)
	log.Println(err)
	if err != nil {
		errors := validate.MapErrorValidate(err)
		return c.JSON(http.StatusBadRequest, model.ResponseError{
			Code:    http.StatusBadRequest,
			Message: "Fail",
			Errors:  errors,
		})
	}

	err = usecase.Delete(customer.UserId)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "Success")
}
