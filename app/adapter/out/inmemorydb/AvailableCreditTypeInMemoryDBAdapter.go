package inmemorydb

import (
	"yofio/app/adapter/out/inmemorydb/entity"
	"yofio/app/adapter/out/inmemorydb/mapper"
	"yofio/app/domain"

	"github.com/google/uuid"
)

type AvailableCreditTypeInMemoryDBAdapter struct {
}

func (a AvailableCreditTypeInMemoryDBAdapter) LoadCreditTypeListPortOUT() ([]domain.Credit, error) {

	availableCredits := make([]entity.AvailableCreditType, 3)

	availableCredits = append(availableCredits, entity.AvailableCreditType{
		CreditID:     uuid.NewString(),
		CreditAmount: 300,
		CreditType:   string(domain.Credit300),
	})

	availableCredits = append(availableCredits, entity.AvailableCreditType{
		CreditID:     uuid.NewString(),
		CreditAmount: 500,
		CreditType:   string(domain.Credit500),
	})

	availableCredits = append(availableCredits, entity.AvailableCreditType{
		CreditID:     uuid.NewString(),
		CreditAmount: 700,
		CreditType:   string(domain.Credit700),
	})

	d := make([]domain.Credit, 3)

	for _, v := range availableCredits {
		d = append(d, mapper.MapToAvailableCreditType(v))
	}

	return d, nil
}
