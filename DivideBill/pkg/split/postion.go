package split

import (
	"fmt"
	"strings"

	"github.com/lashkapashka/divideBill/internal/model"
)


func SplitPosition(name string, dishes *model.DataDishes) {
	var mp = make(map[string]float64)
	var total_price float64

	for i := 1; i <= dishes.NumberClients; i++ {
		for j := i+1; j <= dishes.NumberClients; j++ {
			
			for _, nameProduct := range dishes.Products{
				if strings.Compare(name, nameProduct.Name) == 0 {
					total_price = nameProduct.TotalPrice * (float64(i)/float64(j))
					break
				}
			}

			mp[fmt.Sprintf("%d/%d", i, j)] = total_price
		} 

	}

}