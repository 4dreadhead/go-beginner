package weather_test

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {
	expected := "Moscow"
	format   := 3
	geoData  := geo.GeoData{
		City: expected,
	}
	result, err := weather.GetWeather(&geoData, format)

	if err != nil {
		t.Errorf("unexpected responce: %v", err)
	}
    if !strings.Contains(result, expected) {
		t.Errorf("Expected to contains: %v, got: %v", expected, result)
	}
}

var testCases = []struct {
	name  string
	value int
}{
	{name: "too big format",  value: 150},
	{name: "zero format",     value: 0},
	{name: "negative format", value: -3},
}

func TestWeatherWrongFormat(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			expected := weather.ErrorWrongFormat
			geoData  := geo.GeoData{
				City: "Moscow",
			}
			_, err := weather.GetWeather(&geoData, testCase.value)
			if err != expected {
				t.Errorf("%v: expected to error: %v, got: %v", testCase.name, expected, err)
			}
		})
	}
}
