import React from 'react';
import './App.css';
import AppAppBar from './components/navigation/navbar'
import RedditContent from './components/body';


function App() {
  return (
    <div className="App">
        <AppAppBar/>
        <RedditContent/>
    </div>
  );
}

export default App;
