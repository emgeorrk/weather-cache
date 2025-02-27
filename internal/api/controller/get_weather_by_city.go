package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"weather-cache/internal/api/handler"
	"weather-cache/internal/constants"
	"weather-cache/internal/model"
)

func (c Controller) GetWeatherByCity(e echo.Context) error {
	ctx := e.Request().Context()
	c.log = c.log.With(c.log.RequestID(ctx))

	city := e.QueryParam("name")
	c.log.Debug("GetWeatherByCity", "city", city)

	if city == "" {
		c.log.Debug("City name is required")

		return handler.Return(e, http.StatusBadRequest, model.Weather{}, constants.ErrCityNameRequired)
	}

	w, err := c.weatherService.GetWeatherByCity(ctx, city)
	if err != nil {
		c.log.Error("Failed to get weather by city", c.log.Err(err))

		return handler.Return(
			e,
			http.StatusInternalServerError,
			model.Weather{},
			fmt.Errorf("failed to get weather by city: %w", err),
		)
	}

	return handler.Return(e, http.StatusOK, w, nil)
}
