package database

import (
	"context"
)

// SaveReqData Save req source address coordinates and request timestamp in database
func SaveReqData(latitude string, longitude string) error {
	ctx := context.Background()
	tx, err := Conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	_, err = tx.Exec(context.Background(), "insert into ride_request (geom) values (ST_Point($1, $2, 3857))", latitude, longitude)
	err = tx.Commit(ctx)
	return err
}
