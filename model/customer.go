package model

type Users struct {
	Id     int    `gorm:"primary_key" json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Career string `json:"career"`
	Active bool   `json:"active"`
}

func FullName(FisrtName, LastName string) string {
	return FisrtName + " " + LastName
}
