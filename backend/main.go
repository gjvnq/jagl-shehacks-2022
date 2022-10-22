package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	UsersModuleLoad()
	ProfilesModuleLoad()

	my_cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins
		AllowedMethods: []string{"GET", "PUT", "POST"},
	})

	shutdown_deadline := 15 * time.Second
	addr := "127.0.0.1:3001"
	router := mux.NewRouter()
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})
	router.HandleFunc("/users/list", ListUsers).Methods("GET", "OPTIONS")
	router.HandleFunc("/users/new", AddUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/profiles/{ulid}", ShowProfile).Methods("GET", "OPTIONS")
	router.HandleFunc("/profiles/{ulid}", SaveProfile).Methods("PUT", "OPTIONS")

	srv := &http.Server{
		Handler: my_cors.Handler(router),
		Addr:    addr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  time.Second * 60,
	}

	go func() {
		log.Printf("Listening on %s", "127.0.0.1:3001")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	stop_ch := make(chan os.Signal, 1)
	signal.Notify(stop_ch, os.Interrupt)
	// Wait for SIG_TERM
	<-stop_ch
	ctx, cancel := context.WithTimeout(context.Background(), shutdown_deadline)
	defer cancel()
	srv.Shutdown(ctx)
	UsersModuleSave()
	ProfilesModuleSave()
	log.Print("shutting down")
	os.Exit(0)
}
