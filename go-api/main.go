package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"


	"github.com/Obi-Wan-Codenobi/American-Airlines-Overbookings/go-api/flightapi"
)



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

    flightInfo, err := flightapi.FlightMain(date)
    if err != nil {
        log.Printf("Error getting flight info: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    log.Println("Received flight info and returning json")

    // should convert info to json    
    flightJSON, err := json.Marshal(flightInfo)
    if err != nil {
        log.Printf("Error marshalling JSON: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    //returning json
    w.Header().Set("Content-Type", "application/json")
    w.Write(flightJSON)

}

