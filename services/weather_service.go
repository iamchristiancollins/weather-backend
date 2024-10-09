package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherResponse struct {
	//define struct based on API response
}

func GetWeatherData(location string) (*WeatherResponse, error) {
	apiKey := "API KEY GOES HERE"
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?unitGroup=metric&key=%s", location, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var weatherData WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return nil, err
	}
	return &weatherData, nil
}
