package ha

import (
	"bytes"
	"crypto/tls"
	"io"
	"log"
	"net/http"
)

type Client struct {
	accessToken string
	client      *http.Client
}

func NewClient(accessToken string) *Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	return &Client{client: client, accessToken: accessToken}
}

func (ha Client) FetchProblemData(ch ChallengeName) []byte {
	problemUrl := "https://hackattic.com/challenges/" + string(ch) + "/problem?access_token=" + ha.accessToken

	response, err := ha.client.Get(problemUrl)
	if err != nil {
		log.Fatalf("Couldn't fetch problem data: %v", err)
	}
	defer func(r io.ReadCloser) {
		err = r.Close()
		if err != nil {
			log.Println(err)
		}
	}(response.Body)

	data, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		log.Fatalf("Couldn't parse problem data: %v", err)
	}

	return data
}

func (ha Client) Submit(ch ChallengeName, solution []byte) {
	solutionUrl := "https://hackattic.com/challenges/" + string(ch) + "/solve?access_token=" + ha.accessToken

	response, err := ha.client.Post(solutionUrl, "application/json", bytes.NewBuffer(solution))
	if err != nil {
		log.Fatalf("Couldn't Submit problem: %v", err)
	}
	defer func(r io.ReadCloser) {
		err = r.Close()
		if err != nil {
			log.Println(err)
		}
	}(response.Body)

	log.Printf("%s", response.Status)
}
