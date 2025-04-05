package parserstring

import (
	"encoding/json"
	"log"
)

func ConvertJSON(data any) string {
	mJson, err := json.Marshal(data)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return string(mJson)
}

func ParserString[T any](dataString string) *T {
	var dataModel T
	var unescaped string

	if err := json.Unmarshal([]byte(dataString), &unescaped); err == nil {
		dataString = unescaped
	}

	if err := json.Unmarshal([]byte(dataString), &dataModel); err != nil {
		panic(err)
	}

	return &dataModel
}