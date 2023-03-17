package email

import (
	"Skillbox/internal/codeCountry"
	"log"
	"os"
	"strconv"
	"strings"
)

var file = "../finalTemp/simulator/email.data"

var correctProviders = []string{
	"Gmail",
	"Yahoo",
	"Hotmail",
	"MSN",
	"Orange",
	"Comcast",
	"AOL",
	"Live",
	"RediffMail",
	"GMX",
	"Proton Mail",
	"Yandex",
	"Mail.ru",
}

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"deliveryTime"`
}

func GetData() ([]EmailData, error) {
	bytesData, err := os.ReadFile(file)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	var data []EmailData

	dataSplit := strings.Split(string(bytesData), "\n")
	for _, email := range dataSplit {
		em := strings.Split(email, ";")
		if len(em) == 3 && codeCountry.IsExist(em[0]) {
			if checkProvider(em[1]) {
				time, err := strconv.Atoi(em[2])
				if err != nil {
					log.Printf(err.Error())
					return nil, err
				}
				e := EmailData{
					em[0],
					em[1],
					time,
				}
				data = append(data, e)
			}
		}
	}

	return data, nil
}

func checkProvider(provider string) bool {
	for _, p := range correctProviders {
		if p == provider {
			return true
		}
	}
	return false
}
