package consfactors

import "github.com/lashkapshka/go-Sber/internal/model"

func CalculateTips(total_account int, tips []model.Tips) int{
	for _, valTrips := range tips {
		total_account += valTrips.Number
	}

	return total_account
}