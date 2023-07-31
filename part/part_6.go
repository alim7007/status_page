package part

import (
	"bufio"
	"log"
	"math"
	"os"
)

func binaryToDecemal(str string) int {
	answer := 0
	var x float64
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == 48 {
			x++
			continue
		}
		answer += int(math.Pow(2, x))
		x++
	}
	return answer
}

type BillingData struct {
	CreateCustomer bool
	Purchase       bool
	Payout         bool
	Recurring      bool
	FraudControl   bool
	CheckoutPage   bool
}

func BillingDataInCsv() BillingData {
	data := []BillingData{}
	csvFile, err := os.Open("billing.data")
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewScanner(csvFile)

	for r.Scan() {
		str := r.Text()
		val := binaryToDecemal(str)

		log.Println(val)

		model := BillingData{}

		if str[0] == 49 {
			model.CheckoutPage = true
		}
		if str[1] == 49 {
			model.Purchase = true
		}
		if str[2] == 49 {
			model.Payout = true
		}
		if str[3] == 49 {
			model.Recurring = true
		}
		if str[4] == 49 {
			model.FraudControl = true
		}
		if str[5] == 49 {
			model.CheckoutPage = true
		}

		data = append(data, model)
	}

	return data[0]
}
