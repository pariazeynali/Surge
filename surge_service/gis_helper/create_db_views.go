package gis_helper

import (
	"context"
	"log"
	"os"
)

func CreateDBViews() {
	sql, err := os.ReadFile("./views.sql")
	if err != nil {
		log.Fatalf("ERROR reading SQl files: %v", err)
	}
	sqlQ := string(sql)
	c, _ := Conn.Acquire(context.Background())
	defer c.Release()
	_, err = c.Exec(context.Background(), sqlQ)
	if err != nil {
		log.Fatalf("ERROR could not create views: %v", err)
	}
}
