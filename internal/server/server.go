package server

import (
	"log"
	"net/http"
	"time" // Добавьте этот импорт

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	HTTP   *http.Server
}

// Создает роутер
func CreateRouter(logger *log.Logger) *Server {
	router := http.NewServeMux()

	router.HandleFunc("/", handlers.IndexHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,  // Исправлено: добавлен time.Second
		WriteTimeout: 10 * time.Second, // Исправлено: добавлен time.Second
		IdleTimeout:  15 * time.Second, // Исправлено: добавлен time.Second
	}

	return &Server{
		Logger: logger,
		HTTP:   httpServer,
	}
}

func (s *Server) Start() error {
	s.Logger.Printf("Запуск сервера на порту 8080...")
	return s.HTTP.ListenAndServe()
}
