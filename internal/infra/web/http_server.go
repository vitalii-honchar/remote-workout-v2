package web

import (
	"fmt"
	"net/http"
)

type HttpServer struct {
	Origin *http.ServeMux
}

func (hs *HttpServer) Get(path string, handler func(http.ResponseWriter, *http.Request)) {
	hs.Origin.HandleFunc(path, decorateHttpMethod(http.MethodGet, handler))
}

func (hs *HttpServer) Post(path string, handler func(http.ResponseWriter, *http.Request)) {
	hs.Origin.HandleFunc(path, decorateHttpMethod(http.MethodPost, handler))
}

func (hs *HttpServer) Delete(path string, handler func(http.ResponseWriter, *http.Request)) {
	hs.Origin.HandleFunc(path, decorateHttpMethod(http.MethodPost, handler))
}

func (hs *HttpServer) Listen(port int) {
	http.ListenAndServe(fmt.Sprintf(":%d", port), hs.Origin)
}

func decorateHttpMethod(method string, handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			handler(w, r)
		}
	}
}
