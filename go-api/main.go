package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Obi-Wan-Codenobi/American-Airlines-Overbookings/go-api/flightapi"
)

type FlightInfo struct {
	FlightNumber string `json:"flightNumber"`
	Origin       string `json:"origin"`
	Destination  string `json:"destination"`

  }

func main() {

    flightInfoChan := make(chan []byte)

	go flightapi.FlightMain(flightInfoChan)


	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/user/", userHandler)

	// Start the HTTP server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

    while{
        jsonData := <-flightInfoChan:
        var flightInfo FlightInfo
        err := json.Unmarshal(jsonData, &flightInfo)
        if err != nil {
            log.Printf("Error unmarshaling JSON: %v", err)
            continue
        }

        // Print the decoded data
        fmt.Printf("Received flight info: %+v\n", flightInfo)
    }

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!") 
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/user/")
	fmt.Fprintf(w, "Hello, %s!", name)
}
