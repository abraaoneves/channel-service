package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/abraaoneves/channel/models"
)

/**
* Get all channels from database
* This is an temporary function
**/
func GetAllChannels(w http.ResponseWriter, r *http.Request) {
	channels := []models.Channel{
		{Key: "key-1", Composer: "composer-1", Node: "node-1"},
		{Key: "key-2", Composer: "composer-2", Node: "node-2"},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(channels)
}

/**
* Return Channel by Ids
 */
func GetChannel(w http.ResponseWriter, r *http.Request) {
	channel := models.Channel{
		Key:      "key-1",
		Composer: "composer-1",
		Node:     "node-1",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(channel)
}
