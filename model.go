package main

type Products struct {
	Id    string  `form:"id" json:"id"`
	Name  string  `form:"name" json:"name"`
	Image string  `form:"image" json:"image"`
	Price float64 `form:"price" json:"price"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Products
}
