package constants

import "errors"

var (
	ErrRemoteServerOut            = errors.New("remote server is out")
	ErrConverting                 = errors.New("converting error")
	ErrNotFound                   = errors.New("not found")
	ErrCacheTypeNotSupported      = errors.New("cache type not supported")
	ErrWeatherAPITypeNotSupported = errors.New("weather api type not supported")
	ErrMapsAPITypeNotSupported    = errors.New("maps api type not supported")

	ErrInvalidLatitude    = errors.New("invalid latitude")
	ErrInvalidLongitude   = errors.New("invalid longitude")
	ErrInvalidCoordinates = errors.New("invalid coordinates")
	ErrCityNameRequired   = errors.New("city name is required")
)
