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

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if r.PostFormValue("email") == "" {
		h.Logger.Printf("could not login: %v\n", "Email is missing")
		http.Error(w, "Email is missing", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	input := service.LoginInput{
		Email:    r.PostFormValue("email"),
		Username: formPtr(r.PostForm, "username"),
	}
	user, err := h.Service.Login(ctx, &input)
	if err != nil {
		h.Logger.Printf("could not login: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.Session.Put(r, "user", user)
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Handler) logout(w http.ResponseWriter, r *http.Request) {
	h.Session.Remove(r, "user")
	http.Redirect(w, r, "/", http.StatusFound)
}

func formPtr(form url.Values, key string) *string {
	if !form.Has(key) {
		return nil
	}

	s := form.Get(key)
	return &s
}
