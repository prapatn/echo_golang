package model

type Customer struct {
	Id        int    `gorm:"primary_key" json:"id"`
	Name      string `json:"name"`
	FisrtName string
	LastName  string
	FullName  string
	Age       string `json:"age"`
	Phone     string `json:"phone"`
}

func FullName(FisrtName, LastName string) string {
	return FisrtName + " " + LastName
}

func (c *Customer) SetFullName() {
	c.FullName = c.FisrtName + " " + c.LastName
}
