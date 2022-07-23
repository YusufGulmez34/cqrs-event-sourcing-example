package main

import (
	_ "query-service/docs"
	"query-service/internal/api/useraccountapi"
	"query-service/internal/repository/useraccountrepo"
	"query-service/internal/service/useraccountservice"
	"query-service/pkg"

	"github.com/labstack/echo/v4"
	echoswagger "github.com/swaggo/echo-swagger"
)

// @title Query-Service Backend Api
// @version 1.0
// @description Query-Service backend api documents
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	e := echo.New()
	e.GET("swagger/*", echoswagger.WrapHandler)
	scope := pkg.ConnectCB()
	userAccountRepo := useraccountrepo.NewUserAccountRepository(scope)
	userAccountService := useraccountservice.NewUserAccountService(*userAccountRepo)
	userAccountApi := useraccountapi.NewUserAccountApi(*userAccountService)
	userAccountApi.Group(e.Group("user-accounts"))
	e.Start(":5003")
}
