package storage

import (
	sqlstore "github.com/sea350/ustart_tutorial/record/storage/sql"
)

// Config determines the runtime behavior of the an either SQL or ElasticSearch backed auth server
type Config struct {
	SQLConfig *sqlstore.Config
}

// SQLNewConfig returns a default config object
func SQLNewConfig() *Config {
	return &Config{SQLConfig: sqlstore.NewConfig()}
}
