package consfactors

import "github.com/lashkapshka/go-Sber/internal/model"

func CalculateTips(total_account int, trips []model.Tips) int{
	for _, valTrips := range trips {
		total_account += valTrips.Number
	}

	return total_account
}