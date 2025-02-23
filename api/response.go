package api

import (
	"fmt"
	"io"
	"net/http"
	"slices"
)

var (
	successCodes = []int{
		http.StatusOK,
		http.StatusCreated,
		http.StatusAccepted,
		http.StatusNoContent,
	}
)

type requestError error

type preRequestError struct {
	error error
}

func (p preRequestError) Error() string {
	return p.error.Error()
}

type serverResponseError struct {
	code     int
	response string
}

func (u serverResponseError) Error() string {
	return fmt.Sprintf("Server response faild with status code %d and body: %s", u.code, u.response)
}

type unauthorized struct {
	err serverResponseError
}

func (u unauthorized) Error() string {
	return u.err.Error()
}

type notfound struct {
	err serverResponseError
}

func (n notfound) Error() string {
	return n.err.Error()
}

type unknownError struct {
	err serverResponseError
}

func (u unknownError) Error() string {
	return u.err.Error()
}

type serverResponseSuccess struct {
	response *http.Response
}

func (s serverResponseSuccess) Close() {
	s.response.Body.Close()
}

func parseHttpResponse(response *http.Response) (serverResponseSuccess, requestError) {
	if slices.Contains(successCodes, response.StatusCode) {
		return serverResponseSuccess{response}, nil
	}

	var success serverResponseSuccess
	defer response.Body.Close()
	code := response.StatusCode
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	e := serverResponseError{code, string(body)}
	switch code {
	case http.StatusUnauthorized:
		return success, unauthorized{e}

	case http.StatusNotFound:
		return success, notfound{e}

	default:
		return success, unknownError{e}
	}
}

type postResponseError struct {
	error error
}

func (p postResponseError) Error() string {
	return p.error.Error()
}
