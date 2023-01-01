package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"surge/gis_helper"
)

type Response struct {
	Status int    `json:"status_code"`
	Result string `json:"result"`
}

func GetCoefficient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var res Response
	latitude := r.URL.Query().Get("latitude")
	longitude := r.URL.Query().Get("longitude")
	log.Printf("latidtude: %v, longitude %v", latitude, longitude)
	myPolygon, err := gis_helper.GetMunicipalityDistrict(latitude, longitude)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Bad request : %v ", err)))
	}
	reqCount := gis_helper.GetReqThreshold(myPolygon)
	go func(x string, y string) {
		err = gis_helper.SaveReqData(x, y)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Bad request : %v ", err)))
		}
	}(latitude, longitude)
	coefficient := gis_helper.GetPriceCoefficient(reqCount)
	res = Response{Status: 200, Result: fmt.Sprintf("%v", coefficient)}
	jsonRes, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}
