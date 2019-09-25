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
func GetGeoIP() (Geoip, error) {
	geoip := Geoip{}
	url := "http://www.geoplugin.net/json.gp"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "geoplugin: reading %s: %v\n", url, err)
		return geoip, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "geoplugin: %v\n", err)
		return geoip, err
	}

	err = json.Unmarshal([]byte(b), &geoip)
	if err != nil {
		fmt.Fprintf(os.Stderr, "geoplugin: %v\n", err)
		return geoip, err
	}

	return geoip, nil
}

// Print a Geoip struct
func (ip Geoip) Print() {
	fmt.Printf("%-22s: %v\n", "Request", ip.Request)
	fmt.Printf("%-22s: %v\n", "Status", ip.Status)
	fmt.Printf("%-22s: %v\n", "Delay", ip.Delay)
	fmt.Printf("%-22s: %v\n", "Dredit", ip.Dredit)
	fmt.Printf("%-22s: %v\n", "City", ip.City)
	fmt.Printf("%-22s: %v\n", "Region", ip.Region)
	fmt.Printf("%-22s: %v\n", "RegionCode", ip.RegionCode)
	fmt.Printf("%-22s: %v\n", "RegionName", ip.RegionName)
	fmt.Printf("%-22s: %v\n", "AreaCode", ip.AreaCode)
	fmt.Printf("%-22s: %v\n", "DmaCode", ip.DmaCode)
	fmt.Printf("%-22s: %v\n", "CountryCode", ip.CountryCode)
	fmt.Printf("%-22s: %v\n", "CountryName", ip.CountryName)
	fmt.Printf("%-22s: %v\n", "InEU", ip.InEU)
	fmt.Printf("%-22s: %v\n", "EuVATrate", ip.EuVATrate)
	fmt.Printf("%-22s: %v\n", "ContinentCode", ip.ContinentCode)
	fmt.Printf("%-22s: %v\n", "ContinentName", ip.ContinentName)
	fmt.Printf("%-22s: %v\n", "Latitude", ip.Latitude)
	fmt.Printf("%-22s: %v\n", "Longitude", ip.Longitude)
	fmt.Printf("%-22s: %v\n", "LocationAccuracyRadius", ip.LocationAccuracyRadius)
	fmt.Printf("%-22s: %v\n", "Timezone", ip.Timezone)
	fmt.Printf("%-22s: %v\n", "CurrencyCode", ip.CurrencyCode)
	fmt.Printf("%-22s: %v\n", "CurrencySymbol", ip.CurrencySymbol)
	fmt.Printf("%-22s: %v\n", "CurrencySymbolUTF8", ip.CurrencySymbolUTF8)
	fmt.Printf("%-22s: %v\n", "CurrencyConverter", ip.CurrencyConverter)
}
