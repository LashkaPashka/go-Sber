package service

import (
	"fmt"

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

func (d *DivideService) Divide(req model.RequestBody) string{	
	resp := client.Client(fmt.Sprintf("http://localhost:8000/cache/get-data/cheque:%s", req.Hash))
	
	dish := parserstring.ParserString[model.DataDishes](resp)

	mapPosition := split.SplitPosition([]string{"Espresso"}, dish, &req)
	mapAccount := split.SplitAccount(dish, &req)

	msg := model.Response{
		Position: mapPosition,
		Account: mapAccount,
	}
	
	msgString := parserstring.ConvertJSON(msg)

	//d.rabbit.Producer(msg)
	return msgString
}