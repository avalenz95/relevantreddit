import { combineReducers } from 'redux'
import userdataReducer from './userdataReducer'
import usernameReducer from './usernameReducer'


export default combineReducers({
    userdata: userdataReducer,
    name: usernameReducer
})