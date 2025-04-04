package service

import (
	"encoding/json"
	"fmt"

	"github.com/lashkapashka/divideBill/internal/model"
	"github.com/lashkapashka/divideBill/pkg/client"
	convertstruct "github.com/lashkapashka/divideBill/pkg/convertStruct"
	"github.com/lashkapashka/divideBill/pkg/queue"
	"github.com/lashkapashka/divideBill/pkg/split"
)

type DivideService struct {
	rabbit *queue.RabbitMQ
}

func New() *DivideService {
	return &DivideService{
		rabbit: queue.New(),
	}
}

func (d *DivideService) Divide() {
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

	var msg model.Response

	mp1 := split.SplitPosition([]string{"Garlic Bread"}, &dish)
	mJson := convertstruct.ConvertType(mp1)

	msg.Position = string(mJson)
	
	mp2 := split.SplitAccount(&dish)
	mJson = convertstruct.ConvertType(mp2)

	msg.Account = string(mJson)

	//d.rabbit.Producer(msg)
	fmt.Println(msg)
}