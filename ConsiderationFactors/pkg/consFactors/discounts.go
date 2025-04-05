package consfactors

import (
	"github.com/lashkapshka/go-Sber/internal/model"
	mathoperations "github.com/lashkapshka/go-Sber/pkg/mathOperations"
)

func CalculateDiscount(products []model.Products, discounts []model.Discounts) []model.Products {
	discountMap := make(map[string]int)
	for _, discount := range discounts {
		discountMap[discount.Name] = discount.Number
	}

	for i := range products {
		if discountValue, ok := discountMap[products[i].Name]; ok {
			products[i].Price = mathoperations.ConvertDiscount(
				float64(products[i].Price),
				float64(discountValue),
			)
			products[i].TotalPrice = products[i].NumberServings * products[i].Price
		}
	}

	return products
}