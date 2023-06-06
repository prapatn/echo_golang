package usecase

import (
	"echo_golang/database"
	"echo_golang/model"
	"errors"
	"log"
)

func GetCustomers(customers *[]model.Customer) error {
	db := database.GetDBInstance()
	err := db.Find(&customers).Error
	if err != nil {
		print(err)
		return err
	}

	return nil
}

func GetCustomerById(customer *model.Customer, id string) error {
	db := database.GetDBInstance()
	err := db.Where("id = ?", id).First(&customer).Error
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func Insert(customer *model.Customer) error {
	db := database.GetDBInstance()
	err := db.Create(&customer).Error
	if err != nil {
		print(err)
		return err
	}

	return nil
}

func Update(customer *model.Customer) error {
	db := database.GetDBInstance()
	rowAffected := db.Model(&customer).Updates(customer).RowsAffected
	if rowAffected == 0 {
		return errors.New("Update Fail")
	}
	return nil
}

func Delete(id string) error {
	db := database.GetDBInstance()
	customer := new(model.Customer)
	err := db.Where("id = ?", id).First(&customer).Error
	if err != nil {
		print(err)
		return err
	}
	rowAffected := db.Delete(customer).RowsAffected
	if rowAffected == 0 {
		return errors.New("Delete Fail")
	}
	return nil
}
