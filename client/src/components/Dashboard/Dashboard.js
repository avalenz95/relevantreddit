import React from 'react'
import axios from 'axios'
import Cookies from 'js-cookie'
import Grid from '../Grid/Grid.js'

let endpoint = "http://localhost:8080"

function Dashboard() {
    //Gets the users reddit content from db
    let userName = Cookies.get("username")
    let subreddits = null

//figure out a way to make sure this doesnt happen every single time the dashboard loads
    if (userName) {
        axios.get(endpoint + "/user/" + userName).then(response => {
            console.log(response)
            console.log("endpoint reached")
            subreddits = response.data.subreddits
        })
    }

    return(
        <div className="dashboard">
            <Grid
                endpoint={endpoint}
                userName={userName}
                subreddits={subreddits}

            />
        </div>
    )
}

export default Dashboard