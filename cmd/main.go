package main

import (
	"log"
	"net/http"

	"github.com/abraaoneves/channel/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	handlers.RegisterChannelRoutes(router)

	log.Fatal(http.ListenAndServe(":8000", router))
}
