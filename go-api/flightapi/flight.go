
package flightapi


import (
	"time"
    "log"
    "encoding/json"
)

type FlightInfo struct {
	FlightNumber string `json:"flightNumber"`
	Origin       string `json:"origin"`
	Destination  string `json:"destination"`
}

func FlightMain(flightInfoChan chan<- []byte){

    for {

        log.Printf("Grabbing Info From API")

        flight := FlightInfo{
            FlightNumber: "AA123",
            Origin:       "PHL",
            Destination:  "DFW",
        }

        
        jsonData, err := json.Marshal(flight)
        if err != nil {
           
            log.Printf("Error marshaling JSON: %v", err)
            continue
        }

        // Send the JSON data to the channel
        flightInfoChan <- jsonData

        // Simulate a delay between updates
        time.Sleep(1 * time.Second)
    }
}
