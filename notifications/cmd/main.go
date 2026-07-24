package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/TitusW/notifications/config"
	notification_handler "github.com/TitusW/notifications/internal/handler/notification"
	notification_repo "github.com/TitusW/notifications/internal/repo/notification"
	notificationdispatch_repo "github.com/TitusW/notifications/internal/repo/notification-dispatch"
	notification_usecase "github.com/TitusW/notifications/internal/usecase/notification"
	"github.com/TitusW/notifications/internal/worker"
	unitofwork "github.com/TitusW/unit-of-work"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	ctx := context.Background()

	router := gin.Default()

	uow := unitofwork.NewGormUnitOfWork(db)
	notificationRepo := notification_repo.New(db)
	notificationDispatchesRepo := notificationdispatch_repo.New(db)

	notificationUsecase := notification_usecase.New(notificationRepo, notificationDispatchesRepo, uow)
	notificationHandler := notification_handler.New(notificationUsecase)

	workerPool := worker.NewWorkerPool(5, 100, notificationDispatchesRepo)
	workerPool.Start(ctx)

	dbPoller := worker.NewDBPoller(notificationDispatchesRepo, workerPool)
	dbPoller.StartPolling(ctx)

	router.GET("/health", func(ctx *gin.Context) {
		// logger
		ctx.JSON(200, gin.H{
			"message": "HEALTHY!",
		})
	})

	//Used for metrics test
	router.GET("/", func(ctx *gin.Context) {
		// logger
		ctx.JSON(200, gin.H{
			"message": "Hello this is notification service!",
		})
	})

	router.POST("/notifications/", notificationHandler.CreateNotification)

	return router
}

func setupHttp(router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
}

func setupDB(configData config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		configData.Database.Host,
		configData.Database.Username,
		configData.Database.Password,
		configData.Database.Name,
		configData.Database.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "notification.",
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func main() {
	mainCfg := config.InitializeConfig()

	db := setupDB(mainCfg)

	router := setupRouter(db)

	server := setupHttp(router)

	server.ListenAndServe()
}
