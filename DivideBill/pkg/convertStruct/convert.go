package convertstruct

import (
	"encoding/json"
	"log"
)

func ConvertType(data any) string {
	mJson, err := json.Marshal(data)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return string(mJson)
}