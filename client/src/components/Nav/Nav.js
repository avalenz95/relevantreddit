import React from 'react'
import axios from 'axios'
import { useSelector } from 'react-redux'
import './Nav.css'


//Navigation for website
function Nav() {
    const endpoint = "http://localhost:8080"
    // name of reducer (its where the state is stored)
    const name = useSelector(state => state.name)

    //Authentication request
    function onAuth () {
        axios.get(endpoint + "/r/login").then((response) => {
            console.log(response)

            window.location.assign(response.request.responseURL);
            //window.open(response.request.responseURL, '_blank');
        })
     } 

    return (
        <nav className="navbar">
            <div className="title">
                RELEVANT
            </div>

            <div className="appusername">
                {name !== "" ? 
                <span>u/{name}</span> 
                : 
                <button className="authButton" type="submit" onClick={onAuth}>
                    <div className="authText">
                        Authenticate with Reddit
                    </div>
                    <div className="authIcon">
                        <img className="authIcon" src={process.env.PUBLIC_URL + '/reddit-icon.svg'} alt="redditauth" />  
                    </div>
                </button>
                }
            </div>
        </nav>
    )
}

export default Nav