package flightapi

import (
	"encoding/json"
	"log"
	"net/http"
    "fmt"

	"github.com/Obi-Wan-Codenobi/American-Airlines-Overbookings/go-api/flightJson"
)

func FlightMain(date string) ([]flightJson.FlightInfo, error){

    log.Print("Fetching API data from American Airlines")

    apiURL := fmt.Sprintf("http://localhost:4000/flights?date=%s", date)

    response, err := http.Get(apiURL)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("Unexpected status code: %v", response.Status)
    }

    var flightInfoList []flightJson.FlightInfo
    err = json.NewDecoder(response.Body).Decode(&flightInfoList)
    if err != nil {
        return nil, err
    }

    return flightInfoList, nil
}

