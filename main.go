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
	heartbeat := time.Tick(24 * time.Hour)
	for {
		select {
		case <-heartbeat:
			log.Println("Daily Scrubber Triggered...")
			db.DeleteDailyRecords(time.Now())
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

	go func() {
		log.Println("Go-Bin Server Starting on Port", *port)
		errs <- http.ListenAndServe((":" + *port), r)
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
