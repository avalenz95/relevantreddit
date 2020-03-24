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
            subreddits: [],
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
      
            subreddits: Object.entries(response.data.subreddits).map(([k,v]) => {

              const keywords = []
              //Create buttons for each keyword
              for (var i = 0; i < v.length; i++) {
                keywords.push(<Button>{v[i]}</Button>)
              }

              //Display key alongside keywords
              return (
                <div className="row">
                  <Card>Key:{k} {keywords}</Card>       
                </div>
              )
            })
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
          <br></br>
          <div className="row">{this.state.subreddits}</div>
        </div>
        );
      }
}


export default RedditContent;