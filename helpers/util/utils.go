package util

import (
	"encoding/json"
	"math"
	"strings"
)

func StructToJSON(val interface{}) interface{} {
	jsonEncoded, _ := json.Marshal(val)
	var respJSON interface{}
	json.Unmarshal([]byte(jsonEncoded), &respJSON)
	return respJSON
}

func Contains(str []string, key string) bool {
	for _, v := range str {
		if v == key {
			return true
		}
	}
	return false
}

func IsInteger(val float64) bool {
	return math.Floor(val) == math.Ceil(val)
}

func ValidateCSV(val string) bool {
	return strings.Split(val, ".")[1] == "csv"
}