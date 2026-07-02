package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TitusW/accounts/config"
	financialaccount_handler "github.com/TitusW/accounts/internal/handler/financial_account"
	financialaccount_repo "github.com/TitusW/accounts/internal/repo/financial_account"
	financialaccount_usecase "github.com/TitusW/accounts/internal/usecase/financial_account"
	unitofwork "github.com/TitusW/unit-of-work"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	uow := unitofwork.NewGormUnitOfWork(db)

	financialAccountRepo := financialaccount_repo.New(db)
	financialAccountUsecase := financialaccount_usecase.New(financialAccountRepo, uow)
	financialAccountHandler := financialaccount_handler.New(financialAccountUsecase)

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

	router.POST("/accounts/:ksuid/debit", financialAccountHandler.DebitFinancialAccount)
	router.POST("/accounts/:ksuid/credit", financialAccountHandler.CreditFinancialAccount)

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
			TablePrefix: "game.",
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
