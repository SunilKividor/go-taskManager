package main

import (
	"log"
	"net/http"
	"os"
	"taskManager/handlers"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdin, "Logged message: ", log.LstdFlags)
	td := handlers.NewTodo(l)

	sm := mux.NewRouter()

	//getRouter := sm.Methods(http.MethodGet).Subrouter()
	sm.HandleFunc("/", td.GetTodo).Methods("GET")

	//postRouter := sm.Methods(http.MethodPost).Subrouter()
	sm.HandleFunc("/post", td.PostTodo).Methods("POST")

	// putRouter := sm.Methods(http.MethodPut).Subrouter()
	sm.HandleFunc("/put/{id:[0-9]+}", td.UpdateTodo).Methods("PUT")

	sm.HandleFunc("/del/{id:[0-9]+}", td.DeleteTodo).Methods("DEL")

	s := http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	log.Print("Port running on : 8080")

	s.ListenAndServe()
}
