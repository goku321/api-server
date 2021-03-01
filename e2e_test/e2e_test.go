package e2e_test

import (
	"encoding/json"
	"net/http"
	"testing"

	geo "github.com/goku321/geolocation/geolocation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApiServer(t *testing.T) {
	url := "http://localhost:8080/geolocation/"

	t.Run("ip not found", func(t *testing.T) {
		r, err := http.Get(url + "ip-not-found")
		require.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, r.StatusCode)
	})

	t.Run("ip found in db", func(t *testing.T) {
		want := &geo.GeoData{
			IP:           "200.106.141.15",
			CountryCode:  "SI",
			Country:      "Nepal",
			City:         "DuBuquemouth",
			Latitude:     -84.87503094689836,
			Longitude:    7.206435933364332,
			MysteryValue: 7823011346,
		}
		r, err := http.Get(url + "200.106.141.15")
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, r.StatusCode)

		var actual *geo.GeoData
		d := json.NewDecoder(r.Body)
		err = d.Decode(&actual)
		require.NoError(t, err)
		assert.Equal(t, want, actual)
	})
}
