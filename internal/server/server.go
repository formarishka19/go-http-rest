package server

import (
	"log"
	"net/http"
	"time"

	"go-http-rest/internal/handlers"
)

type HttpServer struct {
	Addr         string
	Handler      *http.ServeMux
	ErrorLog     *log.Logger
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func NewHttpServer(logger *log.Logger) *HttpServer {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", handlers.UploadHandler)
	mux.HandleFunc("/", handlers.MainHandler)
	server := new(HttpServer)
	server.Addr = ":8080"
	server.Handler = mux
	server.ErrorLog = logger
	server.ReadTimeout = 5 * time.Second
	server.WriteTimeout = 10 * time.Second
	server.IdleTimeout = 15 * time.Second

	return server
}
