package main

import (
	"context"
	"github.com/bruceneco/nic-ms/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	// Handlers
	hh := handlers.NewHello(l)
	ph := handlers.NewProducts(l)

	// Create serve mux and register the handlers
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/products", ph)

	// create server
	s := http.Server{Addr: ":9090",
		Handler:      sm,
		IdleTimeout:  2 * time.Minute,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	// start server
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	log.Printf("Server started on port %s\n", s.Addr)

	// trap sigterm or interrupt and gracefully shutdown the server
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received  terminate, gracefully shutting down:", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
