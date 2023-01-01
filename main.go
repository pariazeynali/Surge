package main

import (
	"log"
	"net/http"
	"surge/gis_helper"
	"time"
)

func httpHandler() {
	http.HandleFunc("/get-price-coefficient", GetCoefficient)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func writeOSMData() {
	ticker := time.NewTicker(12 * time.Hour)
	select {
	case <-ticker.C:
		gis_helper.UpdateOSMData()
		gis_helper.SaveOSMToDB()
	}
}

func main() {
	gis_helper.ConnectToDb()
	go writeOSMData()
	httpHandler()
}
