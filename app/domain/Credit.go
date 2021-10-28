package domain

type CreditID struct {
	Value string
}

type Money struct {
	Value int32
}

type Credit struct {
	CreditID     CreditID
	CreditAmount Money
	Quantity     int32
	Total        Money
	CreditType   CreditType
}

type CreditType string

const (
	Credit300 CreditType = "CREDIT_300"
	Credit500 CreditType = "CREDIT_500"
	Credit700 CreditType = "CREDIT_700"
)