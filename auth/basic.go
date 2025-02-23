package auth

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

type BasicAuth struct {
	basic string
}

func NewBasicAuth(email, token string) BasicAuth {
	input := fmt.Sprintf("%s:%s", email, token)
	basic := base64.StdEncoding.EncodeToString([]byte(input))
	return BasicAuth{basic: "Basic " + basic}
}

func (a BasicAuth) AddHeader(r *http.Request) {
	r.Header.Add("Authorization", a.basic)
}
