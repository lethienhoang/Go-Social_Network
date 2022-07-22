package web

import "net/http"

type homeForm struct {
	Session session
}

func (h *Handler) renderHome(w http.ResponseWriter, data homeForm, statusCode int) {
	var homeTmpl = parseTmpl("home.tmpl")
	h.renderTemplate(w, homeTmpl, data, statusCode)
}

func (h *Handler) showHomePageHandler(w http.ResponseWriter, r *http.Request) {
	homeForm := homeForm{
		Session: h.getSession(r),
	}
	h.renderHome(w, homeForm, http.StatusOK)
}
