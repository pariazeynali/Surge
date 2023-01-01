package gis_helper

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

// Db connection pool
var conn *pgxpool.Pool

func ConnectToDb() {
	dbUser := os.Getenv("OSM_USER")
	dbPass := os.Getenv("PGPASSWORD")
	dbName := os.Getenv("OSM_DB")
	dbHost := os.Getenv("OSM_HOST")
	dbPort := os.Getenv("OSM_PORT")

	connStr := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v pool_max_conns=50", dbUser, dbPass, dbHost, dbPort, dbName)
	var err error
	conn, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
}

// GetMunicipalityDistrict Returns municipality district of every point
func GetMunicipalityDistrict(latitude string, longitude string) (int, error) {
	var polygonOSMId int
	// Get polygon within source point
	err := conn.QueryRow(context.Background(), "select tc.osm_id "+
		"from tehran tc where "+
		"ST_Contains(tc.way, ST_POINT($1::float, $2::float, 3857))", latitude, longitude).Scan(&polygonOSMId)
	if err != nil {
		return 0, err
	} else if polygonOSMId == 0 {
		return polygonOSMId, errors.New("polygon not found")
	} else {
		return polygonOSMId, nil
	}
}

// GetReqThreshold Returns source polygon's threshold demand count
func GetReqThreshold(OSMId int) int {
	var reqCount int
	err := conn.QueryRow(context.Background(), "select count from ride_req_statistics where osm_id = $1", OSMId).Scan(&reqCount)
	if err != nil {
		log.Panic(err)
	}
	return reqCount
}

// SaveReqData Save req source address coordinates and request timestamp in database
func SaveReqData(latitude string, longitude string) error {
	_, err := conn.Exec(context.Background(), "insert into ride_request (geom) values (ST_Point($1, $2, 3857))", latitude, longitude)
	return err
}

// GetPriceCoefficient Returns price coefficient of source polygon
func GetPriceCoefficient(threshold int) float32 {
	var coefficient float32
	err := conn.QueryRow(context.Background(), "select price_coefficient from coefficient_t where req_threshold::int4range @> $1::int4", threshold).Scan(&coefficient)
	if err != nil {
		log.Panic(err)
	}
	return coefficient
}
