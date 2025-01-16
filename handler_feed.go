package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/stephmukami/rss-feed-aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerFeedCreate(w http.ResponseWriter, r *http.Request, user database.User){
	type parameters struct{
		Name string `json:"name"`
		URL string `json:"url"`

	}
	 decoder := json.NewDecoder(r.Body)
	 params := parameters{}
	 err := decoder.Decode(&params)

	 if err!=nil{
		respondWithError(w, http.StatusInternalServerError, "could not decode parameters")
		return
	}

	feed,err := apiCfg.DB.CreateFeed(r.Context(),database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		Name:      params.Name,
		Url:       params.URL,
	})

	if err!=nil{
		respondWithError(w, http.StatusInternalServerError, "could not create feed")
		return
	}
	respondWithJSON(w, http.StatusOK,databaseFeedToFeed(feed))
}