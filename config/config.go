package config

import (
	"go.uber.org/fx"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Module = fx.Options(
	fx.Provide(Load),
)

type Config struct {
	WeatherAPIType  string
	WeatherAPIURL   string
	WeatherAPIKey   string
	MapsAPIType     string
	MapsAPIURL      string
	MapsAPIKey      string
	CacheType       string
	CacheTTL        string
	LogLevel        string
	LogLabel        string
	LogAddTimeStamp string
	LogTimeFormat   string
	LogPrefix       string
	LogAddSource    string
	LogFormatter    string
	LogSourceFormat string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return Config{
		WeatherAPIType:  getEnv("WEATHER_API_TYPE", "gismeteo"),
		WeatherAPIURL:   getEnv("WEATHER_API_URL", "api.gismeteo.net"),
		WeatherAPIKey:   getEnv("WEATHER_API_KEY", ""),
		MapsAPIType:     getEnv("MAPS_API_TYPE", "starline_maps"),
		MapsAPIURL:      getEnv("MAPS_API_URL", "maps.starline.ru"),
		MapsAPIKey:      getEnv("MAPS_API_KEY", ""),
		CacheType:       getEnv("CACHE_TYPE", "in-memory"),
		CacheTTL:        getEnv("CACHE_TTL", "5m"),
		LogLevel:        getEnv("LOG_LEVEL", "info"),
		LogLabel:        getEnv("LOG_LABEL", "weather-cache"),
		LogAddTimeStamp: getEnv("LOG_ADD_TIMESTAMP", "yes"),
		LogTimeFormat:   getEnv("LOG_TIME_FORMAT", "2006-01-02 15:04:05 MST"),
		LogPrefix:       getEnv("LOG_PREFIX", ""),
		LogAddSource:    getEnv("LOG_ADD_SOURCE", "yes"),
		LogFormatter:    getEnv("LOG_FORMATTER", "text"),
		LogSourceFormat: getEnv("LOG_SOURCE_FORMAT", "short"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
