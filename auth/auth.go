package auth

import (
	"errors"
	"net/http"

	"github.com/jczornik/fujira/config"
)

type Auth interface {
	AddHeader(*http.Request)
}

func NewAuth() (Auth, error) {
	conf := config.GetConfig()
	if conf == nil {
		panic("Error while loading conf")
	}

	if basic, err := conf.GetBasicAuth(); err == nil {
		return NewBasicAuth(basic.Email, basic.Token), nil
	}

	return nil, errors.New("Cannot construct auth from config")
}
