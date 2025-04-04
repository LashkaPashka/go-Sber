package model

type Products struct {
	Name string `json:"name"`
	NumberServings int `json:"numberServings"`
	Price float64 `json:"price"`
	TotalPrice float64 `json:"total"`
}

type DataDishes struct {
	NumberClients int `json:"numberClients"`
	Total_account float64 `json:"total_account"`
	Products []Products `json:"products"`
}

type Response struct {
	Position string
	Account string
}