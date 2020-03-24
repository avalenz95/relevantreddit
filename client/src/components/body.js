import React, { Component } from "react";
import axios from "axios";
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import Typography from '@material-ui/core/Typography';

let endpoint = "http://localhost:8080";


class RedditContent extends Component {
    constructor(props) {
        super(props);

        this.state = {
            items: []
        };
    }

    componentDidMount() {
        this.getContent();
    }

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
            <div className="row">
            </div>
            <div className="row">
            </div>
            <div className="row">
            </div>
          </div>
        );
      }
}


export default RedditContent;