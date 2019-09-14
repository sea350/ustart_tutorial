package sqlstore

import "context"

// Init initiallizes all required tables
func (dbConn *SQLStore) Init(ctx context.Context) error {
	// the table mirrors the authpb
	_, err := dbConn.db.QueryContext(ctx, `CREATE TABLE IF NOT EXISTS  `+dbConn.recordTablaNAme+` (
	carID text NOT NULL,
	userID text NOT NULL,
	dateStart text NOT NULL,
	dateReturned text,
	rate float NOT NULL,
...
...
..

	PRIMARY KEY  (uuid)
	);
`)

	return err
}

//CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
