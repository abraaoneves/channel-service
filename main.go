package main

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
)


type Channel struct {
    Key         string
    Composer    string
    Node        string
}

func main() {
    router := mux.NewRouter().StrictSlash(true);
    router.HandleFunc("/channels", getChannels).Methods("GET");

    http.ListenAndServe(":8000", router)
}


func getChannels(w http.ResponseWriter, r *http.Request) {
    channels := []Channel {
        {"key-1234", "compose-1234", "node-1234"},
        {"key-5678", "compose-5678", "node-5678"},
    }


    w.WriteHeader(http.StatusOK);
    json.NewEncoder(w).Encode(channels);
}
