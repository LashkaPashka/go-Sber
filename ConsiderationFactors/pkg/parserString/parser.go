package parserString

import (
	"encoding/json"
	"log"
)

func ParseStringMap(data string) map[string]string {
	var mp map[string]string
	
	if err := json.Unmarshal([]byte(data), &mp); err != nil {
		log.Println("не удалось преобразовать в модель")
		return nil
	}

	return mp
}


func Parser[T any](data string) *T{
	var dataModel T
	var unscapped string

	if err := json.Unmarshal([]byte(data), &unscapped); err == nil {
		data = unscapped
	}

	if err := json.Unmarshal([]byte(data), &dataModel); err != nil {
		log.Println("не удалось преобразовать в модель")
		return nil
	}

	return &dataModel
}

func ConvertJSON(data any) string {
	dataString, err := json.Marshal(&data) 
	if err != nil {
		log.Println("Не удалось преобразовать в json")
		return ""
	}

	return string(dataString)
}