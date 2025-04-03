package service

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/lashkapshka/go-Sber/internal/model"
	"github.com/lashkapshka/go-Sber/internal/repository"
	"github.com/lashkapshka/go-Sber/pkg/client"
	consfactors "github.com/lashkapshka/go-Sber/pkg/consFactors"
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
	nameID := strings.Split(msg, " ")
	
	factors := s.repo.GetFactors(nameID[1])
	
	cData := client.Client()

	var dataDishes model.DataDishes

	if err := json.Unmarshal([]byte(cData), &dataDishes); err != nil {
		log.Fatal("ошибка при конвертации строки в структуру Products")
		return nil
	}
	
	switch {
		case factors != nil && factors.Discounts != nil:
			dataDishes.Products = consfactors.CalculateDiscount(dataDishes.Products, factors.Discounts)
			
		case factors != nil && factors.Tips != nil:
			dataDishes.Total_account = consfactors.CalculateTips(dataDishes.Total_account, factors.Tips)

		default:
			log.Println("в модели factors ничего нет")
			return nil
	}

	if err := s.repo.ModifyFactors(nameID[1], dataDishes); err != nil {
		log.Fatal("не удалось модифицировать данные в Redis")
		return nil
	}

	return &dataDishes
}