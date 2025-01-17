package main

import (
	"net/http"
	"github.com/stephmukami/rss-feed-aggregator/internal/auth"
	"github.com/stephmukami/rss-feed-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){

		apiKey, err := auth.GetAPIKey(r.Header)
		if err!=nil{
			respondWithError(w, http.StatusUnauthorized, "could not find api key")
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(),apiKey)
		if err!=nil{
			respondWithError(w, http.StatusNotFound, "could not find user")
			return
		}
		handler(w,r,user)
	}
}