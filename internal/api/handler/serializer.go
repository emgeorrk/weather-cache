package handler

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"weather-cache/internal/model"
)

func Return(e echo.Context, statusCode int, data model.APIResponse) error {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to marshal response")
	}

	e.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	e.Response().WriteHeader(statusCode)
	_, err = e.Response().Write(jsonBytes)

	return err
}
