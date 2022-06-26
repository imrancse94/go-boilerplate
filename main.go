package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-boilerplate/bootstrap"
	"go-boilerplate/database"
	"go-boilerplate/models"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()
	DB := database.InitDb()
	models.Init(DB)
	bootstrap.Init()

	PORT := os.Getenv("APP_PORT")
	fmt.Println("Listening port", PORT)
	http.ListenAndServe(":"+PORT, nil)

}
