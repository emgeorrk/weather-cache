package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"weather-cache/internal/api/handler"
	"weather-cache/internal/model"
)

func (c Controller) GetWeatherByCity(e echo.Context) error {
	ctx := e.Request().Context()
	c.log.Info("GetWeatherByCity", c.log.RequestID(ctx))

	city := e.QueryParam("name")
	if city == "" {
		c.log.Debug("City name is required", c.log.RequestID(ctx))

		return handler.Return(e, http.StatusBadRequest, model.APIResponse{
			APIError: model.APIError{
				Code:    http.StatusBadRequest,
				Error:   http.StatusText(http.StatusBadRequest),
				Message: "City name is required",
			},
		})
	}

	c.log.Debug("GetWeatherByCity", "city", city, c.log.RequestID(ctx))

	w, err := c.weatherService.GetWeatherByCity(ctx, city)
	if err != nil {
		c.log.Error("Failed to get weather by city", c.log.Err(err), c.log.RequestID(ctx))

		return handler.Return(e, http.StatusInternalServerError, model.APIResponse{
			APIError: model.APIError{
				Code:    http.StatusInternalServerError,
				Error:   http.StatusText(http.StatusInternalServerError),
				Message: "Failed to get weather by city",
			},
		})
	}

	return handler.Return(e, http.StatusOK, model.APIResponse{
		Response: w,
	})
}
