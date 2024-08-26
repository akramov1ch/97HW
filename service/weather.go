package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"97HW/config"
	"97HW/models"
)

func UpdateWeatherData() error {
	resp, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=London", config.WeatherAPIKey))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var weatherData models.WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return err
	}

	return saveWeatherDataToDB(weatherData)
}

func saveWeatherDataToDB(data models.WeatherData) error {
	db := getDB()
	_, err := db.Exec("INSERT INTO weather_data (location, temperature) VALUES ($1, $2)",
		data.Location.Name, data.Current.TempC)
	return err
}
