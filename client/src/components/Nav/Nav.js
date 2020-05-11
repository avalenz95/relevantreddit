import React from 'react'
import axios from 'axios'
import usePersistedState from '../../state'
import Cookies from 'js-cookie'

//Navigation for website
function Nav(props) {

    const {endpoint} = props

    const [auth, setAuth] = usePersistedState("auth")
    const [userName, setName] = usePersistedState("userName")

  

    //Authentication request
    function onAuth () {
        if (Cookies.get("username") !== undefined){
            setName(Cookies.get("username"))
        } else {

        
            axios.get(endpoint + "/r/login").then((response) => {
                console.log(response)
                //Redirect user
                setAuth(true)
                window.location.assign(response.request.responseURL);
                //window.open(response.request.responseURL, '_blank');
            })
        }
    } 

    return (
        <nav className="navbar">
            {auth === true ? 
            <span>Welcome! {userName}</span> 
            : 
            <button type="submit" onClick={onAuth}>Authenticate with Reddit</button>
            }
        </nav>
    )
}

export default Nav