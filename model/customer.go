package model

type Users struct {
	Id              int         `gorm:"primary_key" json:"id" validate:"numeric"`
	Name            string      `json:"name" validate:"required,isValidThai"`
	Age             int         `json:"age" validate:"required,min=1,max=50"`
	Career          string      `json:"career" validate:"required"`
	Active          bool        `json:"active" validate:"required"`
	Password        string      `json:"password" validate:"required"`
	ConfirmPassword string      `json:"confirm_password" validate:"eqfield=Password"`
	Address         Address     `json:"address" validate:"required"`
	Contracts       []*Contract `json:"contracts" validate:"required,empty,dive"`
}

type Address struct {
	Zipcode     string `json:"zipcode" validate:"required,customZipcodeValidator"`
	Province    string `json:"province" validate:"required"`
	District    string `json:"district" validate:"required"`
	SubDistrict string `json:"sub_district" validate:"required"`
}

type Contract struct {
	Type string `json:"type" validate:"required"`
	Text string `json:"text" validate:"required"`
}

func FullName(FisrtName, LastName string) string {
	return FisrtName + " " + LastName
}
