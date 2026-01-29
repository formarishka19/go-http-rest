package main

import (
	"log"
	"net/http"

	myserver "go-http-rest/internal/server"
)

func main() {
	httpLogger := new(log.Logger)
	httpServer := myserver.NewHttpServer(httpLogger)
	err := http.ListenAndServe(httpServer.Srv.Addr, httpServer.Srv.Handler)
	if err != nil {
		log.Fatal(err)
	}
}
