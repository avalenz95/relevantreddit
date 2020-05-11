import React from 'react'
import axios from 'axios'
import Cookies from 'js-cookie'
import usePersistedState from '../../state'

//Navigation for website
function Nav(props) {

    const {endpoint, userName, onChange} = props

    const [auth, setAuth] = usePersistedState("auth", false)
    //const [auth, setAuth] = useState(false);

    //Handle the user getting a name
    function handleName(event) {
        event.target.userName = Cookies.get("username")

        onChange(event.target.userName)
    }

    //Authentication request
    function onAuth (event) {
        axios.get(endpoint + "/r/login").then((response) => {
            console.log(response)
            //Redirect user
            window.location.assign(response.request.responseURL);
            //window.open(response.request.responseURL, '_blank');
        }).then(
            handleName(event),
            setAuth(true)
            )
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