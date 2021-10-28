package mapper

import (
	"yofio/app/adapter/in/web/response"
	"yofio/app/domain"
)

func MapToAssignCreditResponse(credits []domain.Credit) []response.AssignCreditResponse {

	r := make([]response.AssignCreditResponse, 0)

	for _, credit := range credits {
		r = append(r, response.AssignCreditResponse{
			CreditAmount: credit.CreditAmount.Value,
			Quantity:     credit.Quantity,
			Total:        credit.Total.Value,
			CreditType:   string(credit.CreditType),
		})
	}

	return r
}
