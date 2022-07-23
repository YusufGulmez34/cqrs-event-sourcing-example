package dto

type ToInvestMoneyDTO struct {
	AccountNumber string `json:"account_number" validate:"required,len=16"`
	Amount        uint   `json:"amount" validate:"required"`
	CurrencyType  string `json:"currency_type"`
}
