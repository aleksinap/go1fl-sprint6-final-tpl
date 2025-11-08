package main

import (
	"log"
	"net/http"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.Default()
	server := server.CreateServer(logger)

	if err := http.ListenAndServe(server.Server.Addr, server.Server.Handler); err != nil {
		logger.Fatal(err)
	}
}
