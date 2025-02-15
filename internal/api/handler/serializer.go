package handler

import (
	"github.com/bytedance/sonic"
	"github.com/labstack/echo"
	"net/http"
)

func SonicJSON(e echo.Context, statusCode int, data interface{}) error {
	jsonBytes, err := sonic.Marshal(data)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to marshal response")
	}

	e.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	e.Response().WriteHeader(statusCode)
	_, err = e.Response().Write(jsonBytes)

	return err
}
