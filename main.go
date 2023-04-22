package main

import (
	"context"
	"handlers/hello/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodBye(l)
	ph := handlers.NewProducts(l)
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/good-bye", gh)
	sm.Handle("/product-api", ph)
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	// })
	// http.HandleFunc("/good-bye", func(w http.ResponseWriter, r *http.Request) {
	// 	log.Println("Good bye !")
	// })
	server := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan

	l.Println("Received terminate command , graceful shutdow ", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)
	//http.ListenAndServe("127.0.0.1:9090", sm)
}
