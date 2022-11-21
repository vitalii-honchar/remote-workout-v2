package request

import "net/http"

type Response struct {
	StatusCode int
	Body       any
}

func CreateResponse(body any) *Response {
	return &Response{Body: body, StatusCode: http.StatusOK}
}

func CreateErrorResponse(statusCode int) *Response {
	return &Response{StatusCode: statusCode}
}
