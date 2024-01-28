package flightapi

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Obi-Wan-Codenobi/American-Airlines-Overbookings/go-api/flightJson"
)

func FlightMain(flightInfoChan chan<- []byte) {
    apiURL := "http://localhost:4000/flights?date=2023-01-01"

    for {
        log.Printf("Grabbing Info From API")

        response, err := http.Get(apiURL)
        if err != nil {
            log.Printf("Error making HTTP request: %v", err)
            continue
        }
        defer response.Body.Close()

        if response.StatusCode != http.StatusOK {
            log.Printf("Error: Unexpected status code: %v", response.Status)
            continue
        }

        var flightInfoList []flightJson.FlightInfo
        err = json.NewDecoder(response.Body).Decode(&flightInfoList)
        if err != nil {
            log.Printf("Error decoding JSON: %v", err)
            continue
        }

        for _, flightInfo := range flightInfoList {
            jsonData, err := json.Marshal(flightInfo)
            if err != nil {
                log.Printf("Error marshaling JSON: %v", err)
                continue
            }

            flightInfoChan <- jsonData
        }

        time.Sleep(5 * time.Second)
    }
}

