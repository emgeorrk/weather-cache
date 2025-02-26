package config

type Server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Weather struct {
	APIType string                `mapstructure:"api_type"`
	APIs    map[string]WeatherAPI `mapstructure:"apis"`
}

type WeatherAPI struct {
	URL string `mapstructure:"url"`
	Key string `mapstructure:"key"`
}

type Maps struct {
	APIType string             `mapstructure:"api_type"`
	APIs    map[string]MapsAPI `mapstructure:"apis"`
}

type MapsAPI struct {
	URL string `mapstructure:"url"`
	Key string `mapstructure:"key"`
}

type Cache struct {
	Type string `mapstructure:"type"`
	TTL  string `mapstructure:"ttl"`
}

type Log struct {
	Level        string `mapstructure:"level"`
	Label        string `mapstructure:"label"`
	AddTimestamp string `mapstructure:"add_timestamp"`
	TimeFormat   string `mapstructure:"time_format"`
	Prefix       string `mapstructure:"prefix"`
	AddSource    string `mapstructure:"add_source"`
	Formatter    string `mapstructure:"formatter"`
	SourceFormat string `mapstructure:"source_format"`
	EnableFxLogs string `mapstructure:"enable_fx_logs"`
}

type Config struct {
	Env     string  `mapstructure:"env"`
	Server  Server  `mapstructure:"server"`
	Weather Weather `mapstructure:"weather"`
	Maps    Maps    `mapstructure:"maps"`
	Cache   Cache   `mapstructure:"cache"`
	Log     Log     `mapstructure:"log"`
}
