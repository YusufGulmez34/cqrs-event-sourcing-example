package dto

type UserAccountDTO struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Surname       string `json:"sur_name"`
	TCKNO         string `json:"tckno"`
	AccountNumber string `json:"account_number"`
	Balance       uint   `json:"balance"`
	CurrencyType  string `json:"currency_type"`
}
