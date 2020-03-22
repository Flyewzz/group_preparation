import AppBar from "@material-ui/core/AppBar"
import Toolbar from "@material-ui/core/Toolbar";
import Button from "@material-ui/core/Button"
import React from "react";
import {makeStyles} from "@material-ui/core/styles";

const useStyles = makeStyles(() => ({
  color: {
    color: '#ffffff',
    borderColor: '#ffffff',
    marginLeft: '10pt',
  },
  appName: {
    fontSize: 'x-large',
    textTransform: 'none',
    color: '#ffffff',
    marginRight: 'auto',
  },
}));

function Header() {
  const classes = useStyles();

  return (
    <AppBar position="static">
      <Toolbar>
        <Button className={classes.appName}>ExamPrep</Button>
        <Button href="#" className={classes.color}>Sign In</Button>
        <Button variant="outlined" className={classes.color}>Sign Up</Button>
      </Toolbar>
    </AppBar>
  );
}

export default Header;
