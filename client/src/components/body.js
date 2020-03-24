import React, { Component } from "react";
import axios from "axios";
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';

let endpoint = "http://localhost:8080";


class RedditContent extends Component {
    constructor(props) {
        super(props);

        this.state = {
            items: []
        };
    }

    //When the user clicks authenticate with reddit
    onAuth = () => {
        axios.get(endpoint + "/r/login").then(response => {
            console.log(response);
            if (response.data) {
                this.getContent(response.data.username)
            }
        })
    };

    getContent = username => {
        axios.get(endpoint + "/u/" + username).then(response => {
            console.log(response);
            if (response.data) {
                this.setState({
                    items: response.data.map(item => {
                        return (
                            <Card>
                              <CardContent>
                                <Typography>
                                  {item.Subreddits}
                                </Typography>
                              </CardContent>
                            </Card>
                          );
                    })
                })
            } else {
                this.setState({
                    items: ["NOTHING"]
                });
            }
        });
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