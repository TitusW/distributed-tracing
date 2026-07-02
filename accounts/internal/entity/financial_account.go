package entity

type FinancialAccount struct {
	Ksuid         string
	CurrentAmount float64
}

type DebitFinancialAccount struct {
	Ksuid  string
	Amount float64
}

type CreditFinancialAccount struct {
	Ksuid  string
	Amount float64
}
