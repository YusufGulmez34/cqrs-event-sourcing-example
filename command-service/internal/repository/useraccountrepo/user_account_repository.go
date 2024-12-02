package useraccountrepo

import (
	"command-service/internal/model"

	"github.com/jinzhu/gorm"
)

type UserAccountRepository struct {
	db *gorm.DB
}

func NewUserAccountRepository(db *gorm.DB) *UserAccountRepository {
	return &UserAccountRepository{db: db}
}

func (u *UserAccountRepository) Migrate() error {
	return u.db.AutoMigrate(&model.UserAccount{}).Error
}

func (u *UserAccountRepository) CreateUserAccount(userAccountModel *model.UserAccount) error {
	return u.db.Create(&userAccountModel).Error
}

func (u *UserAccountRepository) FindUserAccountByAccountNo(accountNo string) (model.UserAccount, error) {
	var userAccountModel model.UserAccount
	err := u.db.Where("account_number=?", accountNo).First(&userAccountModel).Error
	return userAccountModel, err
}

func (u *UserAccountRepository) UpdateUserAccount(userAccountModel model.UserAccount) error {
	return u.db.Save(&userAccountModel).Error
}
