package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"weather-cache/internal/api/handler"
	"weather-cache/internal/constants"
	"weather-cache/internal/model"
	"weather-cache/internal/validator"
)

func (c Controller) GetWeatherByCoords(e echo.Context) error {
	ctx := e.Request().Context()
	c.log = c.log.With(c.log.RequestID(ctx))

	lat := e.QueryParam("lat")
	lon := e.QueryParam("lon")
	c.log.Debug("GetWeather", "lat", lat, "lon", lon)

	latF, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		c.log.Debug("Failed to parse latitude", c.log.Err(err))

		return handler.Return(e, http.StatusBadRequest, model.Weather{}, constants.ErrInvalidLatitude)
	}

	lonF, err := strconv.ParseFloat(lon, 64)
	if err != nil {
		c.log.Debug("Failed to parse longitude", c.log.Err(err))

		return handler.Return(e, http.StatusBadRequest, model.Weather{}, constants.ErrInvalidLongitude)
	}

	if err := validator.ValidateCoordinates(latF, lonF); err != nil {
		c.log.Debug("Invalid coordinates", c.log.Err(err))

		return handler.Return(e, http.StatusBadRequest, model.Weather{}, constants.ErrInvalidCoordinates)
	}

	w, err := c.weatherService.GetWeatherByCoords(ctx, latF, lonF)
	if err != nil {
		c.log.Error("Failed to get weather by coordinates", c.log.Err(err))

		return handler.Return(
			e,
			http.StatusInternalServerError,
			model.Weather{},
			fmt.Errorf("failed to get weather by coordinates: %w", err),
		)
	}

	return handler.Return(e, http.StatusOK, w, nil)
}
