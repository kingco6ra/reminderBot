package times

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const tzAPIServer = "https://api.wheretheiss.at/v1/coordinates"

func GetTimeZoneByLatLon(lat, lon float64) string {
	fullURL := fmt.Sprintf("%s/%f,%f", tzAPIServer, lat, lon)

	response, _ := http.Get(fullURL)

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	var jsonObject map[string]string

	json.Unmarshal([]byte(body), &jsonObject)

	return jsonObject["timezone_id"]
}
