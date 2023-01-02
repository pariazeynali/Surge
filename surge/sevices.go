package surge

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

// GetCoefficient provides rest api to get latitude and longitude of source address and returns price coefficient of thw area
// This api should call in each (get price) request
func GetCoefficient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var res Response
	latitude := r.URL.Query().Get("latitude")
	longitude := r.URL.Query().Get("longitude")
	log.Printf("latidtude: %v, longitude %v", latitude, longitude)
	coefficient := gis_helper.GetPriceCoefficient(latitude, longitude)
	res = Response{Status: 200, Result: fmt.Sprintf("%v", coefficient)}
	jsonRes, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}

// SaveRideRequest saves ride request data into database
// This app simulates a behavior similar to the ride request service
// when registering a request it saves the request information
// in the database so that the desired request data is stored in the database
// and can be used later
func SaveRideRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	latitude := r.URL.Query().Get("latitude")
	longitude := r.URL.Query().Get("longitude")
	if err := gis_helper.SaveReqData(latitude, longitude); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("ERROR: while saving ride request data: %v", err)
	}
	w.WriteHeader(http.StatusCreated)

}
