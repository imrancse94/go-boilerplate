package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go-boilerplate/bootstrap"
	"go-boilerplate/cache"
	"go-boilerplate/database"
	"go-boilerplate/models"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()
	DB := database.InitDB()
	models.Init(DB)
	bootstrap.Init()
	cache.ConnectRedis(context.Background())

	//mail.SendEmail("My subject", "This is test", "", []string{"abquddus.ctg@gmail.com", "jesse.miller.2022.smtp@gmail.com"}, "test.txt")

	PORT := os.Getenv("APP_PORT")
	fmt.Println("Listening port", PORT)
	http.ListenAndServe(":"+PORT, nil)

}
