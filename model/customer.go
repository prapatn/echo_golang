package model

type Users struct {
	Id              int    `gorm:"primary_key" json:"id" validate:"numeric"`
	Name            string `json:"name" validate:"required,customUsernameValidator"`
	Age             int    `json:"age" validate:"required,min=1,max=50"`
	Career          string `json:"career" validate:"required"`
	Active          bool   `json:"active" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"eqfield=Password"`
}

func FullName(FisrtName, LastName string) string {
	return FisrtName + " " + LastName
}
