package entity

type CreditAssignedEvent struct {
	CreditID     string
	CreditAmount int32
	Quantity     int32
	Total        int32
	CreditType   string
}
