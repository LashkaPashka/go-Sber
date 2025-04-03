package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/lashkapshka/go-Sber/internal/db"
	"github.com/lashkapshka/go-Sber/internal/model"
)

type FactorsRepository struct {
	Client *db.DbRedis
}

func New() *FactorsRepository {
	return &FactorsRepository{
		Client: db.New(),
	}
}

func GetData[T any](r *FactorsRepository, keyValue, nameID string) []T {
	var d []T
	cmd := r.Client.Pool.Get(context.Background(), fmt.Sprintf("%s:", keyValue) + nameID)

	if err := json.Unmarshal([]byte(cmd.Val()), &d); err != nil {
		return nil
	}

	return d
}

func (r *FactorsRepository) GetFactors(nameID string) (*model.Factors) {
	var rangeFact = []string{"discounts", "tips", "promtions"}
	var factors model.Factors

	for _, val := range rangeFact {
		switch val {
			case "discounts":
				d := GetData[model.Discounts](r, val, nameID)
				factors.Discounts = d
			case "tips":
				d := GetData[model.Tips](r, val, nameID)
				factors.Tips = d
			case "promtions":
				d := GetData[model.Promtions](r, val, nameID)
				factors.Promtions = d
		}
	}

	return &factors
}

func (r *FactorsRepository) ModifyFactors(nameID string, dataDishes model.DataDishes) error {
	key := "modify:" + nameID

	val, err := json.Marshal(dataDishes)
	if err != nil {
		return err
	}

	err = r.Client.Pool.Set(context.Background(), key, string(val), time.Minute*5).Err()
	if err != nil {
		return err
	}

	return nil
}