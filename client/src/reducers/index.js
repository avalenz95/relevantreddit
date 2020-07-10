import { combineReducers } from 'redux'
import userdataReducer from './userdataReducer'
import usernameReducer from './usernameReducer'
import bannersReducer from './bannersReducer'


export default combineReducers({
    userdata: userdataReducer,
    name: usernameReducer,
    banners : bannersReducer,
})