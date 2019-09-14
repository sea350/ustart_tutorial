package recordapi

import (
	"strconv"

	"github.com/sea350/ustart_tutorial/record"
)

//RESTAPI is the REST API implementation for record
type RESTAPI struct {
	record *record.Record
	port   string
}

// New creates a new record api, given the config
func New(cfg *Config) (*RESTAPI, error) {
	rec, err := record.New(cfg.CustomerCfg)
	if err != nil {
		return nil, err
	}

	return &RESTAPI{
		record: rec,
		port:   strconv.Itoa(cfg.Port),
	}, nil
}
