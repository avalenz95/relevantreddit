import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import GridItem from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';

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

  return (
    <GridItem className={classes.root}>
      <Paper className={classes.paper}>Key {}</Paper>
    </GridItem>
  );
}