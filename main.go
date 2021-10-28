package main

import (
	"log"
	"net/http"
	"os"
	"yofio/app/adapter/in/web"
	"yofio/app/adapter/out/inmemorydb"
	"yofio/app/business/usecase"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	a1 := inmemorydb.AvailableCreditTypeInMemoryDBAdapter{}
	a2 := inmemorydb.StatisticsInMemoryDBAdapter{}

	assingCreditUseCase := usecase.AssignCreditUseCase{
		LoadCreditTypeListPortOUT: a1.LoadCreditTypeListPortOUT,
		LoadStatisticsPortOUT:     a2.LoadStatisticsPortOUT,
		UpdateStatisticsPortOUT:   a2.UpdateStatisticsPortOUT,
	}.AssignCredit

	/*
	assingCreditUseCase(domain.Money{
		Value: 2800,
	})
	*/

	
	assignCreditHandler := web.AssignCreditHandle{
		AssignCreditPortIN: assingCreditUseCase,
	}

	ShowStatisticsHandle := web.ShowStatisticsHandle{
		LoadStatisticsPortOUT: a2.LoadStatisticsPortOUT,
	}
 

	http.HandleFunc("/credit-assignement", assignCreditHandler.AssignCredit)
	http.HandleFunc("/statistics", ShowStatisticsHandle.ShowStatistics)

	serverPort := os.Getenv("server.port")
	log.Println("starting server on port: " + serverPort)
	err = http.ListenAndServe(":"+serverPort, nil)

	if err != nil {
		log.Fatal("Error starting server")
	}

	
	/*
	credits,_ := assingCreditUseCase(domain.Money{
		Value: 2000,
	})

	b,_:= json.Marshal(credits)


	log.Println(string(b))
	*/

}
