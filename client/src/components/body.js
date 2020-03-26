import React, { Component } from "react";
import axios from "axios";
import { makeStyles } from '@material-ui/styles';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import Cookies from 'js-cookie';

let endpoint = "http://localhost:8080";

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    square: 'true',
    backgroundColor: 'blue',
    display: 'flex',
    flexWrap: 'wrap',
  }
}));


class RedditContent extends Component {
    constructor(props) {
        super(props);


        this.state = {
            value: '',
            subreddits: [],
        };


        this.useStyles = makeStyles(theme => ({
          root: {
            flexGrow: 1,
          },
          paper: {
            square: 'true',
            backgroundColor: 'blue',
            display: 'flex',
            flexWrap: 'wrap',
          }
        }));
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
                <Grid item xs={12}>
                  <Paper variant="outlined" elevation={1} className={this.useStyles.paper}>Key:{k} {keywords}</Paper>
                </Grid>
              )
            })
          });
        });
      } else {
          console.log("endpoint not reached")
          this.setState({items: []});
      }

    };

    render() {

      const useStyles = makeStyles(theme => ({
        root: {
          flexGrow: 1,
          backgroundColor: 'red',
        },
        paper: {
          square: 'true',
          backgroundColor: 'blue',
          display: 'flex',
          flexWrap: 'wrap',
        }
      }));

        return (
        <div>
          <div>
            <Typography variant="h1" component="h2"> Subreddits </Typography>
            <Button color="inherit" type="submit" onClick={this.onAuth}>Authenticate with Reddit</Button>
          </div>
          <div className={useStyles.root}>
          <Grid container spacing={3}>
            {this.state.subreddits}
          </Grid>
          </div>
          <div className="row">{this.state.value}</div>
          <br></br>
          <div className="row">{this.state.subreddits}</div>
        </div>
        );
      }
}


export default RedditContent;

