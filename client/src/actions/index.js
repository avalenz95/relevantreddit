import Cookies from 'js-cookie'
// Username
export const LOAD_USERNAME = 'LOAD_USERNAME'
export const USERNAME_SUCCESS = 'USERNAME_SUCCESS'
export const USERNAME_ERROR = 'USERNAME_ERROR'

// Subreddits
export const LOAD_USERDATA = 'LOAD_USERDATA'
export const USERDATA_SUCCESS = 'USERDATA_SUCCESS'
export const USERDATA_ERROR = 'USERDATA_ERROR'
// Add a word
export const ADD_KEYWORD = 'ADD_KEYWORD'
export const KEYWORD_SUCCESS = 'KEYWORD_SUCCESS'
export const KEYWORD_ERROR = 'KEYWORD_ERROR'

export const addKeywordToSub = (subreddit, username, keyword) => {
    return async(dispatch) => {
        // Build Data 
        const data = {
            username: username,
            subreddit: subreddit,
            keyword: keyword,
        }

        // Build url
        const url = 'http://localhost:8080/addkeyword'

        try {
            // Send a post request
            const response = await fetch(url, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
                body: JSON.stringify(data),
            })

            await dispatch(keywordSuccess(response.status))
            await dispatch(loadUserData(username))

        } catch(err) {
            dispatch(keywordError(err))
    }
    }
}

// Retrieving username was successful
export const keywordSuccess = () => {
    return {
        type: KEYWORD_SUCCESS,
    }
}
// Retrieving username was unsuccessful
export const keywordError = (err) => {
    return {
        type: KEYWORD_ERROR,
        payload: {err}
    }
}

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
            await dispatch(usernameSuccess(username))
            await dispatch(loadUserData(username))
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