package server

import (
	"context"
	"net/http"
	"time"
)

//Server is custom wrapper for additional functionality for our server
type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,          // 1 MB //TODO? ?
		ReadTimeout:    10 * time.Second, //TODO params? config?
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
