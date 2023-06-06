package model

type Customer struct {
	Id    int    `gorm:"primary_key" json:"id"`
	Name  string `json:"name"`
	Age   string `json:"age"`
	Phone string `json:"phone"`
}

func FullName(FisrtName, LastName string) string {
	return FisrtName + " " + LastName
}
