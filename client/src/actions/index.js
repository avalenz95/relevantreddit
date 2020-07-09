import Cookies from 'js-cookie'
// Username
export const LOAD_USERNAME = 'LOAD_USERNAME'
export const USERNAME_SUCCESS = 'USERNAME_SUCCESS'
export const USERNAME_ERROR = 'USERNAME_ERROR'

// Subreddits
export const LOAD_USERDATA = 'LOAD_USERDATA'
export const USERDATA_SUCCESS = 'USERDATA_SUCCESS'
export const USERDATA_ERROR = 'USERDATA_ERROR'

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
            dispatch(usernameError())
        }
    }
}

export const userDataSuccess = (data) => {
    return {
        type: USERDATA_SUCCESS,
        payload: {data}
    }
}

export const userDataError = (err) => {
    return {
        type: USERDATA_ERROR,
        payload: {err}
    }
}

//CONSIDER DOING ALL OF THE PARSING FOR SUBREDDITS HERE AND RETURNING THE ARRAY OF COMPONT
// Thunk - similar to a call back, function that wraps another function(action)
export const loadUserData = (username) => {
    return async(dispatch) => {
        // Build url
        const url = "http://localhost:8080/user/" + username
        // Send a request
        try {
            const response = await fetch(url)
            const json = await response.json()
            // Send to dispatcher
            dispatch(userDataSuccess(json))
        } catch(err) {
            dispatch(userDataError(err))
        }
    }
}