import { USERDATA_SUCCESS, USERDATA_ERROR } from '../actions/index.js'
// Deconstruct action to => {type, payload}
const userDataReducer = (state = null, { type, payload }) => {
    switch(type){
        case USERDATA_SUCCESS:
            return payload.data
        
        case USERDATA_ERROR:
            console.log(payload.err)
            return null

        default:
            return state
    }
}

export default userDataReducer