import { USERNAME_SUCCESS, USERNAME_ERROR } from '../actions/index.js'
// Deconstruct action to => {type, payload}
const usernameReducer = (state = "", { type, payload }) => {
    switch(type){
        case USERNAME_SUCCESS:
            return payload.name
        
        case USERNAME_ERROR:
            return ""

        default:
            return state
    }
}

export default usernameReducer