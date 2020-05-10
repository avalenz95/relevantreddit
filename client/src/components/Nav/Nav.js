import React from 'react'
import axios from 'axios'

//Navigation for website
function Nav(props) {

    const {endpoint} = props
    let isAuthenticated = false

    //Authentication request
    function onAuth () {
        axios.get(endpoint + "/r/login").then((response) => {
            console.log(response)
            //Redirect user
            window.location.replace(response.request.responseURL)
            isAuthenticated = true
        })
    } 

    return (
        <nav className="navbar">
            {isAuthenticated ? 
            <span>Welcome!</span> 
            : 
            <button type="submit" onClick={onAuth}>Authenticate with Reddit</button>
            }
        </nav>
    )
}

export default Nav