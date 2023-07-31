package part

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

var provaidersEmail map[string]struct{} = map[string]struct{}{
	"Gmail":      struct{}{},
	"Yahoo":      struct{}{},
	"Hotmail":    struct{}{},
	"MSN":        struct{}{},
	"Orange":     struct{}{},
	"Comcast":    struct{}{},
	"AOL":        struct{}{},
	"Live":       struct{}{},
	"RediffMail": struct{}{},
	"GMX":        struct{}{},
	"Protonmail": struct{}{},
	"Yandex":     struct{}{},
	"Mail.ru":    struct{}{},
}

type EmailData struct {
	Country      string `json:"country,omitempty"`
	Provider     string `json:"provider,omitempty"`
	DeliveryTime int    `json:"delivery_time,omitempty"`
}

func EmailDataInCsv() []EmailData {
	data := []EmailData{}
	csvFile, err := os.Open("email.data")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(csvFile)
	r.Comma = ';'
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		_, ok := provaidersEmail[record[1]]
		if ok == false {
			continue
		}
		_, ok = countries[record[0]]
		if ok == false {
			continue
		}

		deliveryTime, err := strconv.Atoi(record[2])
		if err != nil {
			continue
		}

		modul := EmailData{
			Country:      record[0],
			DeliveryTime: deliveryTime,
			Provider:     record[1],
		}
		data = append(data, modul)
	}

	return data
}
