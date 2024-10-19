package model

// TODO: Separate DTO from Entity

type User struct {
	Id        string `json:"id" db:"ID"`
	Username  string `json:"username" db:"USERNAME"`
	FirstName string `json:"firstName" db:"FIRST_NAME"`
	LastName  string `json:"lastName" db:"LAST_NAME"`
	Age       int    `json:"age" db:"AGE"`
}
