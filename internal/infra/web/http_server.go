package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"remoteworkout/internal/infra/web/request"

	log "github.com/sirupsen/logrus"
)

type HttpServer struct {
	Origin *http.ServeMux
}

func (hs *HttpServer) Get(path string, handler func(*request.Request, chan request.Response)) {
	hs.Origin.HandleFunc(path, handleDecorator(http.MethodGet, handler))
}

func (hs *HttpServer) Post(path string, handler func(*request.Request, chan request.Response)) {
	hs.Origin.HandleFunc(path, handleDecorator(http.MethodGet, handler))
}

func (hs *HttpServer) Delete(path string, handler func(*request.Request, chan request.Response)) {
	hs.Origin.HandleFunc(path, handleDecorator(http.MethodGet, handler))
}

func (hs *HttpServer) Listen(port int) {
	http.ListenAndServe(fmt.Sprintf(":%d", port), hs.Origin)
}

func handleDecorator(method string, handler func(*request.Request, chan request.Response)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			if req, err := request.CreateRequest(r); err != nil {
				log.Error("Unexpected error during parsing request", err)
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				resp := handleRequest(req, handler)
				writeResponse(resp, w)
			}
		}
	}
}

func handleRequest(req *request.Request, handler func(*request.Request, chan request.Response)) *request.Response {
	c := make(chan request.Response)
	go handler(req, c)

	resp := <-c
	close(c)
	return &resp
}

func writeResponse(r *request.Response, w http.ResponseWriter) {
	responseBytes, err := json.Marshal(r.Body)
	if err != nil {
		log.Errorf("Unexpected error during serializing response: %v", r, err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(r.StatusCode)
		w.Write(responseBytes)
	}
}
