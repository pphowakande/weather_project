package configuration

// Configuration used for the main app
type Configuration struct {
	DBConfig        DBConfig
	APIKey          string
	WeatherEndpoint string
}

// DBConfig is the config for database connections
type DBConfig struct {
	DBUser              string
	DBPassword          string
	DBHost              string
	DBName              string
	DBPort              string
	Namespace           string
	CitySet             string
	CodeSet             string
	HistoricalSet       string
	HistoricalHourlySet string
}
