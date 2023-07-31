package part

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

var provider map[string]struct{} = map[string]struct{}{
	"TransparentCalls": struct{}{},
	"E-Voice":          struct{}{},
	"JustPhone":        struct{}{},
}

type VoiceCallData struct {
	Country             string  `json:"country,omitempty"`
	Bandwidth           string  `json:"bandwidth,omitempty"`
	ResponseTime        string  `json:"response_time,omitempty"`
	Provider            string  `json:"provider,omitempty"`
	ConnectionStability float32 `json:"connection_stability,omitempty"`
	TTFB                int     `json:"ttfb,omitempty"`
	VoicePurity         int     `json:"voice_purity,omitempty"`
	MedianOfCallsTime   int     `json:"median_of_call_time,omitempty"`
}

func VoiceCallDataInCsv() []VoiceCallData {
	data := []VoiceCallData{}
	csvFile, err := os.Open("voice.data")
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
		_, ok := provider[record[3]]
		if ok == false {
			continue
		}
		_, ok = countries[record[0]]
		if ok == false {
			continue
		}
		con, err := strconv.ParseFloat(record[4], 32)
		if err != nil {
			continue
		}

		ttfb, err := strconv.Atoi(record[5])
		if err != nil {
			continue
		}
		voicePurity, err := strconv.Atoi(record[6])
		if err != nil {
			continue
		}
		median, err := strconv.Atoi(record[7])
		if err != nil {
			continue
		}

		modul := VoiceCallData{
			Country:             record[0],
			Bandwidth:           record[1],
			ResponseTime:        record[2],
			Provider:            record[3],
			ConnectionStability: float32(con),
			TTFB:                ttfb,
			VoicePurity:         voicePurity,
			MedianOfCallsTime:   median,
		}
		data = append(data, modul)
	}

	return data
}
