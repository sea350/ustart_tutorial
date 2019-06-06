package storage

import (
	elasticstore "github.com/aaron-min/ustart_tutorial/car/storage/elastic"
)

// NewES determines the runtime behavior of the ElasticSearch-backed customer server
func NewES(config *Config) (Storage, error) {
	return elasticstore.New(config.ElasticConfig)
}
