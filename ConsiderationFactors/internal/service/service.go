package service

import (
	"log"

	"github.com/lashkapshka/go-Sber/internal/model"
	"github.com/lashkapshka/go-Sber/internal/repository"
	"github.com/lashkapshka/go-Sber/pkg/client"
	consfactors "github.com/lashkapshka/go-Sber/pkg/consFactors"
	"github.com/lashkapshka/go-Sber/pkg/parserString"
)

type FactorsService struct {
	repo *repository.FactorsRepository
}

func New() *FactorsService{
	return &FactorsService{
		repo: repository.New(),
	}
}

func (s *FactorsService) DivideBill(msg string) *model.DataDishes{
	//nameID := strings.Split(msg, " ")
	
	factors := parserString.Parser[model.Factors](client.ClientGet("factors"))

	dataDishes := parserString.Parser[model.DataDishes](client.ClientGet("key"))

	switch {
		case factors != nil && factors.Discounts != nil:
			dataDishes.Products = consfactors.CalculateDiscount(dataDishes.Products, factors.Discounts)
			fallthrough
		case factors != nil && factors.Tips != nil:
			dataDishes.Total_account = consfactors.CalculateTips(dataDishes.Total_account, factors.Tips)
		default:
			log.Println("в модели factors ничего нет")
			return nil
	}

	client.ClientPost("key", parserString.ConvertJSON(dataDishes))
	log.Println("данные обновлены в Redis")

	return dataDishes
}