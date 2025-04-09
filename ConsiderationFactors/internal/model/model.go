package model


type Factors struct {
	Discounts []Discounts `json:"discounts"`
	Tips []Tips `json:"tips"`
	Promtions []Promtions `json:"promtions"`
}

type Discounts struct {
	Name string `json:"name"`
	Number int `json:"number"`
}

type Tips struct {
	Number int `json:"number"`
}

type Promtions struct {
	Name string `json:"name"`
	Number int `json:"number"`
}

///////////////////////////////

type Products struct {
	Name string `json:"name"`
	NumberServings int `json:"numberServings"`
	Price int `json:"price"`
	TotalPrice int `json:"total"`
}

type DataDishes struct {
	NumberClients int `json:"numberClients"`
	Total_account int `json:"total_account"`
	Products []Products `json:"products"`
}

/////////////////////////////////

type ClientHost struct {
	Port int
	Path string
	Obj string
	Hash string
}