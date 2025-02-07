package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jaiyanth10/goApi/internal/config"
	"github.com/jaiyanth10/goApi/internal/http/handlers/student"
)

func main() {
	// step 1: load config
	cfg := config.MustLoad() //calling must load fun from config packge which sets the config elemnts from local.yaml file
	
	// step 2: database setup

	// step 3: setup router
	//First creating router to handle get request and associated method
	router := http.NewServeMux()
	router.HandleFunc("POST /api/student", student.New())

	// step 4: configuring server
	server := http.Server{ //inbuilt server struct
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	// starting server
	slog.Info("server started", slog.String("address", cfg.HTTPServer.Addr))//logging msg
	//interuption handling
	done := make(chan os.Signal, 1) //created a channel called chan with size 1
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)	// Notify method will pass the occurence of below signals to the channel "done"
	// creating a go routine to start server
	//this will execute  once when project run and execution blocked at below <-done,means server will be running uninteruptedly
	//if interuption happens then done channel will send signal and execution unblocks and the below code of <-done will execute
	//if no interuption happens, then at <-done the exection will be blocked and execution of main wont exit and server will keep running.
	go func() {
		err := server.ListenAndServe()//server start
		if err == nil {
			log.Fatal("failed to start server!")
		}
	}()
	<-done //firstly, execution blocked when goroutine executed and when the interuption happen then here execution will unblock

	//because of interuption here we are gracefully shutting down server
	//putting a time to shutdown server if it didnt, will let us know using context
	//if no timer, server shutdown will go to infinite loop
	slog.Info("Shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()              //as we have only main function differ will excetute before exiting the main function
	err := server.Shutdown(ctx) //shutting server by sending the ctx as parameter
	//if any error thrown
	if err != nil {
		slog.String("error", err.Error())
	}
	slog.Info("server shutdown successfully ")
}

