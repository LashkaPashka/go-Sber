package split

import (
	"fmt"

	"github.com/lashkapashka/divideBill/internal/model"
)


func SplitAccount(dish *model.DataDishes, req *model.RequestBody) map[string]int {
	var mp = make(map[string]int)
	var total_price float64

	for i := 1; i <= req.UseClients; i++ {
		for j := i+1; j <= req.NumberClients; j++ {
			total_price = dish.Total_account * (float64(i)/float64(j))

			mp[fmt.Sprintf("%d/%d", i, j)] = int(total_price)
		}
	}

	return mp
}