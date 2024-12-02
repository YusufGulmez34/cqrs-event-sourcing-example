package useraccountapi

import (
	"query-service/internal/service/useraccountservice"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserAccountApi struct {
	userAccountService useraccountservice.UserAccountService
}

func NewUserAccountApi(userAccountService useraccountservice.UserAccountService) *UserAccountApi {
	return &UserAccountApi{userAccountService: userAccountService}
}

func (u *UserAccountApi) Group(g *echo.Group) {
	g.GET("/:id", u.GetUserAccountByID)
}

// @Tags UserAccountApi
// @Security ApiKeyAuth
// @Summary Get User Account
// @Param id path string true "User Account ID"
// @Success 201 string Success "{"success":true,"data":"","msg":"Success"}"
// @Router /user-accounts/{id} [get]
func (u *UserAccountApi) GetUserAccountByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, "Validation error: "+err.Error())
	}
	userAccountDTO, err := u.userAccountService.GetUserAccountByAccountID(uint(id))
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, map[string]any{"user_account": userAccountDTO, "message": "successfuly"})
}
