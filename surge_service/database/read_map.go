package database

import (
	"context"
	"log"
)

// GetPriceCoefficient gets source coordinates and returns price coefficient of the area
func GetPriceCoefficient(latitude string, longitude string) float32 {
	ctx := context.Background()
	tx, err := Conn.Begin(ctx)
	if err != nil {
		log.Panicf("could not start transaction: %v", err)
	}
	defer tx.Commit(ctx)
	query := "select tc.osm_id from tehran tc where ST_Contains(tc.way, ST_POINT($1::float, $2::float, 3857))"
	var polygonOSMId int
	if err = tx.QueryRow(ctx, query, latitude, longitude).Scan(&polygonOSMId); err != nil {
		log.Panicf("Error reading source polygon: %v", err)
	}
	query = "select count from ride_req_statistics where osm_id = $1"
	var reqCount int
	if err = tx.QueryRow(ctx, query, polygonOSMId).Scan(&reqCount); err != nil {
		log.Panicf("Error reading request count: %v\n", err)
	}
	var coefficient float32
	query = "select price_coefficient from coefficient_t where req_threshold::int4range @> $1::int4"
	if err = tx.QueryRow(ctx, query, reqCount).Scan(&coefficient); err != nil {
		log.Panicf("Error getting price coefficient: %v", err)
	}
	return coefficient
}
