package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/abraaoneves/channel/handlers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/channels", handlers.GetAllChannels).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
