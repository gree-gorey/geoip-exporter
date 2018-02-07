// Package geo implements function for searching
// for a country code by IP address.

package geo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Type GeoIP stores whois info.
type GeoIP struct {
	Ip          string  `json:""`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:""`
	RegionCode  string  `json:"region_code"`
	RegionName  string  `json:"region_name"`
	City        string  `json:"city"`
	Zipcode     string  `json:"zipcode"`
	Lat         float32 `json:"latitude"`
	Lon         float32 `json:"longitude"`
	MetroCode   int     `json:"metro_code"`
	AreaCode    int     `json:"area_code"`
}

var (
	address  string
	err      error
	geo      GeoIP
	response *http.Response
	body     []byte
)

// Function GetCode gets country code by IP address.
func GetCode(address string) (string, error) {
	response, err = http.Get("https://freegeoip.net/json/" + address)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	err = json.Unmarshal(body, &geo)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return geo.CountryCode, nil
}
