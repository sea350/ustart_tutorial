package sqlstore

// Config is the configuration for the sql storage client
type Config struct {
	DriverName      string
	Host            string
	Port            string
	DBName          string
	Username        string
	Password        string
	RecordTableName string
}

// NewConfig creates a default config struct
func NewConfig() *Config {
	return &Config{
		DriverName: "postgresql",
		Host:       "localhost",
		Port:       "5432",
		DBName:     "ustart_auth",
		Username:   "postgres",
		Password:   "postgres",
	}
}
