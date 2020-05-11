import React , {useState} from 'react'
import axios from 'axios'
import Grid from '../Grid/Grid.js'


function Dashboard(props) {
    const {userName, endpoint} = props
    //Gets the users reddit content from db
    //let userName = Cookies.get("username")
    //const [userName, setName] = useState("")
    let [subreddits, setSubs] = useState(null)


//figure out a way to make sure this doesnt happen every single time the dashboard loads
    if (userName) {
        axios.get(endpoint + "/user/" + userName).then(response => {
            console.log(response)
            console.log("endpoint reached")
            setSubs(response.data.subreddits)
        })
    }

    return(
        <div className="dashboard">
            <Grid
                endpoint={endpoint}
                subreddits={subreddits}
                userName={userName}
            />
        </div>
    )
}

export default Dashboard