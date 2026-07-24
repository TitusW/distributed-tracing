package notification

import (
	"time"

	"github.com/TitusW/notifications/internal/entity"
)

type Notification struct {
	Ksuid            string
	IdempotencyKey   string
	UserKsuid        string
	EventType        entity.EventType
	TransactionKsuid string
	ContextData      map[string]any
	InsertedAt       time.Time
}
