package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"remoteworkout/internal/infra/web/route"

	log "github.com/sirupsen/logrus"
)

type HttpServer struct {
	Origin *http.ServeMux
}

type httpEndpoint struct {
	method  string
	handler func(*Request, chan Response)
}

func (hs *HttpServer) Get(path string, handler func(*Request, chan Response)) {
	hs.Origin.HandleFunc(path, handleDecorator(http.MethodGet, handler))
}

func (hs *HttpServer) Post(path string, handler func(*Request, chan Response)) {
	hs.Origin.HandleFunc(path, handleDecorator(http.MethodGet, handler))
}

func (hs *HttpServer) Delete(path string, handler func(*Request, chan Response)) {
	hs.Origin.HandleFunc(path, handleDecorator(http.MethodGet, handler))
}

func (hs *HttpServer) Listen(port int) {
	http.ListenAndServe(fmt.Sprintf(":%d", port), hs.Origin)
}

func handleDecorator(method string, handler func(*Request, chan Response)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			if req, err := createRequest(r); err != nil {
				log.Error("Unexpected error during parsing request", err)
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				resp := handleRequest(req, handler)
				writeResponse(resp, w)
			}
		}
	}
}

func handleRequest(req *Request, handler func(*Request, chan Response)) *Response {
	c := make(chan Response)
	go handler(req, c)

	resp := <-c
	close(c)
	return &resp
}

func writeResponse(r *Response, w http.ResponseWriter) {
	responseBytes, err := json.Marshal(r.Body)
	if err != nil {
		log.Errorf("Unexpected error during serializing response: %v", r, err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(r.StatusCode)
		w.Write(responseBytes)
	}
}

func CreateHttpServer() *HttpServer {
	server := HttpServer{Origin: http.NewServeMux()}
	server.Get("/auth/login", route.GetLogin)
	server.Get("/workout", route.GetWorkouts)
	return &server
}
