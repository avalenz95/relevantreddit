import { SUBREDDITS_SUCCESS, SUBREDDITS_ERROR } from '../actions/index.js'
// Deconstruct action to => {type, payload}
const subredditReducer = (state = null, { type, payload }) => {
    switch(type){
        case SUBREDDITS_SUCCESS:
            return payload.data
        
        case SUBREDDITS_ERROR:
            console.log(payload.err)
            return null

        default:
            return state
    }
}

export default subredditReducer