package service

import (
	"fmt"

	"github.com/lashkapashka/divideBill/internal/model"
	"github.com/lashkapashka/divideBill/pkg/Serializer"
	"github.com/lashkapashka/divideBill/pkg/client"
	"github.com/lashkapashka/divideBill/pkg/split"
)

type DivideService struct {

}

func NewDivideService() *DivideService {
	return &DivideService{}
}

func (s DivideService) Divide(req map[string]string) string{	
	resp := client.Client(fmt.Sprintf("http://localhost:8000/cache/get-data/cheque:%s", req["hash"]))
	
	dish := Serializer.Deserialize[model.DataDishes](resp)

	mapPosition := split.SplitPosition(req["position"], dish, req)
	mapAccount := split.SplitAccount(dish, req)

	msg := model.Response{
		Position: mapPosition,
		Account: mapAccount,
	}
	
	msgString := Serializer.Serialize(msg)
	
	return msgString
}