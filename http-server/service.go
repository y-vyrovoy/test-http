package http_server

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type service struct {
	server *http.Server
}

func New(staticDataPath, port string) *service {

	mux := http.NewServeMux()
	mux.Handle("/data/", newStaticDataHandler(staticDataPath))
	mux.Handle("/", newDynamicHandler())
	mux.Handle("/file", newFileHandler())

	return &service{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%s", port),
			Handler: mux,
		},
	}
}

func (s *service) Run(ctx context.Context, wg *sync.WaitGroup) {

	wg.Add(1)

	go func() {
		defer wg.Done()

		log.Info("http server runs")

		err := s.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Errorf("http server failed to close: %s", err.Error())
		}
	}()

	go func() {
		<-ctx.Done()

		ctxShutdown, _ := context.WithTimeout(context.Background(), 5*time.Second)
		s.server.Shutdown(ctxShutdown)
	}()
}
