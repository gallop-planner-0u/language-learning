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
)

func main() {
	ctx := context.Background()
	client := pkg.GetEntClient(ctx)
	defer client.Close()

	_authRepo := authrepository.New(client)
	_authUsecase := authusecase.New(_authRepo)
	_authHandler := authhandler.New(_authUsecase)

	_dictionaryRepo := dictionaryrepository.New(client)
	_dictionaryUsecase := dictionaryusecase.New(_dictionaryRepo)
	_dictionaryHandler := dictionaryhandler.New(_dictionaryUsecase)

	r := gin.Default()
	authhandler.RegisterRoutes(r, _authHandler)
	dictionaryhandler.RegisterHTTPRoutes(r, _dictionaryHandler, _authHandler.AuthMiddleware)

	r.Run(":" + os.Getenv("PORT"))
}
