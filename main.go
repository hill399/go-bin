package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/gorilla/mux"

	"github.com/hill399/go-bin/api/v1"
	"github.com/hill399/go-bin/db"
)

var (
	version = "1.0.0"
	port    = flag.String("port", "8080", "Port for server to listen on")
)

func init() {
	flag.Parse()
}

func main() {
	db.Database = db.InitDatabase()
	defer db.Database.Close()

	router := mux.NewRouter()
	router.HandleFunc("/submit", api.SubmitData).Methods("POST")
	router.HandleFunc("/request/{id}", api.RequestData).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./client/build")))

	errs := make(chan error, 1)

	go func() {
		fmt.Println("Go-Bin Server Starting on Port", *port)
		errs <- http.ListenAndServe((":" + *port), router)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case <-stop:
		fmt.Println("Shutting Down Safely...")
	case err := <-errs:
		fmt.Println("Failed to start server:", err.Error())
		runtime.Goexit()
	}

}
