package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

func CreateServer(logger *log.Logger) Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HandleMain)
	mux.HandleFunc("/upload", handlers.HandleUpload)

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  time.Duration(5) * time.Second,
		WriteTimeout: time.Duration(10) * time.Second,
		IdleTimeout:  time.Duration(15) * time.Second,
	}
	return Server{
		Logger: logger,
		Server: &server,
	}
}
