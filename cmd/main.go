package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/POMBNK/avitotag/internal/delivery/banner"
	"github.com/POMBNK/avitotag/internal/pkg/client/postgres"
	bannerStorage "github.com/POMBNK/avitotag/internal/repository/banner"
	featStorage "github.com/POMBNK/avitotag/internal/repository/feature"
	tagStorage "github.com/POMBNK/avitotag/internal/repository/tag"
	bannerService "github.com/POMBNK/avitotag/internal/service/banner"
	featService "github.com/POMBNK/avitotag/internal/service/feature"
	tagService "github.com/POMBNK/avitotag/internal/service/tag"
	"github.com/go-chi/chi/v5"
)

func main() {
	pgClient, err := postgres.NewClient(context.Background(), 5)
	if err != nil {
		panic(err)
	}
	// tag init
	tagRepo := tagStorage.New(pgClient)
	tagger := tagService.New(tagRepo)

	featRepo := featStorage.New(pgClient)
	featter := featService.New(featRepo)

	// banner init
	repository := bannerStorage.New(pgClient)
	service := bannerService.New(repository, tagger, featter, pgClient)

	engine := chi.NewMux()
	var listener net.Listener
	var listenErr error
	listener, listenErr = net.Listen("tcp", "127.0.0.1:8080")
	if listenErr != nil {
		log.Fatal(listenErr)
	}
	server := http.Server{
		Handler:      banner.New(service).Register(engine),
		WriteTimeout: 120 * time.Second,
		ReadTimeout:  120 * time.Second,
	}
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	//s.logs.Println("Server started")

	<-interrupt
	//s.logs.Println("Shutting down server...")

}
