import React, { Component } from "react";
import axios from "axios";
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import Cookies from 'js-cookie';

let endpoint = "http://localhost:8080";

class RedditContent extends Component {
    constructor(props) {
        super(props);

        this.state = {
            items: []
        };
    }


    componentDidMount() {
      this.getContent()
    }


    //When the user clicks authenticate with reddit
    onAuth = () => {
        axios.get(endpoint + "/r/login").then((response) => {
            console.log(response);
            window.location.replace(response.request.responseURL);
        })
    };

    getContent = () => {
      var userName = Cookies.get("username")
      if(userName){
        
        axios.get(endpoint + "/user/" + userName).then(response => {
          console.log(response);
          console.log("endpoint reached");
        });
      } else {
          console.log("endpoint not reached")
      }

    };

    render() {
        return (
        <div>
          <div>
            <Typography variant="h1" component="h2"> Subreddits </Typography>
            <Button color="inherit" type="submit" onClick={this.onAuth}>Authenticate with Reddit</Button>
          </div>

          <div className="row">{this.state.items}</div>
        </div>
        );
      }
}


export default RedditContent;