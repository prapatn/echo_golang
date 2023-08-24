package usecase

import (
	"echo_golang/model"
	"echo_golang/repository"
	"errors"
	"log"
)

func GetCustomers(customers *[]model.Users) error {
	err := repository.GetAll(customers)
	if err != nil {
		print(err)
		return err
	}

	return nil
}

func GetCustomerById(customer *model.Users, id string) error {
	err := repository.GetByID(customer, id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func Insert(customer *model.Users) error {
	err := repository.Insert(customer)
	if err != nil {
		print(err)
		return err
	}

	return nil
}

func Update(customer *model.Users) error {
	rowAffected := repository.Update(customer)
	if rowAffected == 0 {
		return errors.New("Update Fail")
	}
	return nil
}

func Delete(id int) error {
	customer := new(model.Users)
	rowAffected := repository.Delete(customer, id)
	if rowAffected == 0 {
		return errors.New("Delete Fail")
	}
	return nil
}
