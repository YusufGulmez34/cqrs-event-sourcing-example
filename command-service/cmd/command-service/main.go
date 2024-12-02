package main

import (
	"command-service/config"
	_ "command-service/docs"
	"command-service/internal/api/useraccountapi"
	"command-service/internal/repository/useraccountrepo"
	"command-service/internal/service/useraccountservice"
	"command-service/pkg/db"
	"command-service/pkg/messagebroker"
	"command-service/pkg/validator/echovalidator"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echoswagger "github.com/swaggo/echo-swagger"
)

// @title Command-Service Backend Api
// @version 1.0
// @description Command-Service backend api documents
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	config := &config.DBConfig{}
	config.ReadConfigFile()

	e := echo.New()
	e.Validator = echovalidator.NewCustomValidator(validator.New())
	e.GET("swagger/*", echoswagger.WrapHandler)

	gormDB := db.ConnectMysql(config.MySql)
	rabbitMQConn, err := messagebroker.ConnectRabbitMQ()
	if err != nil {
		panic(err)
	}

	userAccountRepo := useraccountrepo.NewUserAccountRepository(gormDB)
	userAccountService := useraccountservice.NewUserAccountService(*userAccountRepo, rabbitMQConn)
	userAccountApi := useraccountapi.NewUserAccountApi(*userAccountService)
	userAccountApi.Group(e.Group("user-accounts"))

	if err := userAccountRepo.Migrate(); err != nil {
		panic(err)
	}

	e.Start(":5001")
}
