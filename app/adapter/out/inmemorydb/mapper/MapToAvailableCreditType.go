package mapper

import (
	"yofio/app/adapter/out/inmemorydb/entity"
	"yofio/app/domain"
)

func MapToAvailableCreditType(credit entity.AvailableCreditType) domain.Credit {
	return domain.Credit{
		CreditID: domain.CreditID{
			Value: credit.CreditID,
		},
		CreditAmount: domain.Money{
			Value: credit.CreditAmount,
		},
		CreditType: domain.CreditType(credit.CreditType),
	}
}
