package server

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func appendStringToJson(as string) []byte {
	s := struct {
		Solution string `json:"solution"`
	}{Solution: as}
	sJson, err := json.Marshal(s)
	if err != nil {
		log.Fatalf("Couldn't parse to json: %v", err)
	}
	return sJson
}

func submitSolution(appendString string, w http.ResponseWriter) {
	log.Printf("Submitting \"%s\"...", appendString)
	solJson := appendStringToJson(appendString)
	_, err := w.Write(solJson)
	if err != nil {
		log.Fatalf("Couldn't write to response body: %v", err)
	}
}

func readBody(r io.ReadCloser) []byte {
	data, err := io.ReadAll(r)
	defer func(r io.ReadCloser) {
		err := r.Close()
		if err != nil {
			log.Fatalf("Couldn't close req body: %v", err)
		}
	}(r)
	if err != nil {
		log.Fatalf("Couldn't read request body: %v", err)
	}
	return data
}

func decodeBase64(encoded []byte) []byte {
	decoded, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(string(encoded))
	if err != nil {
		log.Fatalf("Couldn't decode base64 string: %v", encoded)
	}
	return decoded
}
