package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"


	"github.com/Obi-Wan-Codenobi/American-Airlines-Overbookings/go-api/flightJson"
	"github.com/Obi-Wan-Codenobi/American-Airlines-Overbookings/go-api/flightapi"
)



func main() {

    flightInfoChan := make(chan []byte)

	go flightapi.FlightMain(flightInfoChan)


	go func() {
		http.HandleFunc("/", homeHandler(flightInfoChan))  //GET request

		// Start the HTTP server on port 8080
        fmt.Println("Server is running on http://localhost:8080")

        if err := http.ListenAndServe(":8080", nil); err != nil {
            log.Fatal(err)
        }
	}()

    //for Testing output on server 
//    for {
//        jsonData := <-flightInfoChan
//        var flightInfo flightJson.FlightInfo
//        err := json.Unmarshal(jsonData, &flightInfo)
//        if err != nil {
//            log.Printf("Error unmarshalling JSON: %v", err)
//            continue
//        }
//
//        fmt.Printf("Received flight info: %+v\n", flightInfo)
//
//        fmt.Println("#####################################")
//    }
    
    select {} //block to keep it from terminating

}

func homeHandler(flightInfoChan <-chan []byte) http.HandlerFunc{

    return func(w http.ResponseWriter, r *http.Request){

        // grabs info from go channel that is constantly being fed api flight log data from another thread in background
        jsonData := <-flightInfoChan
        var flightInfo flightJson.FlightInfo
        err := json.Unmarshal(jsonData, &flightInfo)
        if err != nil {
            log.Printf("Error unmarshalling JSON: %v", err)
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
}

