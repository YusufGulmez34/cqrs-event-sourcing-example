package useraccountrepo

import (
	"query-service/internal/event"

	"github.com/couchbase/gocb/v2"
)

type UserAccountRepository struct {
	scope *gocb.Scope
}

func NewUserAccountRepository(scope *gocb.Scope) *UserAccountRepository {
	return &UserAccountRepository{scope: scope}
}

func (u *UserAccountRepository) GetUserAccountEventListByAccountID(id uint) ([]event.UserAccountEvent, error) {
	result, err := u.scope.Query("SELECT ua.* FROM `my-db`.`test`.`useraccounts` ua WHERE ua.user_account_id = $1 ORDER BY ua.created_at DESC;", &gocb.QueryOptions{PositionalParameters: []interface{}{id}})
	if err != nil {
		return nil, err
	}
	var userAccountEventList []event.UserAccountEvent
	var userAccountEvent event.UserAccountEvent

	for result.Next() {
		if err := result.Row(&userAccountEvent); err != nil {
			return nil, err
		}
		userAccountEventList = append(userAccountEventList, userAccountEvent)
	}
	return userAccountEventList, nil
}
