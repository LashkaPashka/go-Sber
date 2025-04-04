package split

import (
	"fmt"

	"github.com/lashkapashka/divideBill/internal/model"
)


func SplitAccount(dish *model.DataDishes) map[string]int {
	var mp = make(map[string]int)
	var total_price float64

	for i := 1; i <= dish.NumberClients; i++ {
		for j := i+1; j <= dish.NumberClients; j++ {
			total_price = dish.Total_account * (float64(i)/float64(j))

			mp[fmt.Sprintf("%d/%d", i, j)] = int(total_price)
		}
	}

	return mp
}