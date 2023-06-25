package main

import "net/http"

func handleErr(w http.ResponseWriter, r *http.Request){
	respondWithError(w, 200, "Something went wrong")
}