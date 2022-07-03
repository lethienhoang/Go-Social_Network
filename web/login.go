package web

import (
	"net/http"
	"net/url"
)

type loginForm struct {
	Form url.Values
	Err  error
}

func (h *Handler) renderLogin(w http.ResponseWriter, data loginForm, statusCode int) {
	var loginTmpl = parseTmpl("login.tmpl")
	h.renderTemplate(w, loginTmpl, data, statusCode)
}

func (h *Handler) loginHandler(w http.ResponseWriter, r *http.Request) {
	h.renderLogin(w, loginForm{}, http.StatusOK)
}

// func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
// 	if err := r.ParseForm(); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 	}

// 	ctx := r.Context()
// 	input := nil
// }
