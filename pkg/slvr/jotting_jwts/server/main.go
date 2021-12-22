package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var ip = os.Getenv("LOCAL_IP")
var port = os.Getenv("PORT")
var addr = fmt.Sprintf("%s:%s", ip, port)

type jwtValidator struct {
	jwtSecret string
}

func NewJwtValidator(secret string) *jwtValidator {
	log.Printf("JWT Secret: %s", secret)
	return &jwtValidator{jwtSecret: secret}
}

func (jv jwtValidator) Run() {
	m := http.NewServeMux()
	m.HandleFunc("/", handleJwts(jv.jwtSecret))

	log.Printf("Listening on %s", addr)
	err := http.ListenAndServe(addr, m)
	if err != nil {
		log.Fatalf("Couldn't start server: %v", err)
	}
}

func handleJwts(secret string) func(http.ResponseWriter, *http.Request) {
	finalAppendString := ""
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			log.Println(readBody(req.Body))
			return
		}
		jwt := readBody(req.Body)
		ok, p := validateJwt(jwt, secret)
		if !ok {
			log.Printf("Token invalid: %s", string(jwt))
			return
		} else {
			log.Printf("Token ok: %s", string(jwt))
		}
		log.Printf("Appending: %s", p.Append)
		if p.Append == "" {
			submitSolution(finalAppendString, w)
		} else {
			finalAppendString += p.Append
		}
	}
}
