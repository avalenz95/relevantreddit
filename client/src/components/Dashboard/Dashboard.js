import React from "react"
import Grid from '../Grid/Grid.js'
import Nav from '../Nav/Nav.js'
import { loadUsername } from "../../actions/index.js"

const  endpoint = "http://localhost:8080"


function Dashboard() {

    const [name, setName] = useState("")
    const dispatch = useDispatch() // Get the dispatcher

    // Attempt to load username on component mount
    useEffect(() => {
        dispatch(loadUsername())
    })

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