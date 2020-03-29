import React from 'react';
import { makeStyles, withTheme } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import axios from "axios";
import Paper from '@material-ui/core/Paper';
import { Typography } from '@material-ui/core';
import RedditKeywords from './keywordbutton';

const useStyles = makeStyles(theme => ({
    root: {
      flexGrow: 1,
    },
    paper: {
      padding: theme.spacing(2),
      textAlign: 'center',
      color: 'white',
      backgroundImage: "url(https://www.w3schools.com/w3css/img_lights.jpg)",
      
    },
  }));


export default function RedditItem(props) {
  const classes = useStyles();
  const gridItems = []
  Object.entries(props.subreddit).map(([key,values], index) => {
  

    gridItems.push (
        <Grid item xs={3} className={classes.root}>
            <Paper className={classes.paper}>
                <Typography> {key} </Typography>
                <RedditKeywords values={values}></RedditKeywords>
            </Paper>
        </Grid>
    )
  })

  return (
      gridItems
  );
}