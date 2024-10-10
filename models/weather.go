package models

// WeatherResponse represents the structure of the weather data received from the API.
type WeatherResponse struct {
	QueryCost         int           `json:"queryCost"`
	Latitude          float64       `json:"latitude"`
	Longitude         float64       `json:"longitude"`
	ResolvedAddress   string        `json:"resolvedAddress"`
	Address           string        `json:"address"`
	Timezone          string        `json:"timezone"`
	Tzoffset          float64       `json:"tzoffset"`
	Days              []WeatherDay  `json:"days"`
	CurrentConditions WeatherDetail `json:"currentConditions"`
	// Additional fields can be added here.
}

// WeatherDay represents daily weather details.
type WeatherDay struct {
	Datetime  string  `json:"datetime"`
	Temp      float64 `json:"temp"`
	Feelslike float64 `json:"feelslike"`
	// Include other fields as needed.
}

// WeatherDetail represents detailed current weather information.
type WeatherDetail struct {
	Datetime  string  `json:"datetime"`
	Temp      float64 `json:"temp"`
	Feelslike float64 `json:"feelslike"`
	// Include other fields as needed.
}
