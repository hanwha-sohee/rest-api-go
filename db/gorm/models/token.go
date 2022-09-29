package models

type Token struct {
	Name string `json:"name" binding:"required""`
	//Id string `json:"id" binding:"required,oneof=USD EUR""`
	Id     string `json:"id" binding:"required"`
	Amount int    `json:"default_Amount" binding:"required"`
}
