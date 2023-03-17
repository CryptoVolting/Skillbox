package sms

import (
	"Skillbox/internal/codeCountry"
	"log"
	"os"
	"strings"
)

var file = "../finalTemp/simulator/sms.data"

type SMSData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"responseTime"`
	Provider     string `json:"provider"`
}

func GetData() ([]SMSData, error) {
	bytesSMSData, err := os.ReadFile(file)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	var data []SMSData

	dataSplit := strings.Split(string(bytesSMSData), "\n")
	for _, sms := range dataSplit {
		s := strings.Split(sms, ";")
		if len(s) == 4 && codeCountry.IsExist(s[0]) {
			checkProvider := s[3] == "Topolo" || s[3] == "Rond" || s[3] == "Kildy"
			if checkProvider {
				newSMS := SMSData{
					s[0],
					s[1],
					s[2],
					s[3],
				}
				data = append(data, newSMS)
			}
		}
	}

	return data, nil
}
