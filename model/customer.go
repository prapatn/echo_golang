package model

type Users struct {
	Id     int    `gorm:"primary_key" json:"id" validate:"numeric"`
	Name   string `json:"name" validate:"required"`
	Age    int    `json:"age" validate:"required,numeric,min=1,max=50"`
	Career string `json:"career" validate:"required"`
	Active bool   `json:"active" validate:"required"`
}

func FullName(FisrtName, LastName string) string {
	return FisrtName + " " + LastName
}

// func (a *Users) Validate() error {
// 	validate := validator.New()
// 	return validate.Struct(a)
// }
