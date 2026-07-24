package notificationdispatch

import (
	"time"

	"github.com/TitusW/notifications/internal/entity"
)

type NotificationDispatch struct {
	Ksuid             string
	NotificationKsuid string
	Channel           entity.Channel
	RecipientTarget   string
	Status            entity.Status
	ProviderResponse  map[string]any
	RetryCount        int
	SentAt            time.Time
	InsertedAt        time.Time
}
