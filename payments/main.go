package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupGin() *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(ctx *gin.Context) {
		// logger.KargoLog.Info("Server is healthy")
		ctx.JSON(200, gin.H{
			"message": "HEALTHY!",
		})
	})

	//Used for metrics test
	router.GET("/", func(ctx *gin.Context) {
		// logger.KargoLog.Info("Return ")
		ctx.JSON(200, gin.H{
			"message": "Hello this is PI-Company!",
		})
	})

	router.POST("/transfers", func(ctx *gin.Context) {

	})

	return router
}

func setupHttp(router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
}

func main() {
	router := setupGin()

	server := setupHttp(router)

	server.ListenAndServe()
}
