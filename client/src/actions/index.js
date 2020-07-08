import Cookies from 'js-cookie'
// Username
export const LOAD_USERNAME = 'LOAD_USERNAME'
export const USERNAME_SUCCESS = 'USERNAME_SUCCESS'
export const USERNAME_ERROR = 'USERNAME_ERROR'

// Subreddits
export const LOAD_SUBREDDITS = 'LOAD_SUBREDDITS'
export const SUBREDDITS_SUCCESS = 'SUBREDDITS_SUCCESS'
export const SUBREDDITS_ERROR = 'SUBREDDITS_ERROR'

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
        // Send actions to dispatcher
        if(username){
            dispatch(usernameSuccess(username))
        } else {
            dispatch(usernameError(""))
        }
    }
}

export const subredditsSuccess = (data) => {
    return {
        type: SUBREDDITS_SUCCESS,
        payload: {data}
    }
}

export const subredditsError = (err) => {
    return {
        type: SUBREDDITS_ERROR,
        payload: {err}
    }
}

// Thunk - similar to a call back, function that wraps another function(action)
export const loadSubreddits = (username) => {
    return async(dispatch) => {
        // Build url
        const url = "http://localhost:8080/user/" + username
        // Send a request
        try {
            const response = await fetch(url)
            const json = await response.json()
            // Send to dispatcher
            dispatch(subredditsSuccess(json))
        } catch(err) {
            dispatch(subredditsError(err))
        }
    }
}