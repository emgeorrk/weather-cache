package gismeteo

import "time"

type WeatherDataCurrentWeatherGeoResponse struct {
	Astro         AstroCurrentWeatherGeoResponse         `json:"astro"`
	Icon          IconCurrentWeatherGeoResponse          `json:"icon"`
	Kind          string                                 `json:"kind"`
	Description   string                                 `json:"description"`
	Date          WeatherDateCurrentWeatherGeoResponse   `json:"date"`
	City          CityCurrentWeatherGeoResponse          `json:"city"`
	Wind          WindCurrentWeatherGeoResponse          `json:"wind"`
	Precipitation PrecipitationCurrentWeatherGeoResponse `json:"precipitation"`
	Temperature   TemperatureCurrentWeatherGeoResponse   `json:"temperature"`
	Storm         StormCurrentWeatherGeoResponse         `json:"storm"`
	Cloudiness    CloudinessCurrentWeatherGeoResponse    `json:"cloudiness"`
	Visibility    VisibilityCurrentWeatherGeoResponse    `json:"visibility"`
	Humidity      HumidityCurrentWeatherGeoResponse      `json:"humidity"`
	Pressure      PressureCurrentWeatherGeoResponse      `json:"pressure"`
}

type AstroCurrentWeatherGeoResponse struct {
	Sun  SunCurrentWeatherGeoResponse  `json:"sun"`
	Moon MoonCurrentWeatherGeoResponse `json:"moon"`
}

type SunCurrentWeatherGeoResponse struct {
	Sunrise time.Time   `json:"sunrise"`
	Sunset  time.Time   `json:"sunset"`
	Polar   interface{} `json:"polar"`
}

type MoonCurrentWeatherGeoResponse struct {
	NextFull           time.Time `json:"next_full"`
	PreviousFull       time.Time `json:"previous_full"`
	Phase              string    `json:"phase"`
	PercentIlluminated float64   `json:"percent_illuminated"`
}

type IconCurrentWeatherGeoResponse struct {
	IconWeather string `json:"icon-weather"`
	Emoji       string `json:"emoji"`
}

type WeatherDateCurrentWeatherGeoResponse struct {
	UTC            time.Time `json:"UTC"`
	Local          time.Time `json:"local"`
	Unix           int       `json:"unix"`
	TimeZoneOffset int       `json:"timeZoneOffset"`
}

type CityCurrentWeatherGeoResponse struct {
	Name      string  `json:"name"`
	NameP     string  `json:"nameP"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type WindCurrentWeatherGeoResponse struct {
	Direction          WindDirectionCurrentWeatherGeoResponse `json:"direction"`
	Speed              WindSpeedCurrentWeatherGeoResponse     `json:"speed"`
	GustSpeed          WindSpeedCurrentWeatherGeoResponse     `json:"gust_speed"`
	AlternateDirection bool                                   `json:"alternate_direction"`
}

type WindDirectionCurrentWeatherGeoResponse struct {
	Degree int `json:"degree"`
	Scale8 int `json:"scale_8"`
}

type WindSpeedCurrentWeatherGeoResponse struct {
	MS float64 `json:"m_s"`
}

type PrecipitationCurrentWeatherGeoResponse struct {
	Type      int     `json:"type"`
	TypeExt   int     `json:"type_ext"`
	Amount    float64 `json:"amount"`
	Intensity int     `json:"intensity"`
	Duration  int     `json:"duration"`
}

type TemperatureCurrentWeatherGeoResponse struct {
	Air     TemperatureValueCurrentWeatherGeoResponse `json:"air"`
	Comfort TemperatureValueCurrentWeatherGeoResponse `json:"comfort"`
	Water   TemperatureValueCurrentWeatherGeoResponse `json:"water"`
}

type TemperatureValueCurrentWeatherGeoResponse struct {
	C float64 `json:"C"`
}

type StormCurrentWeatherGeoResponse struct {
	Cape       float64 `json:"cape"`
	Prediction bool    `json:"prediction"`
}

type CloudinessCurrentWeatherGeoResponse struct {
	Percent int `json:"percent"`
	Scale3  int `json:"scale_3"`
}

type VisibilityCurrentWeatherGeoResponse struct {
	Horizontal HorizontalCurrentWeatherGeoResponse `json:"horizontal"`
}

type HorizontalCurrentWeatherGeoResponse struct {
	M int `json:"m"`
}

type HumidityCurrentWeatherGeoResponse struct {
	Percent  int                               `json:"percent"`
	DewPoint DewPointCurrentWeatherGeoResponse `json:"dew_point"`
}

type DewPointCurrentWeatherGeoResponse struct {
	C float64 `json:"C"`
}

type PressureCurrentWeatherGeoResponse struct {
	MmHgAtm int `json:"mm_hg_atm"`
}
