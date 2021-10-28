package out

import "yofio/app/domain"

type SaveCreditAssignedEventPortOUT func(credits []domain.Credit) error
