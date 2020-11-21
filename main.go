package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/beltranbot/go-microservices-introduction/handlers"
)

func main() {
	mainLogger := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(mainLogger)
	goodbyeHandler := handlers.NewGoodbye(mainLogger)
	productHandler := handlers.NewProducts(mainLogger)

	serverMux := http.NewServeMux()
	serverMux.Handle("/", productHandler)
	serverMux.Handle("/hello", helloHandler)
	serverMux.Handle("/goodbye", goodbyeHandler)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      serverMux,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			mainLogger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	mainLogger.Println("Received terminate, graceful shutdown", sig)
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(timeoutContext)
}
