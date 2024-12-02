package event

type MoneyChange struct {
	AccountNumber string `json:"account_number"`
	Amount        uint   `json:"amount"`
	CurrencyType  string `json:"currency_type"`
}
