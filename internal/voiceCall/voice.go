package voiceCall

import (
	"Skillbox/internal/codeCountry"
	"log"
	"os"
	"strconv"
	"strings"
)

var fileName = "../finalTemp/simulator/voice.data"

type VoiceData struct {
	Country             string  `json:"country"`
	Bandwidth           string  `json:"bandwidth"`
	ResponseTime        string  `json:"responseTime"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connectionStability"`
	TTFB                string  `json:"ttfb"`
	VoicePurity         string  `json:"voicePurity"`
	MedianOfCallsTime   string  `json:"medianOfCallsTime"`
}

func GetData() ([]VoiceData, error) {
	bytesData, err := os.ReadFile(fileName)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	var vData []VoiceData

	dataSplit := strings.Split(string(bytesData), "\n")
	for _, call := range dataSplit {
		cSplit := strings.Split(call, ";")
		if len(cSplit) == 8 && codeCountry.IsExist(cSplit[0]) {
			checkProvider := cSplit[3] == "TransparentCalls" || cSplit[3] == "E-Voice" || cSplit[3] == "JustPhone"
			if checkProvider {
				stab, err := strconv.ParseFloat(cSplit[4], 32)
				if err != nil {
					log.Printf(err.Error())
					return nil, err
				}
				stability := float32(stab)
				voiceCall := VoiceData{
					cSplit[0],
					cSplit[1],
					cSplit[2],
					cSplit[3],
					stability,
					cSplit[5],
					cSplit[6],
					cSplit[7],
				}
				vData = append(vData, voiceCall)
			}
		}
	}

	return vData, nil
}
