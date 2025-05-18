package main

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"flag"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	city := flag.String("city", "", "Weather city")
	format := flag.Int("format", 1, "Weather format")
	flag.Parse()

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		color.Red("Can't get geolocation: %v", err)
		return
	}
	fmt.Printf("City is %s\n", geoData.City)

	wthr, err := weather.GetWeather(geoData, *format)
	if err != nil {
		color.Red("Can't get weather: %v", err)
	}
	fmt.Printf("Weather: %v", wthr)
}
