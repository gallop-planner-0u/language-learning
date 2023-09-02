package main

import (
	"context"
	authhandler "language-learning/cmd/auth/controller/http"
	authrepository "language-learning/cmd/auth/repository"
	authusecase "language-learning/cmd/auth/usecase"
	"language-learning/pkg"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	client := pkg.GetEntClient()
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal(err)
	}

	_authRepo := authrepository.New(client)
	_authUsecase := authusecase.New(_authRepo)
	_authHandler := authhandler.New(_authUsecase)

	r := gin.Default()
	authhandler.RegisterRoutes(r, _authHandler)

	r.Run(":" + viper.GetString("app.port"))
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "../config"
	}
	viper.AddConfigPath(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
