package main

import (
	"fmt"
	"log"
	"net/http"

	myserver "go-http-rest/internal/server"
)

func main() {
	httpLogger := new(log.Logger)
	httpServer := myserver.NewHttpServer(httpLogger)
	err := http.ListenAndServe(httpServer.Addr, httpServer.Handler)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Сервер работает")

}
