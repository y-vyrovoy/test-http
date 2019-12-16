package http_server

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

type fileHandler struct {
}

func newFileHandler() *fileHandler {
	return &fileHandler{}
}

func (h *fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Infof("got persist file. Method: %s Host: %s URL: %s", r.Method, r.Host, r.URL)

	fileName := r.Header.Get("fileName")
	if len(fileName) == 0 {
		fileName = "temp.file"
	}

	err := r.ParseMultipartForm(5 * 1024 * 1025)
	if err != nil {
		log.Errorf("failed to parse file: %s", err.Error())
		return
	}

	file, handler, err := r.FormFile("fileName")
	if err != nil {
		log.Errorf("FormFile failed: %s", err.Error())
		return
	}

	defer func() {
		_ = file.Close()
	}()

	log.Infof("file. header fName: %s, handler fName: %s, handler Header: %s", fileName, handler.Filename, handler.Header)

}
