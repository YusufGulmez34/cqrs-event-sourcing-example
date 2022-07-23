package useraccountservice

import (
	"command-service/internal/model"
	"command-service/internal/model/dto"
	"command-service/internal/model/event"
	"command-service/internal/repository/useraccountrepo"
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/copier"
	"github.com/streadway/amqp"
)

type UserAccountService struct {
	userAccountRepo useraccountrepo.UserAccountRepository
	rabbitMQConn    *amqp.Connection
}

func NewUserAccountService(userAccountRepo useraccountrepo.UserAccountRepository, rabbitMQConn *amqp.Connection) *UserAccountService {
	return &UserAccountService{userAccountRepo: userAccountRepo, rabbitMQConn: rabbitMQConn}
}

func (u *UserAccountService) CreateUserAccount(createUserAccountDTO dto.CreateUserAccountDTO) error {
	var userAccountModel model.UserAccount
	if err := copier.Copy(&userAccountModel, createUserAccountDTO); err != nil {
		return err
	}
	if err := u.userAccountRepo.CreateUserAccount(&userAccountModel); err != nil {
		return err
	}

	changedData, _ := json.Marshal(createUserAccountDTO)
	uuid, _ := uuid.NewV4()
	event := &event.UserAccountEvent{
		UUID:          uuid.String(),
		CreatedAt:     time.Now(),
		EventType:     "CREATED",
		ChangedData:   string(changedData),
		UserAccountID: userAccountModel.ID,
	}
	return u.AddEventToQueue(*event)
}

func (u *UserAccountService) ToInvestMoney(toInvestMoney dto.ToInvestMoneyDTO) error {
	userAccountModel, err := u.userAccountRepo.FindUserAccountByAccountNo(toInvestMoney.AccountNumber)
	if err != nil {
		return err
	}

	userAccountModel.Balance += toInvestMoney.Amount
	if err = u.userAccountRepo.UpdateUserAccount(userAccountModel); err != nil {
		return err
	}

	changedData, _ := json.Marshal(toInvestMoney)
	uuid, _ := uuid.NewV4()
	event := &event.UserAccountEvent{
		UUID:          uuid.String(),
		CreatedAt:     time.Now(),
		EventType:     "INVESTED",
		ChangedData:   string(changedData),
		UserAccountID: userAccountModel.ID,
	}
	return u.AddEventToQueue(*event)
}

func (u *UserAccountService) ToDrawMoney(toDrawMoney dto.ToDrawMoneyDTO) error {
	userAccountModel, err := u.userAccountRepo.FindUserAccountByAccountNo(toDrawMoney.AccountNumber)
	if err != nil {
		return err
	}

	userAccountModel.Balance -= toDrawMoney.Amount
	if err = u.userAccountRepo.UpdateUserAccount(userAccountModel); err != nil {
		return err
	}

	changedData, _ := json.Marshal(toDrawMoney)
	uuid, _ := uuid.NewV4()
	event := &event.UserAccountEvent{
		UUID:          uuid.String(),
		CreatedAt:     time.Now(),
		EventType:     "DRAWED",
		ChangedData:   string(changedData),
		UserAccountID: userAccountModel.ID,
	}
	return u.AddEventToQueue(*event)
}
func (u *UserAccountService) AddEventToQueue(event event.UserAccountEvent) error {
	ch, err := u.rabbitMQConn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"useraccount", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return err
	}

	body, _ := json.Marshal(event)
	return ch.Publish("", q.Name, false, false, amqp.Publishing{ContentType: "application/json", Body: body})
}
