package repository

import (
	"echo_golang/database"
)

type CustomerRepo interface {
	GetByID(customer interface{}, id interface{}) error
	GetAll(customers interface{}) error
	Insert(customer interface{}) error
	Update(customer interface{}) error
	Delete(customer interface{}, id interface{}) error
}

func GetByID(customer interface{}, id interface{}) error {
	db := database.GetDBInstance()
	err := db.Where("id = ?", id).First(customer).Error
	return err
}

func GetAll(customers interface{}) error {
	db := database.GetDBInstance()
	err := db.Find(customers).Error
	return err
}

func Insert(customer interface{}) error {
	db := database.GetDBInstance()
	err := db.Create(customer).Error
	return err
}

func Update(customer interface{}) int64 {
	db := database.GetDBInstance()
	rowAffected := db.Model(customer).Updates(customer).RowsAffected

	return rowAffected
}

func Delete(customer interface{}, id interface{}) int64 {
	db := database.GetDBInstance()
	rowAffected := db.Where("id = ?", id).Delete(customer).RowsAffected
	return rowAffected
}
