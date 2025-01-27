package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/stephmukami/rss-feed-aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerFeedFollowCreate(w http.ResponseWriter, r *http.Request, user database.User){
	type parameters struct{
		FeedID uuid.UUID
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	fedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		errorMessage := fmt.Sprintf("Couldn't create feed follow: %v", err)
		respondWithError(w, http.StatusInternalServerError, errorMessage)

		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedFollowToFeedFollow(fedFollow))
}

func (apiCfg *apiConfig) handlerFeedFollowsGet(w http.ResponseWriter, r *http.Request, user database.User){
	feedFollows, err := apiCfg.DB.GetFeedFollowsForUser(r.Context(),user.ID)
	if err !=nil{
		errorMessage := fmt.Sprintf("Couldn't create feed follow: %v", err)
		respondWithError(w, http.StatusInternalServerError, errorMessage)

		return
	}
	var result []FeedFollow
	for _, feedFollow := range feedFollows {
		result = append(result, databaseFeedFollowToFeedFollow(feedFollow))
	}

	respondWithJSON(w, http.StatusOK, result)


}