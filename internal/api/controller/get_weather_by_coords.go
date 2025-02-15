package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"weather-cache/internal/api/handler"
)

func (c *Controller) GetWeatherByCoordinates(e echo.Context) error {
	ctx := e.Request().Context()

	lat := e.QueryParam("lat")
	lon := e.QueryParam("lon")
	c.logger.Info("GetWeather", "lat", lat, "lon", lon)

	latF, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return e.String(http.StatusBadRequest, "Invalid latitude")
	}

	lonF, err := strconv.ParseFloat(lon, 64)
	if err != nil {
		return e.String(http.StatusBadRequest, "Invalid longitude")
	}

	w, err := c.weatherService.GetWeatherByCoords(ctx, latF, lonF)
	if err != nil {
		return err
	}

	return handler.SonicJSON(e, http.StatusOK, w)
}
