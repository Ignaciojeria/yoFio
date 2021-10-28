package usecase

import (
	"errors"
	"yofio/app/business/port/out"
	"yofio/app/domain"
	"yofio/app/utils"
)

type AssignCreditUseCase struct {
	out.LoadCreditTypeListPortOUT
	out.LoadStatisticsPortOUT
	out.UpdateStatisticsPortOUT
}

func (u AssignCreditUseCase) AssignCredit(investment domain.Money) ([]domain.Credit, error) {
	
	isDivisibleBy100 := utils.CheckDivisibleBy(investment.Value,100)

	if !isDivisibleBy100 {
		return nil,errors.New("error distributing investment")
	}

	if investment.Value<300{
		return nil,errors.New("error distributing investment")
	}

	creditTypesMap, _ := u.loadCreditTyeListAsMap()
	switch investment.Value {
	case creditTypesMap[domain.Credit300].CreditAmount.Value :
		return []domain.Credit{u.buildCreditAndSetQuantity(creditTypesMap[domain.Credit300],1)},nil
	case creditTypesMap[domain.Credit500].CreditAmount.Value :
		return []domain.Credit{u.buildCreditAndSetQuantity(creditTypesMap[domain.Credit500],1)},nil
	case creditTypesMap[domain.Credit700].CreditAmount.Value :
		return []domain.Credit{u.buildCreditAndSetQuantity(creditTypesMap[domain.Credit700],1)},nil
	case 600 :
		return []domain.Credit{u.buildCreditAndSetQuantity(creditTypesMap[domain.Credit300],2)},nil
	case 1200 :
		return []domain.Credit{
			u.buildCreditAndSetQuantity(creditTypesMap[domain.Credit500],1),
			u.buildCreditAndSetQuantity(creditTypesMap[domain.Credit700],1)},nil
	}

	credits := u.distributeCredits(investment)

	totalAmount := u.calculateCreditsTotalAmount(credits)

	if totalAmount != investment.Value{
		return nil,errors.New("error distributing investment")
	}

	return u.distributeSimplifiedInvestmentCredits(credits), nil
}

func (u AssignCreditUseCase) calculateCreditsTotalAmount(
	credits map[domain.CreditType]domain.Credit) int32 {
	var totalAmount int32 = 0 
	for _, v := range credits {
		totalAmount += v.Total.Value
	}
	return totalAmount
}


func (u AssignCreditUseCase) distributeSimplifiedInvestmentCredits(
	credits map[domain.CreditType]domain.Credit) []domain.Credit {
	simplifiedCredits := make([]domain.Credit,0)
	for _, v := range credits {
		simplifiedCredits = append(simplifiedCredits, v)
	}

	
	return simplifiedCredits
}

func (u AssignCreditUseCase) loadCreditTyeListAsMap() (map[domain.CreditType]domain.Credit, error) {
	credits,err := u.LoadCreditTypeListPortOUT()

	if err != nil{
		return nil,err
	}

	creditsMap := make(map[domain.CreditType]domain.Credit,3)

	for _, credit := range credits {
		creditsMap[credit.CreditType] = credit
	}

	return creditsMap, nil
}



func (u AssignCreditUseCase) buildCreditAndSetQuantity(c domain.Credit, quantity int32) domain.Credit{
	total := quantity*c.CreditAmount.Value
	credit := domain.Credit{
		CreditAmount: domain.Money{
			Value: c.CreditAmount.Value,
		},
		Quantity: quantity,
		Total: domain.Money{
			Value: total,
		},
		CreditType: c.CreditType,
	}


	return credit
}


