package api

import (
	"fmt"
	"io"
)

const (
	preferencesURL = "/rest/api/3/mypreferences"
)

func MyPreferences(key string) (string, requestError) {
	sender, err := NewSender()
	if err != nil {
		return "", err
	}

	path := fmt.Sprintf("%s?key=%s", preferencesURL, key)
	success, err := sender.Get(path)
	if err != nil {
		return "", err
	}

	defer success.Close()
	body, err := io.ReadAll(success.response.Body)
	if err != nil {
		return "", postResponseError{err}
	}

	return string(body), nil
}
