package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/natefinch/lumberjack.v2"
)

func addRestaurant(w http.ResponseWriter, r *http.Request) {

	log.Printf("Received request for %s\n", view)
	w.Write([]byte(fmt.Sprintf("Hello, %s\n", view)))
}

func viewRestaurant(w http.ResponseWriter, r *http.Request) {

	log.Printf("Received request for %s\n", view)
	w.Write([]byte(fmt.Sprintf("Hello, %s\n", view)))
}

func handler(w http.ResponseWriter, r *http.Request) {

	log.Printf("Received request for %s\n", view)
	w.Write([]byte(fmt.Sprintf("Hello, %s\n", view)))
}

func main() {
	// Create Server and Route Handlers
	r := mux.NewRouter()

	r.HandleFunc("/", handler)

	r.HandleFunc("/addRestaurant", addRestaurant).Methods("POST");

	r.HandleFunc("/viewRestaurant", viewRestaurant).Methods("GET");

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Configure Logging
	LOG_FILE_LOCATION := os.Getenv("LOG_FILE_LOCATION")
	if LOG_FILE_LOCATION != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   LOG_FILE_LOCATION,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		})
	}

	// Start Server
	go func() {
		log.Println("Starting Server")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}