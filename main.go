package main

import (
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
)

func main(){

fmt.Println("hello")

godotenv.Load(".env")

portString := os.Getenv("PORT")

if portString == ""{
	log.Fatal("PORT is not found")
}
fmt.Println("Port: ",portString)
}