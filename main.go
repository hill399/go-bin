package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/hill399/go-bin/api/v1"
	"github.com/hill399/go-bin/db"
)

var (
	version = "1.0.0"
	port    = flag.String("port", "8080", "Port for server to listen on")
)

func init() {
	flag.Parse()
	godotenv.Load(".env")
}

func dailyScrubber() {
	log.Println("Initialising Scrubber")
	db.DeleteExpiredRecords()

	heartbeat := time.Tick(24 * time.Hour)
	for {
		select {
		case <-heartbeat:
			log.Println("Daily Scrubber Triggered")
			db.DeleteExpiredRecords()
		}
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/submit", api.SubmitData).Methods("POST")
	r.HandleFunc("/request/{id}", api.RequestData).Methods("GET")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./client/build")))

	errs := make(chan error, 1)

	go dailyScrubber()

	srv := &http.Server{
		Addr:              (":" + *port),
		Handler:           r,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	go func() {
		log.Println("Go-Bin Server Starting on Port", *port)
		errs <- srv.ListenAndServe()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case <-stop:
		log.Println("Shutting Down Safely...")
	case err := <-errs:
		log.Println("Failed to start server:", err.Error())
		runtime.Goexit()
	}
}
