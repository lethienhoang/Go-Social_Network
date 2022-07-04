package web

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"

	service "go-social-network.com/v1/internal/services"
)

type Handler struct {
	Logger *log.Logger

	once sync.Once

	handler http.Handler

	Service *service.Service
}

func (h *Handler) init() {
	r := mux.NewRouter()

	r.HandleFunc("/login", h.showLoginPageHandler).Methods("GET")
	r.HandleFunc("/login", h.loginFormHandler).Methods("POST")

	h.handler = r
}

func (h *Handler) ServeHTTP() http.Handler {
	h.once.Do(h.init)

	return h.handler
}
