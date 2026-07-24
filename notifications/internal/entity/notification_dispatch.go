package entity

import "time"

type NotificationDispatch struct {
	Ksuid             string
	NotificationKsuid string
	Channel           Channel
	RecipientTarget   string
	Status            Status
	ProviderResponse  map[string]any
	RetryCount        int
	SentAt            time.Time
}

type CreateNotificationDispatch struct {
	NotificationKsuid string
	Channel           Channel
	RecipientTarget   string
	Status            Status
}

type UpdateNotificationDispatch struct {
	Ksuid            string
	Status           Status
	ProviderResponse map[string]any
	SentAt           time.Time
	RetryCount       int
}

type NotificationDispatchFilter struct {
	Status        *[]Status
	MaxRetryCount *int
}

type UpdateNotificationDispatchOption func(*UpdateNotificationDispatch)

func WithKsuid(ksuid string) func(*UpdateNotificationDispatch) {
	return func(und *UpdateNotificationDispatch) { und.Ksuid = ksuid }
}

func WithStatus(status Status) func(*UpdateNotificationDispatch) {
	return func(und *UpdateNotificationDispatch) { und.Status = status }
}

func WithSentAt(sentAt time.Time) func(*UpdateNotificationDispatch) {
	return func(und *UpdateNotificationDispatch) { und.SentAt = sentAt }
}

func WithRetryCount(retryCount int) func(*UpdateNotificationDispatch) {
	return func(und *UpdateNotificationDispatch) { und.RetryCount = retryCount }
}
