package web

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type Handler struct {
	Logger *log.Logger

	once sync.Once

	handler http.Handler
}

func (h *Handler) init() {
	r := mux.NewRouter()

	r.HandleFunc("/login", h.loginHandler).Methods("GET")

	h.handler = r
}

func (h *Handler) ServeHTTP() http.Handler {
	h.once.Do(h.init)

	return h.handler
}
