import React from 'react';
import './App.css';
import Nav from './components/Nav/Nav.js'
import Dashboard from './components/Dashboard/Dashboard.js'
import usePersistedState from './state'


const  ep = "http://localhost:8080"

function App() {


  const [userName, setName] = usePersistedState("userName", "")
  usePersistedState("auth", false)

  return (
      <div className="App">
        <Nav endpoint={ep} />
        <Dashboard userName={userName} endpoint={ep}/>
      </div>
  );
}

export default App;
