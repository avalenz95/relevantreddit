import React from 'react';
import './App.css';
import Nav from './components/Nav/Nav.js'
import Dashboard from './components/Dashboard/Dashboard.js'


const  ep = "http://localhost:8080"

function App() {


  return (
      <div className="App">
        <Dashboard endpoint={ep}/>
      </div>
  );
}

export default App;
