package api

import (
	"weather-cache/internal/api/controller"
	"weather-cache/internal/api/handler"
)

func Setup(r handler.RequestHandler, controller controller.Controller) {
	v1 := r.Group("/v1")

	v1.GET("/weather/city", controller.GetWeatherByCity)
	v1.GET("/weather/coords", controller.GetWeatherByCoords)
}
