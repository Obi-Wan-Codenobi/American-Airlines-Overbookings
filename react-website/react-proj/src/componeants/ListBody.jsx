import ListComponeant from "./ListComponeant"

export default function ListBody( {missedFlights} ){
  const listItems = missedFlights.map(flight => 
      //<ListComponeant indFlight={flight} />
      <li>{flight["flightNumber"]}</li>
    )

  return (
    <body className="list-body">
      <h1>Overloaded Flights:</h1>
      <ul>
        {listItems}
      </ul>
    </body>
  )
}