package flightJson

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Airport struct {
	Code     string   `json:"code"`
	City     string   `json:"city"`
	Timezone string   `json:"timezone"`
	Location Location `json:"location"`
}

type PassengerCapacity struct {
	Total int `json:"total"`
	Main  int `json:"main"`
	First int `json:"first"`
}

type Duration struct {
	Locale  string `json:"locale"`
	Hours   int    `json:"hours"`
	Minutes int    `json:"minutes"`
}

type Aircraft struct {
	Model             string            `json:"model"`
	PassengerCapacity PassengerCapacity `json:"passengerCapacity"`
	Speed             int               `json:"speed"`
}

type FlightInfo struct {
	FlightNumber string    `json:"flightNumber"`
	Origin       Airport   `json:"origin"`
	Destination  Airport   `json:"destination"`
	Distance     int       `json:"distance"`
	Duration     Duration  `json:"duration"`
	DepartureTime string    `json:"departureTime"`
	ArrivalTime   string    `json:"arrivalTime"`
	Aircraft      Aircraft  `json:"aircraft"`
}

