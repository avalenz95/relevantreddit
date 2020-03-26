import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import GridItem from '@material-ui/core/Grid';
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
      color: theme.palette.text.secondary,
    },
  }));

export default function RedditItem() {
  const classes = useStyles();

  var gridItems = Object.entries(this.props.subreddit).map(([key,values]) => {

    return (
        <GridItem className={classes.root}>
            <Paper className={classes.paper}>
                <Typography> {key} </Typography>
                <RedditKeywords values={values}></RedditKeywords>
            </Paper>
        </GridItem>
    )
  })

  return (
      {gridItems}
  );
}