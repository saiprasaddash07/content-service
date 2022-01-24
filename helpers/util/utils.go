package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strings"

	"github.com/saiprasaddash07/content-service.git/config"
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

func CallAPI(URL string) (string, error) {
	response := ""
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		log.Println(err.Error(), err)
		return "", err
	}

	req.Header.Set("User-Interaction-Header", config.Get().UserInteractionHeader)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error(), err)
		return "", err
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error(), err)
		return "", err
	}
	response = string(responseBody)

	return response, nil
}
