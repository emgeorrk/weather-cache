package model

type Weather struct {
	Timestamp   int64    `json:"timestamp"`   // Время обновления данных (Unix)
	Location    Location `json:"location"`    // Город и координаты
	Temperature float64  `json:"temperature"` // Температура воздуха (°C)
	Wind        Wind     `json:"wind"`        // Ветер
	Pressure    int      `json:"pressure"`    // Давление (мм рт. ст.)
	Humidity    int      `json:"humidity"`    // Влажность (%)
	Description string   `json:"condition"`   // Описание погоды (ясно, облачно и т.д.)
	Icon        string   `json:"icon"`        // Иконка погоды
	Emoji       string   `json:"emoji"`       // Эмодзи погоды
}

type Location struct {
	City string  `json:"city"`      // Название города
	Lat  float64 `json:"latitude"`  // Широта
	Lon  float64 `json:"longitude"` // Долгота
}

type Wind struct {
	Speed     float64 `json:"speed"`     // Скорость ветра (м/с)
	Direction int     `json:"direction"` // Направление ветра (градусы)
}
