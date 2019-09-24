package geoplugin

import (
	"testing"
)

func TestGetGeoIp(t *testing.T) {
	t.Run("Should return an IP address from Geoplugin", func(t *testing.T) {
		geoip := GetGeoIP()

		if geoip.Request == "" {
			t.Errorf("Geoplugin should be a valid IP address but got %s", geoip.Request)
		}
	})
}
