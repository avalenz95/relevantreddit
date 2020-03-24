import React from 'react';
import './App.css';
import AppAppBar from './components/navigation/navbar'
import RedditContent from './components/body';
import ToDoList from './components/body2';


function App() {
  return (
    <div className="App">
        <AppAppBar/>
        <RedditContent/>
    </div>
  );
}

export default App;
