package repository

import (
	"echo_golang/database"
	"echo_golang/model"
)

// type CustomerUsecase struct{
// 	CustomerRepo
// }

// type CustomerRepo interface {
// 	GetByID(customer interface{}, id interface{}) error
// 	GetAll(customers interface{}) error
// 	Insert(customer interface{}) error
// 	Update(customer interface{}) error
// 	Delete(customer interface{}, id interface{}) error
// }

func GetByID(customer *model.Customer, id interface{}) error {
	return database.DB.Where("id = ?", id).First(customer).Error
}

func GetAll(customers *[]model.Customer) error {
	return database.DB.Find(customers).Error
}

func Insert(customer *model.Customer) error {
	return database.DB.Create(customer).Error
}

func Update(customer *model.Customer) int64 {
	return database.DB.Model(customer).Updates(customer).RowsAffected
}

func Delete(customer *model.Customer, id interface{}) int64 {
	return database.DB.Where("id = ?", id).Delete(customer).RowsAffected
}
