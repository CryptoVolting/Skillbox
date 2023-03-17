package codeCountry

import (
	"encoding/csv"
	"fmt"
	"os"
)

var (
	codeCountryMap map[string]string
	file           = "../internal/codeCountry/codes.scv"
)

func init() {
	file, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	states, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	codeCountryMap = make(map[string]string)

	for _, st := range states {
		codeCountryMap[st[0]] = st[1]
	}
}

func IsExist(code string) bool {
	if codeCountryMap[code] == "" {
		return false
	}
	return true
}
