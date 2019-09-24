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
	if err != nil {
		fmt.Fprintf(os.Stderr, "getip: reading %s: %v\n", url, err)
	}

	resp.Body.Close()

	geoip := Geoip{}
	err = json.Unmarshal([]byte(b), &geoip)
	if err != nil {
		fmt.Println(err)
	}

	return geoip
}

// Print a Geoip struct
func (ip Geoip) Print() {
	fmt.Printf("%-22v: %v\n", "Request", ip.Request)
	fmt.Printf("%-22v: %v\n", "Status", ip.Status)
	fmt.Printf("%-22v: %v\n", "Delay", ip.Delay)
	fmt.Printf("%-22v: %v\n", "Dredit", ip.Dredit)
	fmt.Printf("%-22v: %v\n", "City", ip.City)
	fmt.Printf("%-22v: %v\n", "Region", ip.Region)
	fmt.Printf("%-22v: %v\n", "RegionCode", ip.RegionCode)
	fmt.Printf("%-22v: %v\n", "RegionName", ip.RegionName)
	fmt.Printf("%-22v: %v\n", "AreaCode", ip.AreaCode)
	fmt.Printf("%-22v: %v\n", "DmaCode", ip.DmaCode)
	fmt.Printf("%-22v: %v\n", "CountryCode", ip.CountryCode)
	fmt.Printf("%-22v: %v\n", "CountryName", ip.CountryName)
	fmt.Printf("%-22v: %v\n", "InEU", ip.InEU)
	fmt.Printf("%-22v: %v\n", "EuVATrate", ip.EuVATrate)
	fmt.Printf("%-22v: %v\n", "ContinentCode", ip.ContinentCode)
	fmt.Printf("%-22v: %v\n", "ContinentName", ip.ContinentName)
	fmt.Printf("%-22v: %v\n", "Latitude", ip.Latitude)
	fmt.Printf("%-22v: %v\n", "Longitude", ip.Longitude)
	fmt.Printf("%-22v: %v\n", "LocationAccuracyRadius", ip.LocationAccuracyRadius)
	fmt.Printf("%-22v: %v\n", "Timezone", ip.Timezone)
	fmt.Printf("%-22v: %v\n", "CurrencyCode", ip.CurrencyCode)
	fmt.Printf("%-22v: %v\n", "CurrencySymbol", ip.CurrencySymbol)
	fmt.Printf("%-22v: %v\n", "CurrencySymbolUTF8", ip.CurrencySymbolUTF8)
	fmt.Printf("%-22v: %v\n", "CurrencyConverter", ip.CurrencyConverter)
}
