package model


type Factors struct {
	Discounts []Discounts
	Tips []Tips
	Promtions []Promtions
}

type Discounts struct {
	Name string `json:"name"`
	Number float64 `json:"number"`
}

type Tips struct {
	Number int `json:"number"`
}

type Promtions struct {
	Name string `json:"name"`
	Number int `json:"number"`
}

///////////////////////////////

type DataDishes struct {
	NumberClients int
	Total_account int
	Products []Products
}

type Products struct {
	GuestID int
	Name string
	NumberServings int
	Price int
	TotalPrice int
}