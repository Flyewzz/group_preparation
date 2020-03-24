import React from "react";
import ListItemText from "@material-ui/core/ListItemText";
import ListItemSecondaryAction from "@material-ui/core/ListItemSecondaryAction"
import Typography from "@material-ui/core/Typography"
import Link from "@material-ui/core/Link";
import {Link as RouterLink} from "react-router-dom";
import {makeStyles} from "@material-ui/core/styles";

const useStyles = makeStyles(() => ({
  text: {
    marginRight: 'auto',
    fontWeight: 'bold',
    fontSize: 'large'
  },
  filesText: {
    fontSize: 'medium',
    fontWeight: 'bold',
  },
}));

function Subject(props) {
  const styles = useStyles();

  return (
    <Link component={RouterLink} to={'/subject/' + props.subject.id}
          underline="none"
          color="inherit">
      <ListItemText disableTypography className={styles.text} primary={props.subject.name}/>
      <ListItemSecondaryAction>
        <Typography >
          {props.subject.semester}
        </Typography>
      </ListItemSecondaryAction>
    </Link>
  );
}

export default Subject;
