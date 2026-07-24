package entity

type Status string

const (
	StatusPending    Status = "PENDING"
	StatusProcessing Status = "PROCESSING"
	StatusSent       Status = "SENT"
	StatusDelivered  Status = "DELIVERED"
	StatusFailed     Status = "FAILED"
	StatusDead       Status = "DEAD"
)
