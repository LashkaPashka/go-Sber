package split

import (
	"fmt"

	"github.com/lashkapashka/divideBill/internal/model"
)

func SplitPosition(names []string, dishes *model.DataDishes, req *model.RequestBody) map[string]int {
	var mp = make(map[string]int)

	// Кэшируем цены продуктов по имени
	priceMap := make(map[string]float64)
	for _, p := range dishes.Products {
		priceMap[p.Name] = p.TotalPrice
	}

	// Предвычисляем сумму
	totalSum := 0.0
	for _, name := range names {
		if price, ok := priceMap[name]; ok {
			totalSum += price
		}
	}

	// Основные циклы
	for i := 1; i <= req.UseClients; i++ {
		for j := i + 1; j <= req.NumberClients; j++ {
			totalPrice := totalSum * (float64(i) / float64(j))
			key := fmt.Sprintf("%d/%d", i, j)
			mp[key] = int(totalPrice)
		}
	}

	return mp
}