import React from 'react';
import './App.css';
import Nav from './components/Nav/Nav.js'

function App() {
  return (
    <div className="App">
      <Nav endpoint="http://localhost:8080"/>
    </div>
  );
}

export default App;
