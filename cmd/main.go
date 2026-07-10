package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	srv := server.CreateRouter(logger)
	srv.Start()
}
