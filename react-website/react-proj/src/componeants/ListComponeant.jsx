import { func } from "prop-types"

export default function ListComponeant( {indFlight} ){
  const [isHovered,setIsHovered] = useState(false)

  function handleMouseOver(){
    setIsHovered(true)
  }

  function handleMouseOut() {
    setIsHovered(false)
  }

  return (
  <li>
    <div onMouseOver={handleMouseOver} onMouseOut={handleMouseOut}>
      {isHovered ? (<>
          We start at {indFlight["origin"]} to {indFlight["destination"]}
        </>) : (<>
          We start at {indFlight["origin"]} to {indFlight["destination"]} and flight number is {indFlight["flightNumber"]}
          </>)
        }
    </div>
  </li>
  )
}