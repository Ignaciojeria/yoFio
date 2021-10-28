package mapper

import (
	"yofio/app/adapter/out/inmemorydb/entity"
	"yofio/app/domain"
)

func MapToCreditAssignedEvent(credit domain.Credit) entity.CreditAssignedEvent {
	return entity.CreditAssignedEvent{}
}
