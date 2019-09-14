package recordapi

import (
	"strconv"

	"github.com/sea350/ustart_tutorial/record"
)

//GPRCAPI is the GRPC API implementation for record
type GPRCAPI struct {
	record *record.Record
	port   string
}

// New creates a new record api, given the config
func New(cfg *Config) (*GPRCAPI, error) {
	rec, err := record.New(cfg.CustomerCfg)
	if err != nil {
		return nil, err
	}

	return &GPRCAPI{
		record: rec,
		port:   strconv.Itoa(cfg.Port),
	}, nil
}
