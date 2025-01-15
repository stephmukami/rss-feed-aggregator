package main

import(
"github.com/stephmukami/rss-feed-aggregator/internal/database"
"time"
"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Name      string `json:"name"`
	APIKey      string `json:"api_key"`

}

func databaseToUser(user database.User) User{
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name :     user.Name,
		APIKey :     user.ApiKey,

 	}
}

