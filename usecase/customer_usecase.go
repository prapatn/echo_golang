package usecase

import (
	"echo_golang/database"
	"echo_golang/model"
	"log"
)

func GetCustomers() ([]model.Customer, error) {
	db := database.GetDBInstance()
	customers := []model.Customer{}

	if err := db.Find(&customers).Error; err != nil {
		print(err)
		return nil, err
	}

	return customers, nil
}

func GetCustomerById(id string) (*model.Customer, error) {
	db := database.GetDBInstance()
	var customer model.Customer
	if err := db.Where("id = ?", id).First(&customer).Error; err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &customer, nil
}

func Insert(customer *model.Customer) error {
	db := database.GetDBInstance()

	if err := db.Create(&customer).Error; err != nil {
		print(err)
		return err
	}

	return nil
}

func Update(customer *model.Customer) int64 {
	db := database.GetDBInstance()

	return db.Model(&customer).Updates(customer).RowsAffected
}

func Delete(id string) int64 {
	db := database.GetDBInstance()
	customer := new(model.Customer)
	if err := db.Where("id = ?", id).First(&customer).Error; err != nil {
		print(err)
		return 0
	}
	return db.Delete(customer).RowsAffected
}
