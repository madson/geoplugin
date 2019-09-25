# geoplugin
A Go module wrapper for Geoplugin API http://www.geoplugin.net/json.gp

### Install

```go
import (
	"github.com/madson/geoplugin"
)
```

Then run `$ go get`.


### Usage example

```go
	ip, err := geoplugin.GetGeoIP()

	if err != nil {
		fmt.Println(err)
	}

	ip.Print()

```


### Geoip struct

```go

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

```

### Print output

```bash
Request               : 98.180.174.3
Status                : 200
Delay                 : 2ms
Dredit                : Some of the returned data includes GeoLite data created by MaxMind, available from <a href='http://www.maxmind.com'>http://www.maxmind.com</a>.
City                  : Scottsdale
Region                : Arizona
RegionCode            : AZ
RegionName            : Arizona
AreaCode              : 
DmaCode               : 753
CountryCode           : US
CountryName           : United States
InEU                  : 0
EuVATrate             : false
ContinentCode         : NA
ContinentName         : North America
Latitude              : 33.6013
Longitude             : -111.8867
LocationAccuracyRadius: 200
Timezone              : America/Phoenix
CurrencyCode          : USD
CurrencySymbol        : $
CurrencySymbolUTF8    : $
CurrencyConverter     : 1
```

### Disclaimer

Use of this wrapper module implies agreement to GeoPlugin Privacy Policy and User Agreement. Check it out at https://www.geoplugin.com/privacy.
