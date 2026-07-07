package server

import (
	"log"
	"net/http"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	HTTP   *http.Server
}

// Создает роутер
func CreateRoter(logger *log.Logger) *Server {
	router := http.NewServeMux()

	router.HandleFunc("/", handlers.IndexHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5,
		WriteTimeout: 10,
		IdleTimeout:  15,
	}

	return &Server{
		Logger: logger,
		HTTP:   httpServer,
	}
}
