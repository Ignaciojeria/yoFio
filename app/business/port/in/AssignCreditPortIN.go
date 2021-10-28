package in

import "yofio/app/domain"

type AssignCreditPortIN func(investment domain.Money) ([]domain.Credit, error)
