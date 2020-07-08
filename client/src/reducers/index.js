import { combineReducers } from 'redux'
import subredditReducer from './subredditReducer'
import usernameReducer from './usernameReducer'


export default combineReducers({
    subreddits: subredditReducer,
    name: usernameReducer
})