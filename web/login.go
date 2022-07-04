package web

import (
	"net/http"
	"net/url"

	service "go-social-network.com/v1/internal/services"
)

type loginForm struct {
	Form url.Values
	Err  error
}

func (h *Handler) renderLogin(w http.ResponseWriter, data loginForm, statusCode int) {
	var loginTmpl = parseTmpl("login.tmpl")
	h.renderTemplate(w, loginTmpl, data, statusCode)
}

func (h *Handler) showLoginPageHandler(w http.ResponseWriter, r *http.Request) {
	h.renderLogin(w, loginForm{}, http.StatusOK)
}

func (h *Handler) loginFormHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	input := service.LoginInput{
		Email:    r.PostFormValue("email"),
		Username: formPtr(r.PostForm, "username"),
	}
	h.Service.Login(ctx, &input)
}

func formPtr(form url.Values, key string) *string {
	if !form.Has(key) {
		return nil
	}

	s := form.Get(key)
	return &s
}
