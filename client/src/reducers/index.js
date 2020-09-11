import { combineReducers } from 'redux'
import userdataReducer from './userdataReducer'
import usernameReducer from './usernameReducer'
import bannersReducer from './bannersReducer'
import authReducer from './authReducer'

export default combineReducers({
    userdata: userdataReducer,
    name: usernameReducer,
    banners : bannersReducer,
    auth: authReducer,
})