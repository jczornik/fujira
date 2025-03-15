package api

import (
	"bytes"
	"encoding/json"
	"io"
)

const (
	searchURL  = "rest/api/3/search/jql"
	maxResults = 500
)

type SearchQuery struct {
	Expand          string   `json:"expand"`
	Fields          []string `json:"fields"`
	FieldsByKeys    bool     `json:"fieldsByKeys"`
	Jql             string   `json:"jql"`
	MaxResults      int      `json:"maxResults"`
	NextPageToken   *string  `json:"nextPageToken"`
	Properties      []string `json:"properties"`
	ReconcileIssues []string `json:"reconcileIssues"`
}

func SearchIssues(fields []string, jql string) (string, error) {
	sender, err := NewSender()
	if err != nil {
		return "", err
	}

	sq := SearchQuery{
		Expand:          "names",
		Fields:          fields,
		FieldsByKeys:    false,
		Jql:             jql,
		MaxResults:      maxResults,
		Properties:      []string{},
		ReconcileIssues: []string{},
	}

	payload, err := json.Marshal(sq)
	if err != nil {
		return "", err
	}

	body := io.NopCloser(bytes.NewReader(payload))
	defer body.Close()

	response, err := sender.Post(searchURL, body)
	if err != nil {
		return "", err
	}
	defer response.Close()

	data, err := io.ReadAll(response.response.Body)
	if err != nil {
		return "", err
	}

	return string(data), err
}
