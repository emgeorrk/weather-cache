package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"weather-cache/internal/api/handler"
	"weather-cache/internal/model"
)

func (c Controller) GetWeatherByCoords(e echo.Context) error {
	ctx := e.Request().Context()

	lat := e.QueryParam("lat")
	lon := e.QueryParam("lon")
	c.log.Debug("GetWeather", "lat", lat, "lon", lon, c.log.RequestID(ctx))

	latF, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		c.log.Debug("Failed to parse latitude", c.log.Err(err), c.log.RequestID(ctx))

		return handler.Return(e, http.StatusBadRequest, model.APIResponse{
			APIError: model.APIError{
				Code:    http.StatusBadRequest,
				Error:   http.StatusText(http.StatusBadRequest),
				Message: "Invalid latitude",
			},
		})
	}

	lonF, err := strconv.ParseFloat(lon, 64)
	if err != nil {
		c.log.Debug("Failed to parse longitude", c.log.Err(err), c.log.RequestID(ctx))

		return handler.Return(e, http.StatusBadRequest, model.APIResponse{
			APIError: model.APIError{
				Code:    http.StatusBadRequest,
				Error:   http.StatusText(http.StatusBadRequest),
				Message: "Invalid longitude",
			},
		})
	}

	w, err := c.weatherService.GetWeatherByCoords(ctx, latF, lonF)
	if err != nil {
		c.log.Error("Failed to get weather by coordinates", c.log.Err(err), c.log.RequestID(ctx))

		return handler.Return(e, http.StatusInternalServerError, model.APIResponse{
			APIError: model.APIError{
				Code:    http.StatusInternalServerError,
				Error:   http.StatusText(http.StatusInternalServerError),
				Message: "Failed to get weather by coordinates",
			},
		})
	}

	return handler.Return(e, http.StatusOK, model.APIResponse{
		Response: w,
	})
}
