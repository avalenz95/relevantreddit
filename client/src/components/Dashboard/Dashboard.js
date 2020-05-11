import React from 'react'
import axios from 'axios'
import Cookies from 'js-cookie'

let endpoint = "http://localhost:8080"

function Dashboard() {
    //Gets the users reddit content from db
    let userName = Cookies.get("username")

    //figure out a way to make sure this doesnt happen every single time the dashboard loads
    function getData() {
        if (userName) {
            axios.get(endpoint + "/user/" + userName).then(response => {
                console.log(response)
                console.log("endpoint reached")
                //response.data.username
                //response.data.subreddits
                //getImages
            })
        }

    }

    return(
        <div className="dashboard">
            
        </div>
    )
}

export default Dashboard