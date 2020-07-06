import Cookies from 'js-cookie'
export const LOAD_USERNAME = 'LOAD_USERNAME'
export const USERNAME_SUCCESS = 'USERNAME_SUCCESS'
export const USERNAME_ERROR = 'USERNAME_ERROR'
export const LOAD_SUBREDDITS = 'LOAD_SUBREDDITS'

// Retrieving username was successful
export const usernameSuccess = (name) => {
    return {
        type: USERNAME_SUCCESS,
        payload: {name}
    }
}
// Retrieving username was unsuccessful
export const usernameError = (err) => {
    return {
        type: USERNAME_ERROR,
        payload: {err}
    }
}

// Get the current user from cookies
export const loadUsername = () => {
    return async(dispatch) => {
        const username = Cookies.get("username")

        if(username){
            dispatch(usernameSuccess)
        } else {
            dispatch(usernameError)
        }
    }
}

export const loadSubreddits = (subs) => {
    return async(dispatch) => {
        
    }
}