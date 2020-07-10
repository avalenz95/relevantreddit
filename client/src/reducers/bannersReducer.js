import { BANNERS_SUCCESS, BANNERS_ERROR } from '../actions/index.js'
// Deconstruct action to => {type, payload}
const bannersReducer = (state = {}, { type, payload }) => {
    switch(type){
        case BANNERS_SUCCESS:
            return payload.data
        
        case BANNERS_ERROR:
            return {}

        default:
            return state
    }
}

export default bannersReducer