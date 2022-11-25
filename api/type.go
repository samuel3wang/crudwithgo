package api

import(
	//"github.com/jinzhu/gorm"
)

type People struct {
	ID     string	`json:"id" binding:"required"`
	Name   string	`json:"name"`
	Ex     string	`json:"ex"`
}
type User struct {
	Userid 		string 	`json:"userid"   binding:"required"`
	Account 	string 	`json:"account"  binding:"required"`
	Password 	string 	`json:"password" binding:"required"`
}