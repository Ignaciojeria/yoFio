package usecase

import (
	"testing"
	"yofio/app/domain"

	"github.com/google/uuid"
)


func TestLoadCreditTyeListAsMap(t *testing.T) {

	credit300 := domain.Credit{
		CreditID:    domain.CreditID{
			Value:  uuid.NewString(),
		},
		CreditAmount: domain.Money{
			Value: 300,
		},
		CreditType:   domain.Credit300,
	}

	credit500 := domain.Credit{
		CreditID:    domain.CreditID{
			Value:  uuid.NewString(),
		},
		CreditAmount: domain.Money{
			Value: 500,
		},
		CreditType:   domain.Credit500,
	}

	credit700 := domain.Credit{
		CreditID:    domain.CreditID{
			Value:  uuid.NewString(),
		},
		CreditAmount: domain.Money{
			Value: 700,
		},
		CreditType:   domain.Credit700,
	}

	loadCreditTypeList := func() ([]domain.Credit, error) {
		availableCredits := make([]domain.Credit, 3)

		availableCredits = append(availableCredits, credit300)
		
	
		availableCredits = append(availableCredits, credit500)
	
		availableCredits = append(availableCredits, credit700)
		return availableCredits,nil
	}

	usecase := AssignCreditUseCase{
	LoadCreditTypeListPortOUT:  loadCreditTypeList,
	}

	creditTypeMap, err := usecase.loadCreditTyeListAsMap()

	if err != nil {
		t.Errorf("got %v, want %v", err, "nil")
	}


	if credit300 !=  creditTypeMap[domain.Credit300]{
		t.Errorf("got %v, want %v", credit300, creditTypeMap[domain.Credit300])
	}

	if credit500 !=  creditTypeMap[domain.Credit500]{
		t.Errorf("got %v, want %v", credit300, creditTypeMap[domain.Credit300])
	}

	if credit700 !=  creditTypeMap[domain.Credit700]{
		t.Errorf("got %v, want %v", credit300, creditTypeMap[domain.Credit300])
	}

}
