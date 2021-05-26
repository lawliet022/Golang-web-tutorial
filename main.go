package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"example.com/hello/tutorial2/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api ", log.LstdFlags) // a new Logger (maybe to a database)

	ph := handlers.NewProduct(l) // Product handler

	sm := http.NewServeMux() // serve mux
	sm.Handle("/", ph)       // Register any path to product handler

	s := &http.Server{ // custom server
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err) // Fatal exists with a status code of err
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate. Gracefully shutting down - ", sig)

	// wait for 30 sec to gracefully finish all work, after that force close.
	// Server do not accepts new requests in this 30 sec of time
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second) // timeout context
	defer cancel()                                                          // releases resources if ops completes before timeout elapses

	s.Shutdown(tc)

}
