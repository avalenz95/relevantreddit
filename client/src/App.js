import React, {useState} from 'react';
import './App.css';
import Nav from './components/Nav/Nav.js'
import Dashboard from './components/Dashboard/Dashboard.js'

const  ep = "http://localhost:8080"

function App() {
  const [userName, setName] = useState("")

  function handleName(updateName){
    setName(updateName)
  }

  return (
    <div className="App">
      <Nav userName={userName} endpoint={ep} onChange={handleName}/>
      <Dashboard userName={userName} endpoint={ep}/>
    </div>
  );
}

export default App;
