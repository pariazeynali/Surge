package database

import "context"

// SaveReqData Save req source address coordinates and request timestamp in database
func SaveReqData(latitude string, longitude string) error {
	_, err := Conn.Exec(context.Background(), "insert into ride_request (geom) values (ST_Point($1, $2, 3857))", latitude, longitude)
	return err
}
