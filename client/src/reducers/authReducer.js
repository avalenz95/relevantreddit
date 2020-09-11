import { AUTH_SUCCESS, AUTH_ERROR } from '../actions/index.js'
// Deconstruct action to => {type, payload}
const authReducer = (state = false, { type, payload }) => {
    switch(type){
        case AUTH_SUCCESS:
            return payload.res
        
        case AUTH_ERROR:
            return payload.err

        default:
            return state
    }
}

export default authReducer