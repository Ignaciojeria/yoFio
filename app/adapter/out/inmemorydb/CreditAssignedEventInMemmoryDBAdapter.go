package inmemorydb

import (
	"sync"
	"yofio/app/adapter/out/inmemorydb/entity"
	"yofio/app/domain"
)

type CreditAssignedEventInMemmoryDBAdapter struct {
	mu sync.Mutex
}

var inMemmoryCreditAssignedEventList []entity.CreditAssignedEvent = make([]entity.CreditAssignedEvent, 0)

func (a *CreditAssignedEventInMemmoryDBAdapter) SaveCreditAssignedEventPortOUT(credits []domain.Credit) error {
	a.mu.Lock()
	for _, v := range credits {
		inMemmoryCreditAssignedEventList = append(inMemmoryCreditAssignedEventList, entity.CreditAssignedEvent{
			CreditID:     v.CreditID.Value,
			CreditAmount: v.CreditAmount.Value,
			Quantity:     v.Quantity,
			Total:        v.Total.Value,
			CreditType:   string(v.CreditType),
		})
	}
	a.mu.Unlock()
	return nil
}
