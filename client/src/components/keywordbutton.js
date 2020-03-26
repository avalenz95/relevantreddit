import React from 'react';
import Button from '@material-ui/core/Button';
import { makeStyles } from '@material-ui/core/styles';
import DeleteIcon from '@material-ui/icons/Delete';
import IconButton from '@material-ui/core/IconButton';
import Typography from '@material-ui/core/Typography';

const useStyles = makeStyles(theme => ({
    root: {
      '& > *': {
        margin: theme.spacing(1),
      },
    },
  }));
  
  
  export default function RedditKeywords() {
    const classes = useStyles();

    var keywordButtons = []

    for (var i = 0; i < this.props.values.length; i++) {
        keywordButtons.push(
            <Button className={classes.root}>
                <Typography>{this.props.values[i]}</Typography>
                <IconButton aria-label="delete">
                    <DeleteIcon />
                </IconButton>

                <IconButton aria-label="delete">
                    <DeleteIcon />
                </IconButton>
            </Button>
        )
      }
  
    return (
        {keywordButtons}
    );
  }
