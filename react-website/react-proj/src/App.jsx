import { useState } from 'react'
import Header from './componeants/Header.jsx'
import './App.css'
//Flights is is list of all the flight objects
import flights from "./data.js"
import ListBody from './componeants/ListBody.jsx'
import Body from './componeants/Body.jsx'
export default function App() {

  return (
    <>
      <Header className="head"/>
      <Body data={flights}/>
    </>
  )
}

//export default App
