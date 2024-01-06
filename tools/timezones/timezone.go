package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const tzAPIServer = "https://api.wheretheiss.at/v1/coordinates"

func GetTimeZoneByLatLon(lat, lon float64) string {
    fullURL := fmt.Sprintf("%s/%f,%f", tzAPIServer, lat, lon)

    response, err := http.Get(fullURL)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    body, err := io.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    var jsonObject map[string]string
    json.Unmarshal([]byte(body), &jsonObject)
    return jsonObject["timezone_id"]
}
