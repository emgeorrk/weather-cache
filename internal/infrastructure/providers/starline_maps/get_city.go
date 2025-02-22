package starline_maps

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"weather-cache/internal/constants"
)

func (a API) GetCityByCoords(ctx context.Context, lat, lon float64) (city string, err error) {
	endpoint := fmt.Sprintf("%s/api/geocoder/v2/reverse?lat=%f&lon=%f", a.URL, lat, lon)
	a.log.Debug("sending request", "endpoint", endpoint)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		a.log.Error("failed to create request", a.log.Err(err))
		return "", errors.Join(constants.ErrRemoteServerOut, err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		a.log.Error("failed to send request", a.log.Err(err))
		return "", errors.Join(constants.ErrRemoteServerOut, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			a.log.Error("failed to close response body", a.log.Err(err))
		}
	}(resp.Body)

	var res []GeocoderResponse

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		a.log.Error("failed to decode response", a.log.Err(err))
		return "", errors.Join(constants.ErrConverting, err)
	}

	a.log.Debug("accepted response", "response", res)

	sort.SliceStable(res, func(i, j int) bool {
		return res[i].Distance < res[j].Distance
	})

	if len(res) == 0 {
		a.log.Error("no results found", "lat", lat, "lon", lon)
		return "", constants.ErrNotFound
	}

	city = res[0].Address.City
	a.log.Debug("matched city", "city", city, "lat", lat, "lon", lon)

	return city, nil
}
