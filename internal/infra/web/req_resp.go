package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Request struct {
	PathParams  map[string]string
	QueryParams map[string]string
	Body        map[string]any
}

type Response struct {
	StatusCode int
	Body       any
}

func createRequest(r *http.Request) (*Request, error) {
	var body map[string]any
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		return nil, err
	}

	return &Request{Body: body}, nil
}

func createResponse(body any) *Response {
	return &Response{Body: body, StatusCode: http.StatusOK}
}
