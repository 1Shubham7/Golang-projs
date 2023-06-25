package main

import (
	"encoding/json"
	"net/http"
	"log"
)

func respondWithError(w http.ResponseWriter, code int, msg string){
	if code > 499 {
		log.Println("Responding with 5XX error", msg)
	}
	type errResponse struct {
		Error string `json:'error'`
	}

	respondWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}){
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed ot marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dat)
}