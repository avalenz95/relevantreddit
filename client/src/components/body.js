import React, { Component } from "react";
import axios from "axios";
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import Card from '@material-ui/core/Card';
import Cookies from 'js-cookie';

let endpoint = "http://localhost:8080";

class RedditContent extends Component {
    constructor(props) {
        super(props);

        this.state = {
            value: '',
            items: [],
            //items: []
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

          this.setState({
            value: "Welcome: " + response.data.username,
            items : Object.keys(response.data.subreddits).map(item =>{
              return (
                <div className="row">
                  <Card>{item}</Card>
                </div>
              )
            }),
            // items: response.data.subreddits.map(item => {

            //   return (
            
            //     <Card>{item}</Card>
            //   );
            // })
          });

        });
      } else {
          console.log("endpoint not reached")
          this.setState({items: []});
      }

    };

    render() {
        return (
        <div>
          <div>
            <Typography variant="h1" component="h2"> Subreddits </Typography>
            <Button color="inherit" type="submit" onClick={this.onAuth}>Authenticate with Reddit</Button>
          </div>

          <div className="row">{this.state.value}</div>
          <div className="row">{this.state.items}</div>
        </div>
        );
      }
}


export default RedditContent;