package main

import (
	"log"
	"net/http"
	"os"

	"github.com/vlonebara/morseConverter/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	srv := server.CreateRouter(logger)
	if err := srv.Start(); err != nil && err != http.ErrServerClosed {
		logger.Fatal(err)
	}
}
