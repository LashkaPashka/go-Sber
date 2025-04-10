package split

import (
	"fmt"
	"strconv"

	"github.com/lashkapashka/divideBill/internal/model"
)


func SplitAccount(dish model.DataDishes, req map[string]string) map[string]int {
	numberClients, _ := strconv.Atoi(req["numClients"])
	useClients, _ := strconv.Atoi(req["useClients"])
	
	var mp = make(map[string]int)

	total_price := dish.Total_account * (float64(numberClients)/float64(useClients))

	mp[fmt.Sprintf("%d/%d", numberClients, useClients)] = int(total_price)


	return mp
}