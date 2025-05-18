package weather

import (
	"demo/weather/geo"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

var ErrorWrongFormat   = errors.New("ERR_WRONG_FORMAT")
var ErrorIncorrectReq  = errors.New("ERR_INCORRECT_REQ")
var ErrorIncorrectResp = errors.New("ERR_INCORRECT_BODY")

func GetWeather(geo *geo.GeoData, format int) (string, error) {
	if format < 1 || format > 4 {
		return "", ErrorWrongFormat
	}
	baseUrl, _ := url.Parse("https://wttr.in/")
	params     := url.Values{}
	params.Add("format", strconv.Itoa(format))

	baseUrl.RawQuery = params.Encode()
	baseUrl.Path     = geo.City

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		return "", ErrorIncorrectReq
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", ErrorIncorrectResp
	}
	return string(body), nil
}
