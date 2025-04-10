package service

import (
	"fmt"

	"github.com/lashkapashka/divideBill/internal/model"
	"github.com/lashkapashka/divideBill/pkg/ParserString"
	"github.com/lashkapashka/divideBill/pkg/client"
	"github.com/lashkapashka/divideBill/pkg/split"
)



func Divide(req map[string]string) string{	
	resp := client.Client(fmt.Sprintf("http://localhost:8000/cache/get-data/cheque:%s", req["hash"]))
	
	dish := parserstring.ParserString[model.DataDishes](resp)

	mapPosition := split.SplitPosition([]string{"Espresso"}, dish, req)
	mapAccount := split.SplitAccount(dish, req)

	msg := model.Response{
		Position: mapPosition,
		Account: mapAccount,
	}
	
	msgString := parserstring.ConvertJSON(msg)
	//d.rabbit.Producer(msg)
	return msgString
}