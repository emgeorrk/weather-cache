package gismeteo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"weather-cache/internal/constants"
	"weather-cache/internal/model"
)

func (g API) FetchWeather(ctx context.Context, city model.Location) (model.Weather, error) {
	endpoint := fmt.Sprintf("%s/v3/weather/current/?latitude=%f&longitude=%f", g.URL, city.Lat, city.Lon)
	g.log.Debug("sending request", "endpoint", endpoint)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		g.log.Error("failed to create request", "error", err)
		return model.Weather{}, errors.Join(constants.ErrRemoteServerOut, err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Gismeteo-Token", g.APIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		g.log.Error("failed to send request", "error", err)
		return model.Weather{}, errors.Join(constants.ErrRemoteServerOut, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			g.log.Error("failed to close response body", "error", err)
		}
	}(resp.Body)

	res := WeatherDataCurrentWeatherGeoResponse{}

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		g.log.Error("failed to decode response", "error", err)
		return model.Weather{}, errors.Join(constants.ErrConverting, err)
	}

	r := model.Weather{
		Timestamp: res.Date.UTC.Unix(),
		Location: model.Location{
			City: res.City.Name,
			Lat:  res.City.Latitude,
			Lon:  res.City.Longitude,
		},
		Temperature: res.Temperature.Air.C,
		Wind: model.Wind{
			Speed:     res.Wind.Speed.MS,
			Direction: res.Wind.Direction.Degree,
		},
		Pressure:    res.Pressure.MmHgAtm,
		Humidity:    res.Humidity.Percent,
		Description: res.Description,
		Icon:        res.Icon.IconWeather,
		Emoji:       res.Icon.Emoji,
	}

	g.log.Debug("accepted response", "response", r)

	return r, nil
}
