import { combineReducers } from 'redux'
import keywordReducer from './keywordReducer'
import usernameReducer from './usernameReducer'


export default combineReducers({
    keywords: keywordReducer,
    name: usernameReducer
})