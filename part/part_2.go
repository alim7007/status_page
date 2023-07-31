package part

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

var countries map[string]string = map[string]string{
	"RU": "russia",
	"US": "unated state america",
	"GB": "grate britain",
	"FR": "france",
	"BL": "belgue",
	"AT": "austria",
	"BG": "bolgaria",
	"DK": "dania",
	"CA": "canada",
	"ES": "spain",
	"CH": "chehia",
	"TR": "turky",
	"PE": "peru",
	"NZ": "new ziland",
	"MC": "monaco",
}

var provaiders map[string]struct{} = map[string]struct{}{
	"Topolo": struct{}{},
	"Rond":   struct{}{},
	"Kildy":  struct{}{},
}

type SMSData struct {
	Country      string `json:"country,omitempty"`
	Bandwidth    string `json:"bandwidth,omitempty"`
	ResponseTime string `json:"response_time,omitempty"`
	Provider     string `json:"provider,omitempty"`
}

func SmsDataInCsv() []SMSData {
	data := []SMSData{}
	csvFile, err := os.Open("sms.data")
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
		_, ok := provaiders[record[3]]
		if ok == false {
			continue
		}
		_, ok = countries[record[0]]
		if ok == false {
			continue
		}

		modul := SMSData{
			Country:      record[0],
			Bandwidth:    record[1],
			ResponseTime: record[2],
			Provider:     record[3],
		}
		data = append(data, modul)
	}

	return data
}
