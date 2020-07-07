import React from "react"
import Grid from '../Grid/Grid.js'
import Nav from '../Nav/Nav.js'

const  endpoint = "http://localhost:8080"


function Dashboard(props) {

    return (
        <div className="dashboard">

            <Nav
                endpoint={endpoint} 
                userName={this.state.userName}
            />
                
            <Grid
                endpoint={endpoint}
                subreddits={this.state.subreddits}
                userName={this.state.userName}
            />
        </div>
    )
}

export default Dashboard