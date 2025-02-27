package validator

import (
	"errors"
	"weather-cache/internal/constants"
)

// ValidateCoordinates checks if the given coordinates are valid.
func ValidateCoordinates(lat, lon float64) error {
	if lat < -90 || lat > 90 {
		return errors.Join(constants.ErrInvalidCoordinates, errors.New("latitude must be between -90 and 90"))
	}
	if lon < -180 || lon > 180 {
		return errors.Join(constants.ErrInvalidCoordinates, errors.New("longitude must be between -180 and 180"))
	}
	return nil
}
