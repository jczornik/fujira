package api

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/jczornik/fujira/auth"
	"github.com/jczornik/fujira/config"
)

type sender struct {
	authenticator auth.Auth
	url           string
	client        http.Client
}

func setJsonHeader(req *http.Request) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
}

func NewSender() (sender, error) {
	auth, err := auth.NewAuth()
	if err != nil {
		return sender{}, err
	}

	url, err := config.GetConfig().GetWorkspaceURL()
	if err != nil {
		return sender{}, err
	}

	client := http.Client{
		Timeout: 20 * time.Second,
	}

	return sender{authenticator: auth, url: url, client: client}, nil
}

func logRequest(req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("Error while logging request", err)
		return
	}

	var b strings.Builder
	b.WriteString("Sending request:\n")
	b.Write(body)

	log.Println(b.String())
}

func (s sender) sendRequest(req *http.Request) (serverResponseSuccess, requestError) {
	s.authenticator.AddHeader(req)
	setJsonHeader(req)

	response, err := s.client.Do(req)
	if err != nil {
		return serverResponseSuccess{}, preRequestError{error: err}
	}

	return parseHttpResponse(response)
}

func (s sender) Get(path string) (serverResponseSuccess, requestError) {
	req, err := http.NewRequest("GET", "https://"+s.url+"/"+path, nil)
	if err != nil {
		return serverResponseSuccess{}, preRequestError{error: err}
	}

	return s.sendRequest(req)
}

func (s sender) Post(path string, body io.ReadCloser) (serverResponseSuccess, requestError) {
	req, err := http.NewRequest("POST", "https://"+s.url+"/"+path, nil)
	if err != nil {
		return serverResponseSuccess{}, preRequestError{error: err}
	}

	req.Body = body

	return s.sendRequest(req)
}
