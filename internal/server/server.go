package server

import (
	"context"
	"log"

	"github.com/SalhSThai/go-first-classes/internal/client"
	"github.com/SalhSThai/go-first-classes/internal/handler"
	"github.com/SalhSThai/go-first-classes/internal/repository"
	"github.com/SalhSThai/go-first-classes/internal/router"
	"github.com/SalhSThai/go-first-classes/internal/service"
)

func Start() error {
	ctx := context.Background()

	// ===========================> create repository
	repo, err := repository.New(ctx)
	if err != nil {
		log.Panicf("unable to init repository: %v", err)
	}

	// ===========================> create client
	webclient, err := client.New(ctx)
	if err != nil {
		log.Panicf("unable to init web client: %v", err)
	}

	// ===========================> create service
	newservice, err := service.New(repo, webclient)
	if err != nil {
		log.Panicf("unable to init service: %v", err)
	}

	// ===========================> create handler
	serviceHandler, err := handler.New(newservice)
	if err != nil {
		log.Panicf("unable to init handler: %v", err)
	}

	// ===========================> create server
	engine := InitServer()

	// ===========================> create router
	router.AddRouter(engine, serviceHandler)

	return nil
}
