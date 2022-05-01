package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/abraaoneves/channel/infra"
	"github.com/abraaoneves/channel/models"
	"github.com/gorilla/mux"
)

func RegisterChannelRoutes(route *mux.Router) {
	route.HandleFunc("/channels", CreateChannel).Methods("POST")
	route.HandleFunc("/channels/{id}", GetChannel).Methods("GET")
}

/**
* Create an channel
**/
//TODO: Create validation and specific methods to create channel
func CreateChannel(w http.ResponseWriter, r *http.Request) {
	request, _ := ioutil.ReadAll(r.Body)

	var channel models.Channel

	json.Unmarshal(request, &channel)

	log.Printf("persist new channel: %v \n", channel)

	redis := infra.CreateDataBaseConnection()
	channelJson, err := json.Marshal(channel)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := make(map[string]string)
		response["error"] = "Error to parse json"
		json.NewEncoder(w).Encode(response)
		return
	}

	err = redis.Set(channel.Key, channelJson, 0).Err()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := make(map[string]string)
		response["error"] = "Error to persist channel at the database"
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(channel)
}

/**
* Return Channel by Ids
 */
func GetChannel(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	id := parameters["id"]

	db := infra.CreateDataBaseConnection()
	channel, err := db.Get(id).Result()

	log.Println(channel)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(channel)
}
