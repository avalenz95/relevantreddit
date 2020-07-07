import { SUBREDDITS_SUCCESS, SUBREDDITS_ERROR } from '../actions/index.js'
// Deconstruct action to => {type, payload}
const keywordReducer = (state = [], { type, payload }) => {
    switch(type){
        case SUBREDDITS_SUCCESS:
            return payload.data
        
        case SUBREDDITS_ERROR:
            return payload.err

        default:
            return state
    }
}

export default keywordReducer