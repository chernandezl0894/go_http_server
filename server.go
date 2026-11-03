package main

import (
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	mux        *http.ServeMux
}

func New(addr string) *Server {
	// NOTE: Enrutador se encarga de manejar cada endpoint.
	mux := http.NewServeMux()

	// NOTE: Configuramos el servidor personalizado.
	httpServer := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       10 * time.Second,
	}

	return &Server{
		httpServer: httpServer,
		mux:        mux,
	}
}

func (server *Server) Start() error {
	// NOTE: Corremos el servidor asociado al mux.
	fmt.Println("Server running on port :8080")
	return server.httpServer.ListenAndServe()
}

func (server *Server) Handle(pattern string, handler http.Handler) {
	server.mux.Handle(pattern, handler)
}