func (u AssignCreditUseCase) distributeCredits(investment domain.Money) map[domain.CreditType]domain.Credit{
	creditTypesMap, _ := u.loadCreditTyeListAsMap()

	creditType300 := creditTypesMap[domain.Credit300]
	creditType500 := creditTypesMap[domain.Credit500]
	creditType700 := creditTypesMap[domain.Credit700]

	minus500Pattern := investment.Value - creditType500.CreditAmount.Value
	minus500Minus300Pattern := investment.Value - creditType500.CreditAmount.Value - creditType300.CreditAmount.Value
	minus500Minus300Minus300Pattern := investment.Value - creditType500.CreditAmount.Value - (creditType300.CreditAmount.Value*2)
	minus300Minus300Minus300Pattern := investment.Value - (creditType300.CreditAmount.Value * 3)
	minus300Minus500Minus500Pattern := investment.Value - creditType300.CreditAmount.Value - (creditType500.CreditAmount.Value*2)
	minus500Minus500Pattern := investment.Value - (creditType500.CreditAmount.Value * 2)

	var remainingAmountDivisibleBy700 int32 = 0
	simplifiedCredits := make(map[domain.CreditType]domain.Credit, 0)
	switch true {
	case utils.CheckDivisibleBy(minus500Minus300Pattern,creditType700.CreditAmount.Value):
		simplifiedCredits[domain.Credit500] = u.buildCreditAndSetQuantity(creditTypesMap[domain.Credit500],1)
		simplifiedCredits[domain.Credit300] = u.buildCreditAndSetQuantity(creditTypesMap[domain.Credit300],1)
		remainingAmountDivisibleBy700 = minus500Minus300Pattern
		break
	case utils.CheckDivisibleBy(minus300Minus300Minus300Pattern,creditType700.CreditAmount.Value):
		simplifiedCredits[domain.Credit300] = u.buildCreditAndSetQuantity(creditTypesMap[domain.Credit300],3)
		remainingAmountDivisibleBy700 = minus500Minus300Pattern
		break
	case utils.CheckDivisibleBy(minus500Minus500Pattern,creditType700.CreditAmount.Value):
		simplifiedCredits[domain.Credit500] = u.buildCreditAndSetQuantity(creditTypesMap[domain.Credit500],2)
		remainingAmountDivisibleBy700 = minus500Minus500Pattern
		break
	case utils.CheckDivisibleBy(minus500Pattern,creditType700.CreditAmount.Value):
		simplifiedCredits[domain.Credit500] = u.buildCreditAndSetQuantity(creditTypesMap[domain.Credit500],1)
		remainingAmountDivisibleBy700 = minus500Pattern
		break
	case utils.CheckDivisibleBy(minus500Minus300Minus300Pattern,creditType700.CreditAmount.Value):
		simplifiedCredits[domain.Credit500] = u.buildCreditAndSetQuantity(creditTypesMap[domain.Credit500],1)
		simplifiedCredits[domain.Credit300] = u.buildCreditAndSetQuantity(creditTypesMap[domain.Credit300],2)
		remainingAmountDivisibleBy700 = minus500Minus500Pattern
		break
	case utils.CheckDivisibleBy(minus300Minus500Minus500Pattern,creditType700.CreditAmount.Value):
		simplifiedCredits[domain.Credit300] = u.buildCreditAndSetQuantity(creditTypesMap[domain.Credit300],1)
		simplifiedCredits[domain.Credit500] = u.buildCreditAndSetQuantity(creditTypesMap[domain.Credit500],2)
		remainingAmountDivisibleBy700 = minus500Minus500Pattern
	default :
		remainingAmountDivisibleBy700 = investment.Value
	}

	redistributedCredits := u.distributeCreditsInvestmentRemainingAmount(simplifiedCredits,investment,remainingAmountDivisibleBy700)

	if redistributedCredits!= nil {
		return redistributedCredits
	}

	credit700Quantity := remainingAmountDivisibleBy700/creditType700.CreditAmount.Value
	simplifiedCredits[domain.Credit700]  = u.buildCreditAndSetQuantity(creditTypesMap[domain.Credit700],credit700Quantity)

	return simplifiedCredits
}

func (u AssignCreditUseCase) distributeCreditsInvestmentRemainingAmount(
	 simplifiedCredits map[domain.CreditType]domain.Credit,
	 investment domain.Money,
	 remainingAmountDivisibleBy700 int32) map[domain.CreditType]domain.Credit{

	sumPatterns := remainingAmountDivisibleBy700 / 2100

	if sumPatterns>0 {

	creditTypesMap, _ := u.loadCreditTyeListAsMap()

	simplifiedCredits[domain.Credit300] = u.buildCreditAndSetQuantity(
		creditTypesMap[domain.Credit300],
		simplifiedCredits[domain.Credit300].Quantity+sumPatterns*3)
	
	simplifiedCredits[domain.Credit500] = u.buildCreditAndSetQuantity(
		creditTypesMap[domain.Credit500],
		simplifiedCredits[domain.Credit500].Quantity+sumPatterns)
		
	simplifiedCredits[domain.Credit700] = u.buildCreditAndSetQuantity(
		creditTypesMap[domain.Credit700],
		simplifiedCredits[domain.Credit700].Quantity+sumPatterns)
	
	remainingAmount := investment.Value - (simplifiedCredits[domain.Credit300].Total.Value +
	simplifiedCredits[domain.Credit500].Total.Value + simplifiedCredits[domain.Credit700].Total.Value)

	if remainingAmount != 0 {
		remainingAmount -= creditTypesMap[domain.Credit700].CreditAmount.Value
		simplifiedCredits[domain.Credit700] = u.buildCreditAndSetQuantity(
			creditTypesMap[domain.Credit700],
			simplifiedCredits[domain.Credit700].Quantity+1)
	}

	if remainingAmount == creditTypesMap[domain.Credit700].CreditAmount.Value {
		remainingAmount -= creditTypesMap[domain.Credit700].CreditAmount.Value
		simplifiedCredits[domain.Credit700] = u.buildCreditAndSetQuantity(
			creditTypesMap[domain.Credit700],
			simplifiedCredits[domain.Credit700].Quantity+1)
	}

	if remainingAmount != 0 {
		remainingAmount -= creditTypesMap[domain.Credit500].CreditAmount.Value
		simplifiedCredits[domain.Credit500] = u.buildCreditAndSetQuantity(
			creditTypesMap[domain.Credit500],
			simplifiedCredits[domain.Credit500].Quantity+1)
	}

	for i := 0; i < 3; i++ {
		if remainingAmount != 0 {
			remainingAmount -= creditTypesMap[domain.Credit300].CreditAmount.Value
			simplifiedCredits[domain.Credit300] = u.buildCreditAndSetQuantity(
				creditTypesMap[domain.Credit300],
				simplifiedCredits[domain.Credit300].Quantity+1)
		}
	}
	simplifiedCredits = u.creditsInvestmentRemainingAmountRedistribution(simplifiedCredits)
	return simplifiedCredits

	}

	return nil

}

