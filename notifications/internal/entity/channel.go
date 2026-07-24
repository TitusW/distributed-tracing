package entity

type Channel string

const (
	InApp Channel = "IN_APP"
	Sms   Channel = "SMS"
	Email Channel = "EMAIL"
	Push  Channel = "PUSH"
)
