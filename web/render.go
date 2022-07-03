package web

import (
	"bytes"
	"embed"
	"fmt"
	"net/http"

	"html/template"
)

//go:embed template/*.tmpl
var templateFS embed.FS

func parseTmpl(tmplName string) *template.Template {
	tmpl := template.New(tmplName)
	tmpl = template.Must(tmpl.ParseFS(templateFS, "template/layout.tmpl"))
	parseFs, err := tmpl.ParseFS(templateFS, "template/"+tmplName)
	if err != nil {
		fmt.Printf("error parsing template: %s", err.Error())
	}

	template := template.Must(parseFs, err)
	if err != nil {
		fmt.Printf("template not found: %s", err.Error())
	}
	return template
}

func (h *Handler) renderTemplate(w http.ResponseWriter, tmpl *template.Template, data interface{}, status int) {
	var buff bytes.Buffer

	err := tmpl.Execute(&buff, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.Logger.Output(2, fmt.Sprintf("could not render %q: %v\n", tmpl.Name(), err))
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=uff-8")
	w.WriteHeader(status)

	_, err = buff.WriteTo(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.Logger.Output(2, fmt.Sprintf("could not render %q: %v\n", tmpl.Name(), err))
	}
}
