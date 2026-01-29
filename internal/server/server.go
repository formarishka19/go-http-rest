package server

import (
	"log"
	"net/http"
	"time"

	"go-http-rest/internal/handlers"
)

type Server struct {
	logger *log.Logger
	Srv    http.Server
}

func NewHttpServer(logger *log.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", handlers.UploadHandler)
	mux.HandleFunc("/", handlers.MainHandler)
	server := new(Server)
	server.Srv.Addr = ":8080"
	server.Srv.Handler = mux
	server.Srv.ReadTimeout = 5 * time.Second
	server.Srv.WriteTimeout = 10 * time.Second
	server.Srv.IdleTimeout = 15 * time.Second
	server.logger = logger

	return server
}
