package http_server

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"path/filepath"
)

type staticDataHandler struct {
	staticDataPath string
	handler        http.Handler
}

func newStaticDataHandler(staticDataPath string) *staticDataHandler {

	return &staticDataHandler{
		staticDataPath: staticDataPath,
		handler:        http.FileServer(http.Dir(filepath.Dir(staticDataPath))),
	}
}

func (h *staticDataHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Infof("got request for data. Host: %s URL: %s", r.Host, r.URL)

	h.handler.ServeHTTP(w, r)
}
