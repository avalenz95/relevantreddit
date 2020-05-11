import React, {Component} from "react"
import axios from "axios"
import Cookies from 'js-cookie'
import Grid from '../Grid/Grid.js'

const  endpoint = "http://localhost:8080"

class Dashboard extends Component {
    constructor(props){
        super(props)

        this.state = {
            subreddits: {},
            userName: ""
        }
    }

    //Check if component is active
    componentDidMount() {
        this.getContent()
    }

    //Get a users content
    getContent() {

        let userName = Cookies.get("username")

        if (userName) {
            this.setState({userName: {userName}})
            axios.get(endpoint + "/user/" + userName).then(response => {
                console.log(response);
                console.log("endpoint reached")
            })
        }
    }

    render() {
        return (
            <div className="dashboard">
                <Grid
                    endpoint={endpoint}
                    subreddits={this.state.subreddits}
                    userName={this.state.userName}
                />
            </div>
        )
    }
}

export default Dashboard;