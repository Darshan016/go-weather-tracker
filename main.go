package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type apiConfig struct {
	ApiKey string `json:"ApiKey"`
}

type weatherData struct {
	Location struct {
		Name      string `json:"name"`
		State     string `json:"region"`
		Country   string `json:"country"`
		LocalTime string `json:"localtime"`
	} `json:"location"`
	Current struct {
		TempC      float64 `json:"temp_c"`
		WindInKmph float64 `json:"wind_kph"`
	} `json:"current"`
}

func loadApiConfig(filename string) (apiConfig, error) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return apiConfig{}, err
	}

	var c apiConfig

	err = json.Unmarshal(bytes, &c)

	if err != nil {
		return apiConfig{}, err
	}

	return c, nil
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!\n"))
}

func query(cityname string) (weatherData, error) {
	apiConfig, err := loadApiConfig(".apiConfig")
	if err != nil {
		return weatherData{}, err
	}
	resp, err := http.Get("http://api.weatherapi.com/v1/current.json?key=" + apiConfig.ApiKey + "&q=" + cityname)
	fmt.Printf("queried for city: %v\n", cityname)

	if err != nil {
		return weatherData{}, err

	}
	defer resp.Body.Close()

	var d weatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}
	return d, nil

}

func main() {
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/weather/",
		func(w http.ResponseWriter, r *http.Request) {
			city := strings.SplitN(r.URL.Path, "/", 3)[2]
			data, err := query(city)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(w).Encode(data)
		})
	fmt.Println("Server running on port: 8000")
	http.ListenAndServe(":8000", nil)

}
