package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/google/uuid"
	"time"
	"github.com/stephmukami/rss-feed-aggregator/internal/auth"

	"github.com/stephmukami/rss-feed-aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request){
	type parameters struct{
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err!=nil{
		respondWithError(w,400,fmt.Sprintf("Error parsing JSON %v",err))
		return
	}

	user, err := apiCfg.DB.CreateUser(
		r.Context(),
		database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,})

		if err!=nil{
			respondWithError(w,400,fmt.Sprintf("Couldn't create user %v",err))
			return
		}

	// respondWithJSON(w,200,user)
	respondWithJSON(w,201,databaseToUser(user))

}

func (apiCfg *apiConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request){
	apiKey, err := auth.GetAPIKey(r.Header)

	if err !=nil{
		respondWithError(w,http.StatusUnauthorized,"could not find API key")
		return
	}

	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(),apiKey)
	if err !=nil{
		respondWithError(w,400, fmt.Sprintf("could not get user: %v",err))
		return
	}

	
	respondWithJSON(w,200,databaseToUser(user))

}