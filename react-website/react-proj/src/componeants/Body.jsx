import ListBody from "./ListBody"
import OutlierBody from "./OutlierBody"

export default function Body( { data }){
  const overloadedFlights = data.filter(flight => 
    flight["passengerCount"] > flight["capacity"])
  return (
    <div className="body">
      <ListBody missedFlights={overloadedFlights}/>
      <OutlierBody flights={overloadedFlights}/>
    </div>
  )
}