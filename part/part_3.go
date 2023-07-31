package part

import (
	"encoding/json"
	"log"
	"net/http"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

func GetMMSData() []MMSData {
	resp, err := http.Get("http://127.0.0.1:8383/mms")
	if err != nil {
		log.Fatal(err)
	}
	data := []MMSData{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	return data

}
