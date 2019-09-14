package recordapi

import (
	"github.com/sea350/ustart_tutorial/record"
)

// Config for auth api
type Config struct {
	CustomerCfg *record.Config
	Port        int //Recomended 5101
}
