package entity

type Notification struct {
	Ksuid            string
	IdempotencyKey   string
	UserKsuid        string
	EventType        EventType
	TransactionKsuid string
	ContextData      map[string]any
}

type CreateNotification struct {
	IdempotencyKey   string
	UserKsuid        string
	EventType        EventType
	TransactionKsuid string
	ContextData      map[string]any
	Channels         []Channel
}
