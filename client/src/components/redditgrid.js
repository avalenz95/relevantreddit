import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import RedditItem from './reddititem';

const useStyles = makeStyles(theme => ({
    root: {
      flexGrow: 1,
      direction: "column",
      justify: "space-around",
      alignItems: "center"
    },
  }));
  
  
  export default function RedditGrid(props) {
    const classes = useStyles();
    return (
      <div className={classes.root}>
        <Grid container container direction="row" justify="space-around" alignItems="center" spacing={4}>
            <RedditItem subreddit={props.subreddits}> </RedditItem>
        </Grid>
      </div>
    );
  }