package part

import (
	"sort"
	"strings"
)

func GetResultData() ResultSetT {
	resultSetT := ResultSetT{}
	smsData := SmsDataInCsv()
	for _, val := range smsData {
		val.Country = countries[val.Country]
	}
	sms1 := smsData
	sort.Slice(sms1, func(i, j int) bool {
		a := strings.Compare(sms1[i].Provider, sms1[j].Provider)
		return a == -1
	})
	sms2 := smsData
	sort.Slice(sms2, func(i, j int) bool {
		a := strings.Compare(sms2[i].Country, sms2[j].Country)
		return a == -1
	})
	sms := [][]SMSData{sms1, sms2}
	resultSetT.SMS = sms

	mmsData := GetMMSData()
	for _, val := range mmsData {
		val.Country = countries[val.Country]
	}
	mms1 := mmsData
	sort.Slice(mms1, func(i, j int) bool {
		a := strings.Compare(mms1[i].Provider, mms1[j].Provider)
		return a == -1
	})
	mms2 := mmsData
	sort.Slice(mms2, func(i, j int) bool {
		a := strings.Compare(mms2[i].Country, mms2[j].Country)
		return a == -1
	})
	mms := [][]MMSData{mms1, mms2}
	resultSetT.MMS = mms

	voice := VoiceCallDataInCsv()
	resultSetT.VoiceCall = voice

	email := EmailDataInCsv()
	emailMap := make(map[string][]EmailData)
	for _, val := range email {
		v := emailMap[val.Country]
		v = append(v, val)
		emailMap[val.Country] = v
	}

	emailMapNew := make(map[string][][]EmailData)
	for key, val := range emailMap {

		sort.Slice(val, func(i, j int) bool {
			return val[i].DeliveryTime < val[j].DeliveryTime
		})

		email1 := []EmailData{}
		email2 := []EmailData{}
		for i := 0; i < 3; i++ {
			email1 = append(email1, val[i])
		}
		for i := len(val) - 1; i >= len(val)-3; i-- {
			email2 = append(email2, val[i])
		}
		emailMapNew[key] = [][]EmailData{email1, email2}
	}
	emailRes := [][]EmailData{}
	for _, val := range emailMapNew {
		emailRes = append(emailRes, val...)
	}
	resultSetT.Email = emailRes

	billing := BillingDataInCsv()
	resultSetT.Billing = billing

	support := GetSupportData()
	count := 0
	for _, val := range support {
		count += val.ActiveTickets
	}
	var supportArr []int
	time := count * (60 / 18)
	rang := 0
	if count < 9 {
		rang = 1
	} else if count > 16 {
		rang = 3
	} else {
		rang = 2
	}

	supportArr = append(supportArr, time)
	supportArr = append(supportArr, rang)

	resultSetT.Support = supportArr

	incident := GetIncidentData()
	sort.Slice(incident, func(i, j int) bool {
		if incident[i].Status == "active" && incident[j].Status != "active" {
			return true
		}
		return false
	})
	resultSetT.Incidents = incident
	return resultSetT
}
