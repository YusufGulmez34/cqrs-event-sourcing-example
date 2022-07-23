package useraccountapi

import (
	"command-service/internal/model/dto"
	"command-service/internal/service/useraccountservice"

	"github.com/labstack/echo/v4"
)

type UserAccountApi struct {
	userAccountService useraccountservice.UserAccountService
}

func NewUserAccountApi(userAccountService useraccountservice.UserAccountService) *UserAccountApi {
	return &UserAccountApi{userAccountService: userAccountService}
}

func (u *UserAccountApi) Group(g *echo.Group) {
	g.POST("", u.CreateUserAccount)
	g.PUT("/invest", u.ToInvestMoney)
	g.PUT("/draw", u.ToDrawMoney)

}

// @Tags UserAccountApi
// @Security ApiKeyAuth
// @Summary Create User Account
// @Param data body dto.CreateUserAccountDTO true "CreateUserAccount DTO"
// @Success 201 string Success "{"success":true,"msg":"Success"}"
// @Router /user-accounts [post]
func (u *UserAccountApi) CreateUserAccount(c echo.Context) error {
	var createUserAccountDTO dto.CreateUserAccountDTO
	if err := c.Bind(&createUserAccountDTO); err != nil {
		return c.JSON(400, err.Error())
	}
	if err := c.Validate(&createUserAccountDTO); err != nil {
		return c.JSON(400, err.Error())
	}
	if err := u.userAccountService.CreateUserAccount(createUserAccountDTO); err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(201, "Created Successfuly")
}

// @Tags UserAccountApi
// @Security ApiKeyAuth
// @Summary Invest Money to User Account
// @Param data body dto.ToInvestMoneyDTO true "To Invest Money DTO"
// @Success 200 string Success "{"success":true,"msg":"Success"}"
// @Router /user-accounts/invest [put]
func (u *UserAccountApi) ToInvestMoney(c echo.Context) error {
	var toInvestMoneyDTO dto.ToInvestMoneyDTO
	if err := c.Bind(&toInvestMoneyDTO); err != nil {
		return c.JSON(400, err.Error())
	}
	if err := c.Validate(&toInvestMoneyDTO); err != nil {
		return c.JSON(400, err.Error())
	}
	if err := u.userAccountService.ToInvestMoney(toInvestMoneyDTO); err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, "Successfuly")
}

// @Tags UserAccountApi
// @Security ApiKeyAuth
// @Summary Draw Money from User Account
// @Param data body dto.ToDrawMoneyDTO true "To Draw Money DTO"
// @Success 200 string Success "{"success":true,"msg":"Success"}"
// @Router /user-accounts/draw [put]
func (u *UserAccountApi) ToDrawMoney(c echo.Context) error {
	var toDrawMoneyDTO dto.ToDrawMoneyDTO
	if err := c.Bind(&toDrawMoneyDTO); err != nil {
		return c.JSON(400, err.Error())
	}
	if err := c.Validate(&toDrawMoneyDTO); err != nil {
		return c.JSON(400, err.Error())
	}
	if err := u.userAccountService.ToDrawMoney(toDrawMoneyDTO); err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, "Successfuly")
}
