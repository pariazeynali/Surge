package main

import (
	"log"
	"net/http"
	"surge_service/database"
	"time"
)

func httpHandler() {
	http.HandleFunc("/get-price-coefficient", GetCoefficient)
	http.HandleFunc("/ride-request", SaveRideRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func writeOSMData() {
	ticker := time.NewTicker(12 * time.Hour)
	select {
	case <-ticker.C:
		database.UpdateOSMData()
		database.SaveOSMToDB()
	}
}

func main() {
	database.ConnectToDb()
	database.CreateDBRelations()
	go writeOSMData()
	httpHandler()
}
