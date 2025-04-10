package service

import (
	"fmt"
	"log"

	"github.com/lashkapshka/go-Sber/internal/model"
	"github.com/lashkapshka/go-Sber/internal/repository"
	"github.com/lashkapshka/go-Sber/pkg/client"
	"github.com/lashkapshka/go-Sber/pkg/parserString"
	"github.com/lashkapshka/go-Sber/pkg/consFactors"
)

type FactorsService struct {
	repo *repository.FactorsRepository
}

func New() *FactorsService{
	return &FactorsService{
		repo: repository.New(),

	}
}

func (s *FactorsService) DivideBill(mpHash string) string{
	mp := parserString.ParseStringMap(mpHash)

	factors := parserString.Parser[model.Factors](client.ClientGetRedis(":8000", "cache/get-data/factors", mp["hash"]))
	dataDishes := parserString.Parser[model.DataDishes](client.ClientGetRedis(":8000", "cache/get-data/cheque", mp["hash"]))

	switch {
		case factors != nil && factors.Discounts != nil:
			dataDishes.Products = consfactors.CalculateDiscount(dataDishes.Products, factors.Discounts)
			fallthrough
		case factors != nil && factors.Tips != nil:
			dataDishes.Total_account = consfactors.CalculateTips(dataDishes.Total_account, factors.Tips)
		default:
			log.Println("в модели factors ничего нет")
			return ""
	}

	client.ClientPost(":8000", fmt.Sprintf("cache/set-data/cheque:%s", mp["hash"]), parserString.ConvertJSON(dataDishes))
	dataDivideBill := client.ClientGetGo(":8085", "/divide-bill", mp["hash"])
	
	return dataDivideBill
}