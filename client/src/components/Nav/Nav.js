import React, {useState}from 'react'
import axios from 'axios'
import Cookies from 'js-cookie'

//Navigation for website
function Nav(props) {

    const {endpoint, userName, onChange} = props

    const [auth, setAuth] = useState(false);

    //Handle the user getting a name
    function handleName(event) {
        event.target.userName = Cookies.get("username")

        onChange(event.target.userName)
    }

    //Authentication request
    function onAuth (event) {
        axios.get(endpoint + "/r/login").then((response) => {
            console.log(response)
            setAuth(true)
            //Redirect user
            window.location.replace(response.request.responseURL)
            //Get username
            handleName(event)
        })
    } 

    return (
        <nav className="navbar">
            {auth ? 
            <span>Welcome! {userName}</span> 
            : 
            <button type="submit" onClick={onAuth}>Authenticate with Reddit</button>
            }
        </nav>
    )
}

export default Nav