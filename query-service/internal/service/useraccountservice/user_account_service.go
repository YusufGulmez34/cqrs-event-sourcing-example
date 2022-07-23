package useraccountservice

import (
	"encoding/json"
	"query-service/internal/dto"
	"query-service/internal/event"

	"query-service/internal/repository/useraccountrepo"
)

type UserAccountService struct {
	userAccountRepo useraccountrepo.UserAccountRepository
}

func NewUserAccountService(userAccountRepo useraccountrepo.UserAccountRepository) *UserAccountService {
	return &UserAccountService{userAccountRepo: userAccountRepo}
}

func (u *UserAccountService) GetUserAccountByAccountID(accountID uint) (*dto.UserAccountDTO, error) {
	eventList, err := u.userAccountRepo.GetUserAccountEventListByAccountID(accountID)
	if err != nil {
		return nil, err
	}

	var userAccountDTO dto.UserAccountDTO
	var changedEvent event.MoneyChange

	for _, e := range eventList {
		switch e.EventType {
		case "CREATED":
			if err := json.Unmarshal([]byte(e.ChangedData), &userAccountDTO); err != nil {
				return nil, err
			}
			userAccountDTO.ID = e.UserAccountID
			break
		case "INVESTED":
			if err := json.Unmarshal([]byte(e.ChangedData), &changedEvent); err != nil {
				return nil, err
			}
			userAccountDTO.Balance += changedEvent.Amount
			break
		case "DRAWED":
			if err := json.Unmarshal([]byte(e.ChangedData), &changedEvent); err != nil {
				return nil, err
			}
			userAccountDTO.Balance -= changedEvent.Amount
			break
		}
	}
	return &userAccountDTO, nil
}
