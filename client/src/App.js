import React from 'react'
import { Provider } from 'react-redux'
import thunk from 'redux-thunk'
import { createStore , applyMiddleware } from 'redux'
import './App.css'
import Dashboard from './components/Dashboard/Dashboard.js'
import rootReducer from './reducers/index.js'


const  ep = "http://localhost:8080"

const store = createStore(rootReducer, applyMiddleware(thunk))

function App() {


  return (
    <Provider store={store}>
      <div className="App">
        <Dashboard endpoint={ep}/>
      </div>
    </Provider>
  )
}

export default App
