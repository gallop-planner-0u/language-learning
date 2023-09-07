package main

import (
	"context"
	authhandler "language-learning/cmd/auth/controller/http"
	authrepository "language-learning/cmd/auth/repository"
	authusecase "language-learning/cmd/auth/usecase"
	dictionaryhandler "language-learning/cmd/dictionary/controller/http"
	dictionaryrepository "language-learning/cmd/dictionary/repository"
	dictionaryusecase "language-learning/cmd/dictionary/usecase"
	"language-learning/pkg"
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
		panic(err)
	}

	_authRepo := authrepository.New(client)
	_authUsecase := authusecase.New(_authRepo)
	_authHandler := authhandler.New(_authUsecase)

	_dictionaryRepo := dictionaryrepository.New(client)
	_dictionaryUsecase := dictionaryusecase.New(_dictionaryRepo)
	_dictionaryHandler := dictionaryhandler.New(_dictionaryUsecase)

	r := gin.Default()
	authhandler.RegisterRoutes(r, _authHandler)
	dictionaryhandler.RegisterHTTPRoutes(r, _dictionaryHandler, _authHandler.AuthMiddleware)

	r.Run(":" + viper.GetString("app.port"))
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "/app/config"
	}
	viper.AddConfigPath(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
