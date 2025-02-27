package controller

import (
	"github.com/labstack/echo"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"weather-cache/internal/constants"
	"weather-cache/internal/domain/mocks"
	"weather-cache/internal/model"
	"weather-cache/pkg/logger"
)

func TestControllerGetWeatherByCity(t *testing.T) {
	type args struct {
		city string
	}

	tests := []struct {
		name     string
		args     args
		mock     func(weatherService *mocks.MockWeatherService, args args)
		wantCode int
		wantJSON string
	}{
		{
			name: "OK",
			args: args{
				city: "Moscow",
			},
			mock: func(weatherService *mocks.MockWeatherService, args args) {
				weatherService.EXPECT().
					GetWeatherByCity(gomock.Any(), args.city).
					Return(model.Weather{
						Timestamp: 7230013133,
						Location: model.Location{
							City: "Moscow",
							Lat:  55.751244,
							Lon:  37.618423,
						},
						Temperature: 21.4,
						Wind: model.Wind{
							Speed:     12.21,
							Direction: 2,
						},
						Pressure:    31,
						Humidity:    14,
						Description: "Clear",
						Icon:        "01d",
						Emoji:       "☀️",
					}, nil)
			},
			wantCode: http.StatusOK,
			wantJSON: `{"response":{"timestamp":7230013133,"location":{"city":"Moscow","latitude":55.751244,"longitude":37.618423},"temperature":21.4,"wind":{"speed":12.21,"direction":2},"pressure":31,"humidity":14,"condition":"Clear","icon":"01d","emoji":"☀️"}}`,
		},
		{
			name: "Empty city",
			args: args{
				city: "",
			},
			mock: func(weatherService *mocks.MockWeatherService, args args) {
				// No expectations
			},
			wantCode: http.StatusBadRequest,
			wantJSON: `{"code":400,"error":"Bad Request","message":"city name is required"}`,
		},
		{
			name: "weatherService error",
			args: args{
				city: "Saint Petersburg",
			},
			mock: func(weatherService *mocks.MockWeatherService, args args) {
				weatherService.EXPECT().
					GetWeatherByCity(gomock.Any(), args.city).
					Return(model.Weather{}, constants.ErrRemoteServerOut)
			},
			wantCode: http.StatusInternalServerError,
			wantJSON: `{"code":500,"error":"Internal Server Error","message":"failed to get weather by city: remote server is out"}`,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	lg := logger.NewTestLogger()
	weatherService := mocks.NewMockWeatherService(ctrl)

	c := NewController(lg, weatherService)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(weatherService, tt.args)

			path := "/test"

			e := echo.New()
			e.GET(path, c.GetWeatherByCity)

			u, err := url.Parse(path)
			if err != nil {
				t.Fatal(err)
			}

			query := u.Query()
			query.Set("name", tt.args.city)
			u.RawQuery = query.Encode()

			req := httptest.NewRequest(http.MethodGet, u.String(), nil)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			if rec.Body.String() != tt.wantJSON {
				t.Errorf("unexpected response: got %v, want %v", rec.Body.String(), tt.wantJSON)
			}

			if rec.Code != tt.wantCode {
				t.Errorf("unexpected status code: got %v, want %v", rec.Code, http.StatusOK)
			}
		})
	}
}

func TestControllerGetWeatherByCoords(t *testing.T) {
	type args struct {
		lat string
		lon string
	}

	tests := []struct {
		name     string
		args     args
		mock     func(weatherService *mocks.MockWeatherService, args args)
		wantCode int
		wantJSON string
	}{
		{
			name: "OK",
			args: args{
				lat: "55.751244",
				lon: "37.618423",
			},
			mock: func(weatherService *mocks.MockWeatherService, args args) {
				weatherService.EXPECT().
					GetWeatherByCoords(gomock.Any(), 55.751244, 37.618423).
					Return(model.Weather{
						Timestamp: 7230013133,
						Location: model.Location{
							City: "Moscow",
							Lat:  55.751244,
							Lon:  37.618423,
						},
						Temperature: 21.4,
						Wind: model.Wind{
							Speed:     12.21,
							Direction: 2,
						},
						Pressure:    31,
						Humidity:    14,
						Description: "Clear",
						Icon:        "01d",
						Emoji:       "☀️",
					}, nil)
			},
			wantCode: http.StatusOK,
			wantJSON: `{"response":{"timestamp":7230013133,"location":{"city":"Moscow","latitude":55.751244,"longitude":37.618423},"temperature":21.4,"wind":{"speed":12.21,"direction":2},"pressure":31,"humidity":14,"condition":"Clear","icon":"01d","emoji":"☀️"}}`,
		},
		{
			name: "Invalid latitude",
			args: args{
				lat: "invalid",
				lon: "37.618423",
			},
			mock: func(weatherService *mocks.MockWeatherService, args args) {
				// No expectations
			},
			wantCode: http.StatusBadRequest,
			wantJSON: `{"code":400,"error":"Bad Request","message":"invalid latitude"}`,
		},
		{
			name: "Invalid longitude",
			args: args{
				lat: "55.751244",
				lon: "invalid",
			},
			mock: func(weatherService *mocks.MockWeatherService, args args) {
				// No expectations
			},
			wantCode: http.StatusBadRequest,
			wantJSON: `{"code":400,"error":"Bad Request","message":"invalid longitude"}`,
		},
		{
			name: "weatherService error",
			args: args{
				lat: "55.751244",
				lon: "37.618423",
			},
			mock: func(weatherService *mocks.MockWeatherService, args args) {
				weatherService.EXPECT().
					GetWeatherByCoords(gomock.Any(), 55.751244, 37.618423).
					Return(model.Weather{}, constants.ErrRemoteServerOut)
			},
			wantCode: http.StatusInternalServerError,
			wantJSON: `{"code":500,"error":"Internal Server Error","message":"failed to get weather by coordinates: remote server is out"}`,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	lg := logger.NewTestLogger()
	weatherService := mocks.NewMockWeatherService(ctrl)

	c := NewController(lg, weatherService)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(weatherService, tt.args)

			path := "/test"

			e := echo.New()
			e.GET(path, c.GetWeatherByCoords)

			u, err := url.Parse(path)
			if err != nil {
				t.Fatal(err)
			}

			query := u.Query()
			query.Set("lat", tt.args.lat)
			query.Set("lon", tt.args.lon)
			u.RawQuery = query.Encode()

			req := httptest.NewRequest(http.MethodGet, u.String(), nil)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			if rec.Body.String() != tt.wantJSON {
				t.Errorf("unexpected response: got %v, want %v", rec.Body.String(), tt.wantJSON)
			}

			if rec.Code != tt.wantCode {
				t.Errorf("unexpected status code: got %v, want %v", rec.Code, http.StatusOK)
			}
		})
	}
}
