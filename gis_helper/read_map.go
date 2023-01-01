package gis_helper

import (
	"context"
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

// SaveReqData Save req source address coordinates and request timestamp in database
func SaveReqData(latitude string, longitude string) error {
	_, err := conn.Exec(context.Background(), "insert into ride_request (geom) values (ST_Point($1, $2, 3857))", latitude, longitude)
	return err
}

// GetPriceCoefficient gets source coordinates and returns price coefficient of the area
func GetPriceCoefficient(latitude string, longitude string) float32 {
	ctx := context.Background()
	tx, err := conn.Begin(ctx)
	if err != nil {
		log.Panicf("could not start transaction: %v", err)
	}
	defer tx.Commit(ctx)
	query := "select tc.osm_id from queries.tehran tc where ST_Contains(tc.way, ST_POINT($1::float, $2::float, 3857))"
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
