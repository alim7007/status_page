package part

import (
	"encoding/json"
	"log"
	"net/http"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

func GetSupportData() []SupportData {
	resp, err := http.Get("http://127.0.0.1:8383/support")
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode == 500 {
		return []SupportData{}
	}

	data := []SupportData{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
