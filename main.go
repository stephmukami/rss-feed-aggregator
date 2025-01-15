package main

import (
	
	"log"
	"net/http"
	"os"
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/stephmukami/rss-feed-aggregator/internal/database"
	_ "github.com/lib/pq"
)

type apiConfig struct{
	DB*database.Queries
}
func main(){


godotenv.Load(".env")

portString := os.Getenv("PORT")

if portString == ""{
	log.Fatal("PORT is not found")
}

dbURL := os.Getenv("DB_URL")

if dbURL == "" {
	log.Fatal("DATABASE_URL environment variable is not set")
}

db, err := sql.Open("postgres", dbURL)
if err != nil {
	log.Fatal(err)
}

dbQueries := database.New(db)

apiCfg := apiConfig{
	DB: dbQueries,
}

router := chi.NewRouter()

router.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"https://*", "http://*"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300, 
  }))

v1Router := chi.NewRouter()
v1Router.Get("/health-check",handlerReadiness)
v1Router.Get("/err",handlerErr)
v1Router.Post("/users",apiCfg.handlerCreateUser)
v1Router.Get("/users",apiCfg.handlerUsersGet)

//v1Router.Post("/users", apiCfg.handlerUsersCreate)

//for api versioning
router.Mount("/v1",v1Router)

srv := &http.Server{
	Handler: router,
	Addr: ":"+ portString,
}
log.Printf("Serving on port: %s\n", portString)
log.Fatal(srv.ListenAndServe())
}