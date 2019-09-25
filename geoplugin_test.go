package geoplugin

import (
	"testing"
)

func TestGetGeoIp(t *testing.T) {
	t.Run("Should return an IP address from Geoplugin", func(t *testing.T) {
		geoip, err := GetGeoIP()

		if err != nil {
			t.Error("GetGeoIP should not return an error")
		}

		geoip.Print()

		if geoip.Request == "" {
			t.Errorf("Geoplugin should be a valid IP address but got %s", geoip.Request)
		}
	})
}
