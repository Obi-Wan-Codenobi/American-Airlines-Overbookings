export default function OutlierBody({ flights }) {
  let overbookedFlightsObject = {}
  
  for (const flight in flights){
    if (flight["origin"] in overbookedFlightsObject){
      overbookedFlightsObject["origin"] = overbookedFlightsObject["origin"] + (flight["passengerCount"] - flight["capacity"]) 
    } else{
      overbookedFlightsObject[flight["origin"]] = flight[(flight["passengerCount"] - flight["capacity"])]
    }
  }
  
  let cityItems = []
  for (const [key,value] of Object.entries(overbookedFlightsObject)){
    cityItems.push(<li>The City is {`${key}`} and has an overbooked capacity of {`${value}`} </li>)
  }
  return(
    <div className="outlier-body">
      <h1>City Data</h1>
      <ul>{cityItems}</ul>
    </div>
    
  )
}