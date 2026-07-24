package entity

type EventType string

const (
	PaymentInitiated        EventType = "PAYMENT.INITIATED"
	PaymentDebited          EventType = "PAYMENT.DEBITED"
	PaymentCredited         EventType = "PAYMENT.CREDITED"
	PaymentDebitFailed      EventType = "PAYMENT.DEBIT.FAILED"
	PaymentCreditFailed     EventType = "PAYMENT.CREDIT.FAILED"
	PaymentDebitCompensated EventType = "PAYMENT.DEBIT.COMPENSATED"
)
