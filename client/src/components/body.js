import React, { Component } from "react";
import axios from "axios";
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import Cookies from 'js-cookie';
import RedditGrid from './redditgrid';

let endpoint = "http://localhost:8080";



class RedditContent extends Component {
    constructor(props) {
        super(props);


        this.state = {
            paperImgs: [],
            value: '',
            subreddits: {},
        };
}


    componentDidMount() {
      this.getContent()
    };

    //Get all images from endpoint
    getImages = () => {
      var images = []
      Object.entries(this.state.subreddits).map(([key,values], index) => {
        axios.get(endpoint + "/img/" + key).then((response) => {
    
          //Use image
          if(response.data != ""){
            images.push(response.data)
          } else {
            //Generic image
            images.push("https://www.w3schools.com/w3css/img_lights.jpg")
          }
          
          console.log(images)
        });
      });

      this.setState({paperImgs: images});

   }



    //When the user clicks authenticate with reddit
    onAuth = () => {
        axios.get(endpoint + "/r/login").then((response) => {
            console.log(response);
            //Redirect user
            window.location.replace(response.request.responseURL);
        })
    };

    //Gets the users reddit content from db
    getContent = () => {
      var userName = Cookies.get("username")
      if(userName){
        
        axios.get(endpoint + "/user/" + userName).then(response => {
          console.log(response);
          console.log("endpoint reached");

          this.setState({value: "Welcome: " + response.data.username, subreddits: response.data.subreddits,}
          , this.getImages);
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
          <RedditGrid subreddits={this.state.subreddits}></RedditGrid>
        </div>
        );
      }
}


export default RedditContent;

