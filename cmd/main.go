package main

import (
	"Skillbox/internal/billing"
	"Skillbox/internal/email"
	"Skillbox/internal/incident"
	"Skillbox/internal/mms"
	"Skillbox/internal/sms"
	"Skillbox/internal/support"
	"Skillbox/internal/voiceCall"
	"fmt"
)

func main() {
	fmt.Println("Сбор данных о системе SMS:")
	resultSms, _ := sms.GetData()
	fmt.Println(resultSms)

	fmt.Println("Сбор данных о системе MMS:")
	resultMMS, _ := mms.GetData()
	fmt.Println(resultMMS)

	fmt.Println("Сбор данных о системе Voice Call:")
	resultVoice, _ := voiceCall.GetData()
	fmt.Println(resultVoice)

	fmt.Println("Сбор данных о системе Email:")
	resultEmail, _ := email.GetData()
	fmt.Println(resultEmail)

	fmt.Println("Сбор данных о системе Billing:")
	fmt.Println(billing.GetData())

	fmt.Println("Сбор данных о системе Support:")
	resultSup, _ := support.GetData()
	fmt.Println(resultSup)

	fmt.Println("Сбор данных о системе инцидентов:")
	resultInc, _ := incident.GetData()
	fmt.Println(resultInc)
}
