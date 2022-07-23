package dto

type CreateUserAccountDTO struct {
	Name          string `json:"name" validate:"required,max=25"`
	Surname       string `json:"sur_name" validate:"required,max=25"`
	TCKNO         string `json:"tckno" validate:"required,len=11"`
	AccountNumber string `json:"account_number" validate:"required,len=16"`
}
