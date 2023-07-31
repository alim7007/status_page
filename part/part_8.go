package part

import (
	"encoding/json"
	"log"
	"net/http"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы active и closed
}

func GetIncidentData() []IncidentData {
	resp, err := http.Get("http://127.0.0.1:8383/accendent")
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode == 500 {
		return []IncidentData{}
	}

	data := []IncidentData{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
