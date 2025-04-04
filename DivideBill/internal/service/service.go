package service

import (
	"encoding/json"
	"fmt"

	"github.com/lashkapashka/divideBill/internal/model"
	"github.com/lashkapashka/divideBill/pkg/client"
)

// type DivideService struct {
// 	repo *repository.
// }

func DivideService() {

	var unescaped string

	var dish model.DataDishes
	
	resp := client.Client("http://localhost:8000/cache/get-data/key")
	
	// Избавляемся от экранирования для дальнейшего успешного парсинга
	if err := json.Unmarshal([]byte(resp), &unescaped); err == nil {
		resp = unescaped
	}

	// Парсим строку в объект
	if err := json.Unmarshal([]byte(resp), &dish); err != nil {
		panic(err)
	}

	fmt.Println(dish)
}