func (u AssignCreditUseCase) creditsInvestmentRemainingAmountRedistribution(
	simplifiedCredits map[domain.CreditType]domain.Credit) map[domain.CreditType]domain.Credit{

	creditTypesMap, _ := u.loadCreditTyeListAsMap()

	credit300TotalAmountValue := simplifiedCredits[domain.Credit300].Total.Value

	creditType300AmountValue := creditTypesMap[domain.Credit300].CreditAmount.Value
	minus300Minus300Pattern := credit300TotalAmountValue - (creditType300AmountValue *2)
	
	minus300Minus300Minus300Pattern := credit300TotalAmountValue - (creditType300AmountValue *3)

	minus300Minus300Minus300Minus300Pattern := credit300TotalAmountValue - (creditType300AmountValue *4)

	switch true {

	case utils.CheckDivisibleBy(credit300TotalAmountValue,1800):
		
		matchPatternsQuantity := credit300TotalAmountValue/1800
		
		simplifiedCredits[domain.Credit300] = u.
		buildCreditAndSetQuantity(creditTypesMap[domain.Credit300],matchPatternsQuantity*2)
		
		simplifiedCredits[domain.Credit500] =  u.
		buildCreditAndSetQuantity(creditTypesMap[domain.Credit500],
			simplifiedCredits[domain.Credit500].
			Quantity+matchPatternsQuantity)
	
		simplifiedCredits[domain.Credit700] =  u.
		buildCreditAndSetQuantity(creditTypesMap[domain.Credit700],
			simplifiedCredits[domain.Credit700].
			Quantity+matchPatternsQuantity)
		return simplifiedCredits
		
	case utils.CheckDivisibleBy(minus300Minus300Pattern,1800):
		matchPatternsQuantity := minus300Minus300Pattern/1800

		simplifiedCredits[domain.Credit300] = u.
		buildCreditAndSetQuantity(creditTypesMap[domain.Credit300],(matchPatternsQuantity*2)+2)
		
		simplifiedCredits[domain.Credit500] =  u.
		buildCreditAndSetQuantity(creditTypesMap[domain.Credit500],
			simplifiedCredits[domain.Credit500].
			Quantity+matchPatternsQuantity)

		simplifiedCredits[domain.Credit700] =  u.
		buildCreditAndSetQuantity(creditTypesMap[domain.Credit700],
			simplifiedCredits[domain.Credit700].
			Quantity+matchPatternsQuantity)

	return simplifiedCredits

	case utils.CheckDivisibleBy(minus300Minus300Minus300Pattern,1800):
		
		matchPatternsQuantity := minus300Minus300Minus300Pattern/1800
		
		simplifiedCredits[domain.Credit300] = u.
		buildCreditAndSetQuantity(creditTypesMap[domain.Credit300],(matchPatternsQuantity*2)+3)
		
		simplifiedCredits[domain.Credit500] =  u.
		buildCreditAndSetQuantity(creditTypesMap[domain.Credit500],
			simplifiedCredits[domain.Credit500].
			Quantity+matchPatternsQuantity)

		simplifiedCredits[domain.Credit700] =  u.
		buildCreditAndSetQuantity(creditTypesMap[domain.Credit700],
			simplifiedCredits[domain.Credit700].
			Quantity+matchPatternsQuantity)

	case utils.CheckDivisibleBy(minus300Minus300Minus300Minus300Pattern,1800):
		
		matchPatternsQuantity := minus300Minus300Minus300Pattern/1800
		
		simplifiedCredits[domain.Credit300] = u.
		buildCreditAndSetQuantity(creditTypesMap[domain.Credit300],(matchPatternsQuantity*2)+4)
		
		simplifiedCredits[domain.Credit500] =  u.
		buildCreditAndSetQuantity(creditTypesMap[domain.Credit500],
			simplifiedCredits[domain.Credit500].
			Quantity+matchPatternsQuantity)

		simplifiedCredits[domain.Credit700] =  u.
		buildCreditAndSetQuantity(creditTypesMap[domain.Credit700],
			simplifiedCredits[domain.Credit700].
			Quantity+matchPatternsQuantity)

	return simplifiedCredits
	
	}
	
	return simplifiedCredits
}

