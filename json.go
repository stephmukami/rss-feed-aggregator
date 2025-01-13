package main

import (
	"encoding/json"
	"net/http"
	"log"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}){
	dat,err := json.Marshal(payload)

	if err != nil{
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code)
	w.Write(dat)
}

func respondWithError(w http.ResponseWriter, code int, msg string){
	if code > 499{
		log.Println("Respond with SXX error",msg)
	}
	// the struct will be turned to json, code specifies there is an Error object of type string and its key is error
	type errResponse struct{
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errResponse{
		Error:msg,
	})
}