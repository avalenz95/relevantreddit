import { combineReducers } from 'redux'
import keywordReducer from './keywordReducers'


export default combineReducers({
    keywords: keywordReducer
})