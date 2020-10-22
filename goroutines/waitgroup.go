package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

const apiKey = ""

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go func(w *sync.WaitGroup) {
		defer wg.Done()
		cityName := "Florianopolis"
		t := GetTemperature(cityName)
		log.Printf("Temperatura da cidade %s = %f graus celsius", cityName, t)
	}(&wg)

	go func(w *sync.WaitGroup) {
		defer wg.Done()
		cityName := "Salvador"
		t := GetTemperature(cityName)
		log.Printf("Temperatura da cidade %s = %f graus celsius", cityName, t)
	}(&wg)

	go func(w *sync.WaitGroup) {
		defer wg.Done()
		cityName := "London"
		t := GetTemperature(cityName)
		log.Printf("Temperatura da cidade %s = %f graus celsius", cityName, t)
	}(&wg)

	wg.Wait()
}

// GetTemperature returns the temperature of the city in Celsius degrees
func GetTemperature(cityName string) float32 {
	url := "http://api.openweathermap.org/data/2.5/weather"
	req, err := http.NewRequest(http.MethodGet, url, nil)

	q := req.URL.Query()
	q.Add("q", cityName)
	q.Add("appid", apiKey)
	q.Add("units", "metric")
	req.URL.RawQuery = q.Encode()

	if err != nil {
		log.Print(err)
	}

	c := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := c.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
	}

	type temperature struct {
		CurrentTemperature   float32 `json:"temp"`
		SensationTemperature float32 `json:"feels_like"`
	}
	type respOpenWeather struct {
		ID   int         `json:"id"`
		Name string      `json:"name"`
		Main temperature `json:"main"`
	}

	respTemp := respOpenWeather{}
	json.Unmarshal(body, &respTemp)

	return respTemp.Main.CurrentTemperature
}
