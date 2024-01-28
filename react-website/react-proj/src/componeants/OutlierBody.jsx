export default function OutlierBody({ flights }) {
  let overbookedFlightsObject = {}
  
  for (const flight in flights){
    //console.log(`the current flight origin is ${flights[flight]["origin"]}`)
    console.log(flights[flight]["origin"])
    let city = flights[flight]["origin"]
    if (flights[flight]["origin"] in overbookedFlightsObject){
      overbookedFlightsObject["origin"] = overbookedFlightsObject["origin"] + (flights[flight]["passengerCount"] - flights[flight]["capacity"]) 
    } else{
      overbookedFlightsObject[city] = flights[flight]["passengerCount"] - flights[flight]["capacity"]
    }
  }
  //console.log(`the object is ${JSON.stringify(overbookedFlightsObject)}`)
  let cityItems = []
  for (const [key,value] of Object.entries(overbookedFlightsObject)){
    if (key != "origin" || value < 0){

      cityItems.push(<li>The Airport is {`${key}`} and has an overbooked capacity of {`${value}`} </li>)
    }
  }

  //console.log(cityItems)
  return(
    <div className="outlier-body">
      <h1>City Data</h1>
      <ul>{cityItems}</ul>
    </div>
    
  )
}