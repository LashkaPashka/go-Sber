package consfactors

import (
	"strings"

	"github.com/lashkapshka/go-Sber/internal/model"
	mathoperations "github.com/lashkapshka/go-Sber/pkg/mathOperations"
)

func CalculateDiscount(products []model.Products, discounts []model.Discounts) []model.Products {
	for _, valProduct := range products {
		for _, valDiscount := range discounts {
			if strings.Compare(valProduct.Name, valDiscount.Name) == 0 {
				valProduct.Price = mathoperations.ConvertDiscount(
					float64(valProduct.Price),
					valDiscount.Number,
				)
				valProduct.TotalPrice = valProduct.NumberServings * valProduct.Price
			}
		}
	}

	return products
}