package main

import (
	"os"

	"github.com/SalhSThai/go-first-classes/internal/server"
	"github.com/google/martian/log"
)

func main() {
	if err := server.Start(); err != nil {
		log.Errorf("Run server failed!", err)
		os.Exit(1)
	}
}
