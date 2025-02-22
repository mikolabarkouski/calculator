package api

import (
	"net/http"
	"text/template"

	"github.com/mikolabarkouski/calculator/templates"
)

func (h *Handler) RenderApp(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("index").Parse(templates.IndexHTML))
	packagesDefault := h.app.GetPackagesMap()
	tmpl.Execute(w, packagesDefault)
}
