package sqlstore

import (
	"context"
	"fmt"
	"strconv"
)

//Insert creates a new row for a new registering record
func (dbConn *SQLStore) Insert(ctx context.Context, carID, userID, startDate string, rate int) error {

	queryString := fmt.Sprintf(
		"INSERT INTO %s (carID, userID, startDate, rate) VALUES ( '%s', %s', '%s', '"+strconv.Itoa(rate)+"');",
		dbConn.recordTablaNAme, carID, userID, startDate)

	_, err := dbConn.db.QueryContext(ctx, queryString)

	return err

}
