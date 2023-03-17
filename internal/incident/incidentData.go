package incident

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"`
}

func GetData() ([]IncidentData, error) {
	var data []IncidentData

	r, err := http.Get("http://127.0.0.1:8383/accendent")
	if err != nil {
		log.Printf(err.Error())
		return data, err
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if r.StatusCode != 200 {
		log.Printf("Статус ответа: %d\nBody: %s\n", r.StatusCode, body)
		return data, err
	}
	if err != nil {
		log.Printf(err.Error())
		return data, err
	}

	if err := json.Unmarshal(body, &data); err != nil {
		log.Printf(err.Error())
		return data, err
	}

	return data, nil
}
