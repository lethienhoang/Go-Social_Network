package web

import (
	"net/http"

	"go-social-network.com/v1/internal/db"
)

type session struct {
	IsLoggedIn bool
	User       db.User
}

func (h *Handler) getSession(r *http.Request) session {
	var out session

	if h.Session.Exists(r, "user") {
		user, ok := h.Session.Get(r, "user").(db.User)
		out.User = user
		out.IsLoggedIn = ok
	}

	return out
}
