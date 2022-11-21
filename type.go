package main

import(
	//"github.com/jinzhu/gorm"
)

type People struct {
	ID     string	`json:"id" binding:"required"`
	Name   string	`json:"name"`
	Ex     string	`json:"ex"`
}