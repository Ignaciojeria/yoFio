package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"yofio/app/adapter/in/web/mapper"
	"yofio/app/adapter/in/web/request"
	"yofio/app/business/port/in"
	"yofio/app/domain"
)

type AssignCreditHandle struct {
	in.AssignCreditPortIN
}

func (h AssignCreditHandle) AssignCredit(w http.ResponseWriter, req *http.Request) {

	b, _ := ioutil.ReadAll(req.Body)

	var r request.AssignCreditRequest

	json.Unmarshal(b, &r)

	investment := domain.Money{
		Value: r.Investment,
	}

	credits, err := h.AssignCreditPortIN(investment)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mapper.MapToAssignCreditResponse(credits))
}
