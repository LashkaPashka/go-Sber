package split

import (
	"fmt"
	"strconv"

	"github.com/lashkapashka/divideBill/internal/model"
)

func SplitPosition(name string, dishes model.DataDishes, req map[string]string) map[string]int {
	numberClients, _ := strconv.Atoi(req["numClients"])
	useClients, _ := strconv.Atoi(req["useClients"])
	
	var mp = make(map[string]int)

	// Кэшируем цены продуктов по имени
	priceMap := make(map[string]float64)
	for _, p := range dishes.Products {
		priceMap[p.Name] = p.TotalPrice
	}
	// Предвычисляем сумму
	totalSum := 0.0
	if price, ok := priceMap[name]; ok {
		totalSum = price
	}
	
	totalPrice := totalSum * (float64(numberClients) / float64(useClients))
	key := fmt.Sprintf("%d/%d", numberClients, useClients)
	mp[key] = int(totalPrice)

	return mp
}