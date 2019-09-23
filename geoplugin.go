package geoplugin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Geoip data coming from http://www.geoplugin.net/json.gp
type Geoip struct {
	Request                string `json:"geoplugin_request"`
	Status                 int    `json:"geoplugin_status"`
	Delay                  string `json:"geoplugin_delay"`
	Dredit                 string `json:"geoplugin_credit"`
	City                   string `json:"geoplugin_city"`
	Region                 string `json:"geoplugin_region"`
	RegionCode             string `json:"geoplugin_regionCode"`
	RegionName             string `json:"geoplugin_regionName"`
	AreaCode               string `json:"geoplugin_areaCode"`
	DmaCode                string `json:"geoplugin_dmaCode"`
	CountryCode            string `json:"geoplugin_countryCode"`
	CountryName            string `json:"geoplugin_countryName"`
	InEU                   int    `json:"geoplugin_inEU"`
	EuVATrate              bool   `json:"geoplugin_euVATrate"`
	ContinentCode          string `json:"geoplugin_continentCode"`
	ContinentName          string `json:"geoplugin_continentName"`
	Latitude               string `json:"geoplugin_latitude"`
	Longitude              string `json:"geoplugin_longitude"`
	LocationAccuracyRadius string `json:"geoplugin_locationAccuracyRadius"`
	Timezone               string `json:"geoplugin_timezone"`
	CurrencyCode           string `json:"geoplugin_currencyCode"`
	CurrencySymbol         string `json:"geoplugin_currencySymbol"`
	CurrencySymbolUTF8     string `json:"geoplugin_currencySymbol_UTF8"`
	CurrencyConverter      int    `json:"geoplugin_currencyConverter"`
}

// GetGeoIP function gets IP struct from Geoip
func GetGeoIP() Geoip {
	url := "http://www.geoplugin.net/json.gp"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "getip: %v\n", err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "getip: reading %s: %v\n", url, err)
	}

	geoip := Geoip{}
	bytes := []byte(b)
	err = json.Unmarshal(bytes, &geoip)

	if err != nil {
		fmt.Println(err)
	}

	return geoip
}

// Print a geoip object
func (ip Geoip) Print() {
	fmt.Println(ip.Request)
	fmt.Println(ip.City)
	fmt.Println(ip.RegionCode)
	fmt.Println(ip.CountryName)
	fmt.Printf("%s, %s\n", ip.Latitude, ip.Longitude)
}
