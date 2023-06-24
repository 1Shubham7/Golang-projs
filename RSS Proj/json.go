package main

import (
	"encoding/json"
	"net/http"
	"log"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}){
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed ot marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
}