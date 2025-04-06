package service

import (
	"github.com/lashkapashka/divideBill/internal/model"
	"github.com/lashkapashka/divideBill/pkg/ParserString"
	"github.com/lashkapashka/divideBill/pkg/client"
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

func (d *DivideService) Divide() string{	
	resp := client.Client("http://localhost:8000/cache/get-data/key")
	
	dish := parserstring.ParserString[model.DataDishes](resp)

	mapPosition := split.SplitPosition([]string{"Espresso"}, dish)
	mapAccount := split.SplitAccount(dish)

	msg := model.Response{
		Position: mapPosition,
		Account: mapAccount,
	}
	
	msgString := parserstring.ConvertJSON(msg)

	//d.rabbit.Producer(msg)
	return msgString
}