package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

func (c *Controller) GetWeatherByCity(e echo.Context) error {
	city := e.QueryParam("name")
	c.logger.Info("GetWeather", "city", city)

	return e.String(http.StatusOK, "GetWeather")
}
