package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
    "time"
    "math/rand"
    "github.com/Obi-Wan-Codenobi/American-Airlines-Overbookings/go-api/flightJson"
	"github.com/Obi-Wan-Codenobi/American-Airlines-Overbookings/go-api/flightapi"
)


type FlightInfoPassed struct {
	FlightNumber   string `json:"flightNumber"`
	Origin         string `json:"origin"`
	Destination    string `json:"destination"`
	Capacity       int    `json:"capacity"`
	PassengerCount int    `json:"passengerCount"`
	Date           string `json:"date"`
}

func main() {



    http.HandleFunc("/", homeHandler)  //GET request

    // Start the HTTP server on port 8080
    fmt.Println("Server is running on http://localhost:8080")

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {


    // Parse the query parameters from the request
    queryValues := r.URL.Query()
    date := queryValues.Get("date") 
    flightInfoList, err := flightapi.FlightMain(date)
    if err != nil {
        log.Printf("Error getting flight info: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    log.Println("Received flight info and returning json")

    var flightInfoPassedList []FlightInfoPassed

    for _, flightInfo := range flightInfoList {
        flightInfoPassed := convertToPassedFormat(flightInfo, date)
        flightInfoPassedList = append(flightInfoPassedList, flightInfoPassed)
    }


    rand.Seed(time.Now().UnixNano())
    totalLen := len(flightInfoList)
    var finalFlightInfo []FlightInfoPassed

    for i := 0; i < totalLen; i += 2 {
        flightInfoPassed := flightInfoPassedList[i]
        randomOffset := rand.Intn(50) + 1
        flightInfoPassed.PassengerCount = flightInfoPassed.Capacity + randomOffset
        finalFlightInfo = append(finalFlightInfo, flightInfoPassed)
    }   


    finalFlightJSON, err := json.Marshal(finalFlightInfo)
    if err != nil {
        log.Printf("Error marshalling JSON: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Returning JSON
    w.Header().Set("Content-Type", "application/json")
    w.Write(finalFlightJSON)
}


func convertToPassedFormat(flightInfo flightJson.FlightInfo, date string) FlightInfoPassed {
	return FlightInfoPassed{
		FlightNumber:   flightInfo.FlightNumber,
		Origin:         flightInfo.Origin.Code,	
        Destination:    flightInfo.Destination.Code,
		Capacity:       flightInfo.Aircraft.PassengerCapacity.Total,
		PassengerCount: flightInfo.Aircraft.PassengerCapacity.Total, 
        Date:           date,
	}
}
