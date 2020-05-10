import React from 'react';
import logo from './logo.svg';
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
