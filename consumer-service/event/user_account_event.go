package event

import "time"

type UserAccountEvent struct {
	UUID          string    `json:"uuid"`
	CreatedAt     time.Time `json:"created_at"`
	EventType     string    `json:"event_type"`
	ChangedData   string    `json:"changed_data"`
	UserAccountID uint      `json:"user_account_id"`
}
