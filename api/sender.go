package api

import (
	"net/http"
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
		Timeout: 2 * time.Second,
	}

	return sender{authenticator: auth, url: url, client: client}, nil
}

func (s sender) Get(path string) (serverResponseSuccess, requestError) {
	req, err := http.NewRequest("GET", "https://"+s.url+"/"+path, nil)
	if err != nil {
		return serverResponseSuccess{}, preRequestError{error: err}
	}

	s.authenticator.AddHeader(req)
	setJsonHeader(req)

	response, err := s.client.Do(req)
	if err != nil {
		return serverResponseSuccess{}, preRequestError{error: err}
	}

	return parseHttpResponse(response)
}
