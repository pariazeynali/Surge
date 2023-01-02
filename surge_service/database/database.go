package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

var Conn *pgxpool.Pool

func ConnectToDb() {
	dbUser := os.Getenv("OSM_USER")
	dbPass := os.Getenv("OSM_PASS")
	dbName := os.Getenv("OSM_DB")
	dbHost := os.Getenv("OSM_HOST")
	dbPort := os.Getenv("OSM_PORT")

	connStr := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v pool_max_conns=50", dbUser, dbPass, dbHost, dbPort, dbName)
	var err error
	Conn, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
}

func CreateDBRelations() {
	sql, err := os.ReadFile("./relation.sql")
	if err != nil {
		log.Fatalf("ERROR reading SQl files: %v", err)
	}
	sqlQ := string(sql)
	c, err := Conn.Acquire(context.Background())
	if err != nil {
		log.Panicf("ERROR accuring connection: %v", err)
	}
	defer c.Release()
	_, err = c.Exec(context.Background(), sqlQ)
	if err != nil {
		log.Fatalf("ERROR could not create relatiosn: %v", err)
	}
}
