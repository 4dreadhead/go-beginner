package geo_test

import (
	"demo/weather/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	city := "London"
	expected := geo.GeoData{City: city}

	got, err := geo.GetMyLocation(city)
	if err != nil {
		t.Error("Can't get city")
	}
	if got.City != expected.City {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	city := "SomeInvalidCity"
	expected := geo.ErrorNoCity

	_, err := geo.GetMyLocation(city)
	if err != expected {
		t.Errorf("Expected: %v, got: %v", expected, err)
	}
}
