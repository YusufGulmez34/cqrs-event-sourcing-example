package model

import "github.com/jinzhu/gorm"

type UserAccount struct {
	gorm.Model
	Name          string `gorm:"not null;type:varchar(25)"`
	Surname       string `gorm:"not null;type:varchar(25)"`
	TCKNO         string `gorm:"not null;type:varchar(11)"`
	AccountNumber string `gorm:"not null;type:varchar(16)"`
	Balance       uint   `gorm:"default=0"`
	CurrencyType  string `gorm:"type:varchar(3);default=TRY"`
}
