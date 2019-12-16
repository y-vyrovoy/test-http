package http_server

import (
	log "github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"path/filepath"
)

type dynamicHandler struct {
	tmpl *template.Template
}

func newDynamicHandler() *dynamicHandler {

	tmpl, err := template.ParseFiles(filepath.Join("template", "form.html"))
	if err != nil {
		log.Fatalf("failed to parse template file: %s", err.Error())
	}

	return &dynamicHandler{
		tmpl: tmpl,
	}
}

func (h *dynamicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Infof("got dynamic request. Method: %s Host: %s URL: %s", r.Method, r.Host, r.URL)

	data := struct {
		Text string
	}{Text: r.URL.String()}

	w.Header().Add("Content-Type", "text/html")
	_ = h.tmpl.Execute(w, data)
}
