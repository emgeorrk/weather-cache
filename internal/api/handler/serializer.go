package handler

import (
	"github.com/bytedance/sonic"
	"github.com/labstack/echo"
	"net/http"
	"weather-cache/internal/model"
)

func Return(e echo.Context, statusCode int, weather model.Weather, err error) error {
	resp := model.APIResponse{}

	if err != nil {
		resp.APIError = model.APIError{
			Code:    statusCode,
			Error:   http.StatusText(statusCode),
			Message: err.Error(),
		}
	} else {
		resp.Response = weather
	}

	jsonBytes, err := sonic.Marshal(resp)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to marshal response")
	}

	e.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	e.Response().WriteHeader(statusCode)
	_, err = e.Response().Write(jsonBytes)

	return err
}
