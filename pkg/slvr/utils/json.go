package utils

import (
	"encoding/json"
	"log"
)

func FromJson(data []byte, problem interface{}) {
	if data == nil {
		log.Fatalln("Data is empty")
	}
	err := json.Unmarshal(data, problem)
	if err != nil {
		log.Fatalf("Couldn't parse from json: %v", err)
	}
}

func ToJson(sol interface{}) []byte {
	jsonBytes, err := json.Marshal(sol)
	if err != nil {
		log.Fatalf("Couldn't parse to json: %v", err)
	}
	return jsonBytes
}
