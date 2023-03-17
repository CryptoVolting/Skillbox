package billing

import (
	"log"
	"os"
	"strconv"
)

var fileName = "../finalTemp/simulator/billing.data"

type BillingData struct {
	CreateCustomer bool `json:"create_customer"`
	Purchase       bool `json:"purchase"`
	Payout         bool `json:"payout"`
	Recurring      bool `json:"recurring"`
	FraudControl   bool `json:"fraudControl"`
	CheckoutPage   bool `json:"checkoutPage"`
}

func GetData() BillingData {
	bytesData, err := os.ReadFile(fileName)
	if err != nil {
		log.Printf(err.Error())
	}

	var stats []bool

	for i := len(bytesData) - 1; i >= 0; i-- {
		status, err := strconv.ParseBool(string(bytesData[i]))
		if err != nil {
			log.Printf(err.Error())
		}
		stats = append(stats, status)
	}

	billing := BillingData{
		stats[0],
		stats[1],
		stats[2],
		stats[3],
		stats[4],
		stats[5],
	}

	return billing
}
