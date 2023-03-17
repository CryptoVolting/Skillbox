package mms

import (
	"Skillbox/internal/codeCountry"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

func GetData() ([]MMSData, error) {
	r, err := http.Get("http://127.0.0.1:8383/mms")
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if r.StatusCode != 200 {
		log.Printf("Статус ответа: %d\nBody: %s\n", r.StatusCode, body)
		return nil, err
	}
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	var data []MMSData
	if err := json.Unmarshal(body, &data); err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	var result []MMSData
	for _, m := range data {
		checkProvider := m.Provider == "Topolo" || m.Provider == "Rond" || m.Provider == "Kildy"
		if codeCountry.IsExist(m.Country) && checkProvider {
			result = append(result, m)
		}
	}

	return result, nil
}
