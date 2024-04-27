package main

import (
	"context"
	"github.com/POMBNK/avitotag/internal/delivery/tag"
	"github.com/POMBNK/avitotag/internal/pkg/client/postgres"
	tagStorage "github.com/POMBNK/avitotag/internal/repository/tag"
	tagService "github.com/POMBNK/avitotag/internal/service/tag"
	"github.com/go-chi/chi/v5"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	pgClient, err := postgres.NewClient(context.Background(), 5)
	if err != nil {
		panic(err)
	}
	repository := tagStorage.New(pgClient)
	service := tagService.New(repository)
	_ = service

	engine := chi.NewRouter()
	var listener net.Listener
	var listenErr error
	listener, listenErr = net.Listen("tcp", "127.0.0.1:8080")
	if listenErr != nil {
		panic(listenErr)
	}
	server := http.Server{
		Handler:      tag.New(engine).Register(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
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
