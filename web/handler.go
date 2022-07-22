package web

import (
	"encoding/gob"
	"log"
	"net/http"
	"sync"

	"github.com/golangcollege/sessions"
	"github.com/gorilla/mux"

	"go-social-network.com/v1/internal/db"
	service "go-social-network.com/v1/internal/services"
)

type Handler struct {
	Logger *log.Logger

	once sync.Once

	handler http.Handler

	Service *service.Service

	Session *sessions.Session

	SessionKey []byte
}

func (h *Handler) init() {
	r := mux.NewRouter()

	r.HandleFunc("/login", h.showLoginPageHandler).Methods("GET")
	r.HandleFunc("/", h.showHomePageHandler).Methods("GET")
	r.HandleFunc("/login", h.login).Methods("POST")
	r.HandleFunc("/logout", h.logout).Methods("POST")

	gob.Register(db.User{})

	h.Session = sessions.New(h.SessionKey)
	h.handler = r
	h.handler = h.Session.Enable(h.handler)
}

func (h *Handler) ServeHTTP() http.Handler {
	h.once.Do(h.init)

	return h.handler
}
