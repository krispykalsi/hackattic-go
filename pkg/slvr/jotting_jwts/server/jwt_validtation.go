package server

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"log"
	"time"
)

func validateJwt(jwt []byte, secret string) (bool, *jwtPayload) {
	if !verifySignature(jwt, secret) {
		return false, nil
	}
	p := getPayloadFromJwt(jwt)
	if isValidByDate(p) {
		return true, p
	} else {
		return false, nil
	}
}

func isValidByDate(p *jwtPayload) bool {
	tNow := time.Now()
	if p.Exp != 0 {
		tExp := time.Unix(p.Exp, 0)
		if tNow.After(tExp) {
			return false
		}
	}
	if p.Nbf != 0 {
		tNbf := time.Unix(p.Nbf, 0)
		if tNow.Before(tNbf) {
			return false
		}
	}
	return true
}

func verifySignature(jwt []byte, secret string) bool {
	jwtParts := bytes.Split(jwt, []byte("."))
	headerAndPayload := bytes.Join(jwtParts[:2], []byte("."))
	hm := hmac.New(sha256.New, []byte(secret))
	hm.Write(headerAndPayload)
	actualSignature := hm.Sum(nil)
	expectedSignature := decodeBase64(jwtParts[2])
	return hmac.Equal(actualSignature, expectedSignature)
}

func getPayloadFromJwt(jwt []byte) *jwtPayload {
	encodedPayload := bytes.Split(jwt, []byte("."))[1]
	decodedPayload := decodeBase64(encodedPayload)
	payload := jwtPayload{}
	err := json.Unmarshal(decodedPayload, &payload)
	if err != nil {
		log.Fatalf("Couldn't parse from json: %v", err)
	}
	return &payload
}
